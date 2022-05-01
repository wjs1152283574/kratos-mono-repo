package tool

import (
	"fmt"

	"gorm.io/gorm"
)

type ExecFun func(db *gorm.DB) error

// 简易全局 [GORM] 事务封装
// 事务具有隔离性，如果一个事务内业务需要等待其他事务的结果数据，则最好合并在同一个事务中避免脏读/幻读
// 以下方法适用于多个事务在数据上互不干涉，只需要保证同时成功或失败
// 如果业务需要必须要依赖于其他事物正在进行的数据，则可以添加一个 *gorm.DB 入参，保证所有事务串行在同一个操作中
func Tx(execs []ExecFun, gdb *gorm.DB) (err error) {
	tx := gdb.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			err = fmt.Errorf("%v", err)
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

// NewExec 生成新的事务执行方法 暂时只支持sql语句
func NewExec(args string) ExecFun {
	return func(db *gorm.DB) error {
		return db.Exec(args).Error
	}
}
