package main

import (
	"fmt"
	"net"
	//"message包"
	"io"
	"massivecom/server/utils"
	"massivecom/server/process"

)
//先创建Processor 的结构体
type Processor struct {
	Conn net.Conn
}
func (this *Processor)serverProcessMes(mes *message.Message)(err error){
	//看看是否能就受到客户端的群发消息
	fmt.Println("mes=",mes)
	
	switch  mes.Type{
		case message.LoginMesType:
			//处理登录
			//创建一个UserProcess实例
			up := &process2.Userprocess {
				Conn : this.Conn,
			}
			err = up.ServerProcessLogin(mes)
		case message.RegisterMesType:
			//处理注册
			up := &process2.Userprocess {
				Conn : this.Conn,
			}
			err = up.ServerProcessRegister(mes)
		case message.SmsMesType :
			//创建一个SmsProcess实例完成转发群聊消息
			smsProcess := &process2.SmsProcess{}
			smsProcess.SendGroupMes(mes)
		default:
			fmt.Println("消息类型不存在")
		
	}
	return
}


func (this *Processor)process2()(err error){
	//读客户端发送的信息
	for {
		//这里我们将读取数据包，直接封装成一个函数readPkg()，返回Message,Err
		
		//创建Transfer实例，完成读包任务
		tf := &utils.Transfer {
			Conn : this.Conn,
		}
		mes,err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器也退出")
				return err
			} else {
				fmt.Println("readPkg err=",err)
				return err
			}
		}

		//fmt.Println("mes=",mes)
		err = this.serverProcessLogin(&mes)
		if err != nil {
			return err
		}
	}
}

