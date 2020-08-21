package message


const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMesType         = "RegosterMes"
	RegisterResMesType      = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMes"
)

type Message struct {
	Type string `json:"type"`//消息类型
	Data string `json:"data"`//消息内容
}

//定义几个用户状态的常量
const (
	UserOnline = iota
	UserOffline 
	UserBusyStatus
)




//先定义两个消息..后面需要再增加

type LoginMes struct {
	UserID int `json:"userID"`
	UserPwd string `json:"userPwd"`
	UserName string `jsonL:"userName"`
}

type LoginResMes struct {
	Code int `json:"code"`//返回状态码 500表示用户未注册 200表示登录成功
	UsersID []int          //增加字段，保存用户ID的切片
	Error string `json:"error"`//返回错误信息
}


type RegisterMes struct {
	User User `json:"user"`//类型就是User结构体
}

type RegisterResMes struct {
	Code int `json:"code"`//返回状态码 500表示用户未注册 200表示成功
	Error string `json:"error"`//返回错误信息
}

//为了配合服务器端推送通知用户状态变化的消息
type NotifyUserStatusMes struct {
	UserID int `json:"userID"`
	Status int `json:"status"`
}

//增加一个SmsMes //发送的消息
type SmsMes sturct {
	Content string `json:"content"`
	User //匿名结构体 继承
}


//SmsResMes