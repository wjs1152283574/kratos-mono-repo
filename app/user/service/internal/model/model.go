package model

// user model

var (
	UserTableName = "user"
)

type User struct {
	ID     uint   `gorm:"primarykey"`
	Mobile string `gorm:"unique;COMMENT:手机号"`
	Pass   string `gorm:"COMMENT:密码"`
	Name   string `gorm:"COMMENT:用户名"`
	Age    int64  `gorm:"COMMENT:年龄"`

	UpdatedTime int64 `gorm:"type:bigint(20);COMMENT:最后修改时间"`
	CreatedTime int64 `gorm:"type:bigint(20);COMMENT:创建时间"`
	DeleteTime  int64 `gorm:"type:bigint(20);COMMENT:删除时间"`
}

func (u *User) TableName() string {
	return UserTableName
}
