package message

//定义一个用户的结构体
//为了序列和反序列话，加上json tag
type User struct {
	//确定字段信息
	UserID int 'json:"userID"'
	UserPwd string'json:"userPwd"'
	UserName string'json:"userName"'
	UserStatus int 'json:"userStatus"'//用户状态
	Gender string 'json:"gender"'
}