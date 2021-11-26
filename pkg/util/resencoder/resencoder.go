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
	"github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/grpc/encoding"
)

type Res struct {
	Code int
	Data interface{}
	Msg  string
}

func CustomResponeDeco() http.EncodeResponseFunc {
	return func(w http.ResponseWriter, r *http.Request, v interface{}) error {
		reply := &Res{ // 将状态码都改为200
			Code: 200,
			Data: v,
		}
		codc := encoding.GetCodec("json")
		data, err := codc.Marshal(reply)
		if err != nil {
			return err
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
		return nil
	}
}
