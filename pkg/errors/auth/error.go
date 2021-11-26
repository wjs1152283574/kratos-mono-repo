/*
 * @Author: Casso
 * @Date: 2021-11-17 16:24:19
 * @LastEditors: Casso
 * @LastEditTime: 2021-11-26 17:51:03
 * @Description: file content
 * @FilePath: /kratos-mono-repo/pkg/errors/auth/error.go
 */
package auth

import "github.com/go-kratos/kratos/v2/errors"

var ErrAuthFail = errors.New(401, "Authentication failed", "Missing token or token incorrect")
