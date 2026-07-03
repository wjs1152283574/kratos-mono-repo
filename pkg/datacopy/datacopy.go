package datacopy

import (
	"encoding/json"
	"reflect"
)

// Copy 基于 JSON 序列化/反序列化的数据复制
// 适用于请求体数据映射到 DTO 模型
func Copy(data, res interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, res)
}

// CopyReflect 基于 reflect 的数据复制，将 src 字段逐一赋值到 dst
func CopyReflect(src, dst *reflect.Value) *reflect.Value {
	for i := 0; i < src.NumField(); i++ {
		fieldName := src.Type().Field(i).Name
		dstField := dst.FieldByName(fieldName)
		if dstField.IsValid() && dstField.CanSet() {
			dstField.Set(src.Field(i))
		}
	}
	return dst
}
