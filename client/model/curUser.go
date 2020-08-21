package model

import (
	"fmt"
	"net"
	//引入"message"包
	"massive/client/model"
)
var onlineUsers map[int]*message.User = make(map[int]*message.User,10)
var CurUser model.CurUser//我们在用户登陆成功后，完成对CurUser的初始化

//因为在客户端，很多地方会使用curUser，我们将其作为一个全局变量
type CurUser struct {
	Conn net.Conn
	message.User
}