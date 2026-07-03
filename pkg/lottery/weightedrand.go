package lottery

import (
	"math/rand"
	"sort"
	"time"

	wr "github.com/mroth/weightedrand"
)

// Prize 奖品
type Prize struct {
	PlayerId int64
	Weight   int
}

// RandomDraw 权重随机抽奖
func RandomDraw(prizes []*Prize) int64 {
	var weightSum int
	for _, v := range prizes {
		weightSum += v.Weight
	}

	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(weightSum)

	randomNumTmp := &Prize{PlayerId: -1, Weight: randomNum}
	concatWeightArr := make([]*Prize, 0, len(prizes)+1)
	for _, p := range prizes {
		concatWeightArr = append(concatWeightArr, &Prize{PlayerId: p.PlayerId, Weight: p.Weight})
	}
	concatWeightArr = append(concatWeightArr, randomNumTmp)

	sort.Slice(concatWeightArr, func(i, j int) bool {
		return concatWeightArr[i].Weight < concatWeightArr[j].Weight
	})
	sort.Slice(prizes, func(i, j int) bool {
		return prizes[i].Weight < prizes[j].Weight
	})

	var randomNumIndex = -1
	for p, v := range concatWeightArr {
		if v.Weight == randomNum {
			randomNumIndex = p
		}
	}
	if randomNumIndex >= len(prizes) {
		randomNumIndex = len(prizes) - 1
	}

	res := prizes[randomNumIndex]
	return res.PlayerId
}

// GetRandPrize 随机获奖（推荐方案），输入盲盒数量，输出中奖 ID:数量
func GetRandPrize(boxNum int64, prize []Prize) map[int64]int64 {
	p := make(map[int64]int64, boxNum)

	for i := 0; i < int(boxNum); i++ {
		var wrs []wr.Choice
		for _, v := range prize {
			wrs = append(wrs, wr.Choice{
				Item:   v.PlayerId,
				Weight: uint(v.Weight),
			})
		}
		result := RandPick(wrs)
		p[result] += 1
		for i, v := range prize {
			if v.PlayerId == result {
				prize[i].Weight -= 1
			}
		}
	}

	return p
}

// RandPick 传入奖池，返回中奖 item
func RandPick(wrs []wr.Choice) int64 {
	rand.Seed(time.Now().UTC().UnixNano())
	chooser, _ := wr.NewChooser(wrs...)
	return chooser.Pick().(int64)
}
