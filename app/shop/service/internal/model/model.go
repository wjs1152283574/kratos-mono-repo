package model

// shop model

type ShopLog struct {
	ID      uint   `gorm:"primarykey"`
	OrderSn string `gorm:"NOT NULL;COMMENT:订单号（不可使用order关键词，mysql查询会报错）"`

	UpdatedTime int64 `gorm:"type:bigint(20);COMMENT:最后修改时间"`
	CreatedTime int64 `gorm:"type:bigint(20);COMMENT:创建时间"`
	DeleteTime  int64 `gorm:"type:bigint(20);COMMENT:删除时间"`
}
