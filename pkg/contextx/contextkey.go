package contextx

// Key 自定义类型，用于 context 赋值，避免键冲突
type Key string

// UserID 是上下文中存储用户 ID 的键名
const UserID Key = "userid"

// NewKey 返回 UserID 键
func NewKey() Key {
	return UserID
}
