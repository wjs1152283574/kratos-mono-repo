package contextkey

// 自定义类型，用户context赋值
type Key string
type Val string

var (
	UserID = "userid"
)

// NewKey return Key with key name
func NewKey() Key {
	return Key(UserID)
}
