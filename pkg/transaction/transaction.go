package transaction

import (
	"fmt"

	"gorm.io/gorm"
)

// ExecFun 事务执行函数
type ExecFun func(db *gorm.DB) error

// Tx 简易 GORM 事务封装
// 事务具有隔离性，如果事务内业务需要等待其他事务的结果数据，则最好合并在同一个事务中
// 此方法适用于多个事务在数据上互不干涉，只需保证同时成功或失败
func Tx(execs []ExecFun, gdb *gorm.DB) (err error) {
	tx := gdb.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			err = fmt.Errorf("%v", r)
		}
	}()

	for _, f := range execs {
		err = f(tx)
		if err != nil {
			tx.Rollback()
			return
		}
	}

	return tx.Commit().Error
}

// NewExec 生成新的事务执行方法，暂时只支持 SQL 语句
func NewExec(args string) ExecFun {
	return func(db *gorm.DB) error {
		return db.Exec(args).Error
	}
}
