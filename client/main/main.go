package main

import (
	"fmt"
	"os"
	"massivecom/client/process"
)
//定义两个全局变量，一个表示用户ID
//另一个表示密码
var userID int
var unserPwd string

func main(){
	//接收用户的选择
	var key int
	//判断是否还继续显示菜单
	//var loop = true
	for {
		fmt.Println("------------欢迎登陆多人聊天系统------------")
		fmt.Println("\t\t\t 1 登录聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择(1-3)")
		fmt.Scanf("%d\n",&key)
		switch key {
			case 1 :
				fmt.Println("登陆聊天室")
				fmt.Println("请输入用户的ID")
				fmt.Scanf("%d\n",&userID)
				fmt.Println("请输入密码")
				fmt.Scanf("%s\n",&userPwd)
				//完成登录
				//1.创建一个UserProcess的实例
				up := &process.UserProcess {}
				up.Login(userID,userPwd)
			case 2 :
				fmt.Println("注册用户")
				fmt.Println("请输入用户的ID:")
				fmt.Scanf("%d\n",&userID)
				fmt.Println("请输入用户密码:")
				fmt.Scanf("%s\n",&userPwd)
				fmt.Println("请输入用户名字(nickname):")
				fmt.Scanf("%s\n",&userName)
				//2.调用一个UserProcess实例，完成注册请求
				up := &process.UserProcess {}
				up.Register(userID,userPwd,userName)

			case 3 :
				fmt.Println("退出系统")
				os.Exit(0)
				//loop = false
			default :
				fmt.Println("输入有误，请重新输入")
		}
	}



	//根据用户的输入显示新的提示信息
	// if key == 1{
	// 	//用户登录
	// 	fmt.Println("请输入用户的ID")
	// 	fmt.Scanf("%d\n",&userID)
	// 	fmt.Println("请输入密码")
	// 	fmt.Scanf("%s\n",&userPwd)
		
	// 	//因为使用了新的程序结构，我们创建一个
		
		
		
	// 	//先把登录的函数写到另外一个文件，比如login.go
		
		
		
	// 	//login(userID,userPwd)
	// 	//if err != nil {
	// 	//	fmt.Println("登陆信息错误")
	// 	//} else {
	// 	//	fmt.Println("登录成功")
	// 	//}
	// } else  if key == 2 {
	// 	fmt.Println("进行用户注册的逻辑")
	// }
}