package process

import (
	"fmt"
	"net"
	"encoding/binary"
	"encoding/json"
	//记得引入"message包"
	"massivecom/client/utils"
	"os"
)


type UserProcess struct {
	//暂时不需要任何字段...
	
}


func(this *UserProcess)Register(userID int, 
	userPwd string,userName string)(err error){
	
	//1.连接到服务器端
	conn,err := net.Dial("tcp","localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err= ",err)
		return
	}
	//延时关闭
	defer conn.Close()
	//2.准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.RegisterMesType
	//3.创建一个LoginMes 结构体
	var registerMes message.RegisterMesType
	registerMes.User.UserID = userID
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = user Name
	//4.将registerMes序列化
	data,err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("register json.Marshal err=",err)
		return
	}
	//5.把data赋给mes.Data字段
	mes.Data = string(data)
	//6.将mes进行序列化
	data,err = json.Marshal(mes)  //data是[]byte
	if err != nil {
		fmt.Println("json.Marshal err=",err)
		return
	}
	//创建一个Transfer实例
	tf := &utils.Transfer {
		Conn : conn,
	}
	//发送data给服务器端
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("注册发送信息错误 err="，err)
	}
	
	mes,err = tf.ReadPkg()//mes 就是RegisterResMes
	
	if err != nil{
		fmt.Println("readPkg(conn) failed=",err)
		return
	}
	//将mes的data部分反序列化成 RegisterResMes
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data),&registerResMes)
	if registerResMes.Code == 200{
		fmt.Println("注册成功！请重新登陆")
		os.Exit(0)
	} else {
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}
	return

}






//关联一个用户登录的方法
func (this *UserProcess)Login(userID int, userPwd string)(err error){
	//下一步开始定义协议..
	//fmt.Printf("userID = %d userPwd=%s",userID,userPwd)
	//return nil


	//1.连接到服务器端
	conn,err := net.Dial("tcp","localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err= ",err)
		return
	}
	//延时关闭
	defer conn.Close()
	//2.准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.LoginMesType
	//3.创建一个LoginMes
	var loginMes message.LoginMes
	loginMes.UserID = userID
	loginMes.UserPwd = userPwd

	//4.将loginMes序列化
	data,err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err=",err)
		return
	}
	//5.把data赋给mes.Data字段
	mes.Data = string(data)
	//6.将mes进行序列化
	data,err = json.Marshal(mes)  //data是[]byte
	if err != nil {
		fmt.Println("json.Marshal err=",err)
		return
	}
	//7.到这个时候，data就是我们要发送的消息
	//7.1 先把data的长度发送给服务器
	// 先获取到data的长度 -> 转成一个表示长度的byte切片
	var pkgLen uint32
	pkg = uint32(len(data))
	var buf [4]byte 
	binary.BigEndian.PutUint32(buf[0:4],pkgLen)
	//发送长度
	n,err := conn.Write(buf[:4])
	if n!=4 || err !=nil {
		fmt.Println("conn.Write(bytes) failed= ",err)
		return
	}
	fmt.Println("客户端，发送消息的长度=%d 内容是 =%s",len(data),string(data))
	
	
	//发送消息本身
	_,err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) failed= ",err)
		return
	}
	//这里还需要处理服务器端返回的消息
	
	//创建一个Transfer实例
	tf := &utils.Transfer{
		Conn : conn,
	}
	mes,err = tf.ReadPkg()//mes 就是
	
	if err != nil{
		fmt.Println("readPkg(conn) failed=",err)
		return
	}
	
	
	//将mes的data部分反序列化成LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data),&loginResMes)
	if loginResMes.Code == 200{
		//初始化CurUser
		CurUser.Conn = conn
		CurUser.UserID = userID
		CurUser.UserStatus = message.UserOnline
		//fmt.Println("用户登陆成功")
		
		//现在可以显示当前在线用户的列表，遍历loginResMes.UsersID
		fmt.Println("当前在线用户:")
		for _, v := range loginResMes.UsersID {
			
			//如果我们要求不显示自己在线，下面增加代码
			if v == userID {
				continue
			}
			
			
			fmt.Println("用户ID=\t",v)
			//完成客户端的onlineUsers的初始化
			user := &message.User {
				UserID : v,
				UserStatus : message.UserOnline,
			}
			onlineUsers[v]= user
		}
		fmt.Print("\n\n")

		//这里我们还需要在客户端启动一个协程
		//该协程保持和服务器端的通讯，如果有服务器有数据推送
		//则接受并显示在客户端的终端
		go serverProcessMes(conn)
		
		
		//1.显示我们的登录成功的菜单[循环]
		for {
			ShowMenu()
		}
	} else {
		fmt.Println(loginResMes.Error)
	}
	
	return
}