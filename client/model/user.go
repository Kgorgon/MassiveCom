package model

//定义一个用户的结构体
//为了序列和反序列话，加上json tag
type Userstruct {
	//确定字段信息
	UserID int 'json:"userID"'
	UserPwd string'json:"userPwd"'
	UserName string'json:"userName"'
}