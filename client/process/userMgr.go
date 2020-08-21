package process

import (
	"fmt"
	//引入"message"包
)


//客户端要维护的map
var onlineUsers map[int]*message.User = make(map[int]*message.User,10)

//在客户端显示当前的在线客户
func outputOnlineUser(){
	//遍历一把 onlineUsers
	fmt.Println("当前在线用户列表")
	for id, _ := range onlineUsers {
		fmt.Println("用户ID：\t",id)
	}
}







//编写一个方法，处理返回的NotifyUserStatusMes
func updataUserStatus(notifyUserStatusMes *message.notifyUserStatusMes){

	//适当优化一下
	user,ok := onlineUsers[notifyUserStatusMes.UserID]
	if !ok{//原来没有
		user = &message.User{
			UserID : notifyUserStatusMes.UserID,
		}
	}
	user.UserStatus =  notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserID] = user
	outputOnlineUser()
}