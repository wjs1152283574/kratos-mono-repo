package pagination

// GetPageOffset 根据页码和每页数量计算偏移量
func GetPageOffset(pageNum, pageSize int64) int64 {
	return (pageNum - 1) * pageSize
}
