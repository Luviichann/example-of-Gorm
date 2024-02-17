package models

type User struct {
	Id       int
	Username string
	Age      int
	Email    string
	AddTime  int
}

// 结构体对应的表明，Gorm给了默认情况的表名，这里是自定义表名。
// 意思是User这个结构体连接的数据表名字是user。
func (User) TableName() string {
	return "user"
}
