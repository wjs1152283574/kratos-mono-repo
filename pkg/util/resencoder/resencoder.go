/*
 * @PackageName: resencoder
 * @Description: 自定义返回数据编码方式
 * @Author: Casso-Wong
 * @Date: 2021-10-29 12:32:52
 * @Last Modified by: Casso-Wong
 * @Last Modified time: 2021-10-29 12:32:52
 */
package resencoder

import (
	"net/http"
)

func CustomResponeDeco(w http.ResponseWriter, r *http.Request, v interface{}) error {
	w.Write([]byte("casso"))
	return nil
}
