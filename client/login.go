package main

import (
	"fmt"
	"encoding/json"
	"encoding/binary"
	"net"
	//记得引入message包
)
//写一个函数，完成登录
func login(userID int, userPwd string)(err error){
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
	//fmt.Println("客户端，发送消息的长度=%d 内容是 =%s",len(data),string(data))
	
	
	//发送消息本身
	_,err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) failed= ",err)
		return
	}
	//这里还需要处理服务器端返回的消息
	mes,err = readPkg(conn)//mes 就是
	if err != nil{
		fmt.Println("readPkg(conn) failed=",err)
	}
	
	
	//将mes的data部分反序列化成LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data),&loginResMes)
	if loginResMes.Code == 200{
		fmt.Println("用户登陆成功")
	}else if loginResMes.Code == 500 {
		fmt.Println("用户不存在请注册")
	}
	
	return
}