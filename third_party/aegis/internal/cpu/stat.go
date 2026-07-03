package cpu

import (
	"sync/atomic"
	"time"
)

const (
	interval time.Duration = time.Millisecond * 500
)

var (
	stats CPU
	usage uint64
)

// CPU is cpu stat usage.
type CPU interface {
	Usage() (u uint64, e error)
	Info() Info
}

func init() {
	var (
		err error
	)
	stats, err = newCgroupCPU()
	if err != nil {
		stats, err = newPsutilCPU(interval)
		if err != nil {
			// On macOS or non-Linux systems, cpu stats may not be available.
			// Use a no-op implementation instead of panicking.
			stats = &noopCPU{}
		}
	}
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for {
			<-ticker.C
			u, err := stats.Usage()
			if err == nil && u != 0 {
				atomic.StoreUint64(&usage, u)
			}
		}
	}()
}

// Stat cpu stat.
type Stat struct {
	Usage uint64 // cpu use ratio.
}

// Info cpu info.
type Info struct {
	Frequency uint64
	Quota     float64
}

// ReadStat read cpu stat.
func ReadStat(stat *Stat) {
	stat.Usage = atomic.LoadUint64(&usage)
}

// GetInfo get cpu info.
func GetInfo() Info {
	return stats.Info()
}

// noopCPU is a no-op CPU stats implementation for systems where
// cgroup and psutil are unavailable (e.g., macOS with CGO_ENABLED=0).
type noopCPU struct{}

var _ CPU = (*noopCPU)(nil)

func (n *noopCPU) Usage() (uint64, error) { return 0, nil }
func (n *noopCPU) Info() Info             { return Info{} }
