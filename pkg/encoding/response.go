package encoding

import (
	"encoding/json"
	ht "net/http"

	"github.com/go-kratos/kratos/v2/transport/http"
)

// ResponseEncoder 自定义 HTTP JSON 响应编码器
func ResponseEncoder() http.EncodeResponseFunc {
	return func(w ht.ResponseWriter, r *ht.Request, v interface{}) error {
		data, err := json.Marshal(v)
		if err != nil {
			return err
		}
		_, err = w.Write(data)
		return err
	}
}
