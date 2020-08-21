package process

import (
	"fmt"
	"os"
	"massivecom/client/utils"
	"net"
	//引入"message"包
	"encoding/json"
)

//显示登陆成功后的界面..
func ShowMenu() {
	fmt.Println("--------恭喜xxx登陆成功--------")
	fmt.Println("-------1.显示在线用户列表-------")
	fmt.Println("-------2.发送消息-------")
	fmt.Println("-------3.信息列表--------")
	fmt.Println("-------4.退出系统--------")
	fmt.Println("请选择1-4")
	var key int
	var content string
	
	//因为我们总要使用到SmsProcess实例，因此我们将其定义在switch外部
	smsProcess = &SmsProcess{}
	fmt.Scanf("%d\n",&key)
	switch key {
		case 1:
			//fmt.Println("显示在线用户列表~")
			outputOnlineUser()
		case 2:
			fmt.Println("相要对大伙儿说点啥？：")\
			fmt.Scanf("%s\n",&content)
			smsProcess.SendGroupMes(content)
		case 3:
			fmt.Println("信息列表")
		case 4:
			fmt.Println("你选择退出系统")
			os.Exit(0)
		default :
			fmt.Println("error!")
		
	}
}

//和服务器保持通讯
func serverProcessMes(conn net.Conn){
	//创建一个Transfer实例，不停地读取服务器发送的消息
	tf := &utils.Transfer{
		Conn : conn,
	}
	for {
		fmt.Println("客户端正在等待读取服务器发送的消息")
		mes,err := tf.ReadPkg()
		if err != nil{
			fmt.Pritnln("服务端error！")
		}
		//如果读取到消息，又是下一步处理逻辑
		switch mes.Type {
			case message.NotifyUserStatusMesType://有人上线
				//处理
				//1.取出NotifyUserStatusMes
				var notifyUserStatusMes message.NotifyUserStatusMes
				json.Unmarshal([]byte(mes.data),&notifyUserStatusMes)
				//2.把这个用户信息，状态保存到我们客户map[int]User中
				updataUserStatus(&notifyUserStatusMes)
			case message.SmsMesType ://有人群发消息
				outputGroupMes(&mes)
			default :
				fmt.Println("服务器端返回了一个未知的消息类型")
		}
		//fmt.Printf("mes=%v\n",mes)
	}
}