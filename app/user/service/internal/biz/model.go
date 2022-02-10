/*
 * @PackageName: biz
 * @Description: data model
 * @Author: Casso
 * @Date: 2022-02-10 17:12:57
 * @LastModifiedBy: Casso
 * @LastEditTime: 2022-02-10 17:12:57
 */

package biz

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Mobile string `gorm:"unique"`
	Pass   string
	Name   string
	Age    int64
}

type UserReply struct {
	Name, Mobile string
	Age, ID      int64
}

type UserForToken struct {
	Mobile string
	Pass   string
	ID     int
}
