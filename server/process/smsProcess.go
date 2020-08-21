package process2

import (
	"fmt"
	"encoding/json"
	//引入message包
	"net"
	"massivecom/server/utils"
)

type SmsProcess struct {
	//暂时不需要字段
}

//写方法转发消息
func (this *SmsProcess)SendGroupMes(mes *message.Message){
	//遍历服务器端的onlineUsers map[int]*UserProcess
	//将消息转发出去
	//取出mes的内容SmsMes

	var smsMes message.smsMes
	err := json.Unmarshal([]byte(mes.data),&smsMes)
	if err != nil {
		fmt.Println("反序列化failed=",err)
		return
	}


	data,err := json.Marshal(mes)
	if err != nil {
		fmt.Println("序列化failed=",err)
		return
	}


	for id, up := range userMgr.onlineUsers{
		//这里还需要过滤掉自己，不要再发给自己
		if id == smsMes.UserID {
			continue
		}
		this.SendMesToEachOnlineUser(data,up.Conn)
	}
}



func (this *SmsProcess)SendMesToEachOnlineUser(data []byte,conn net.Conn){
	//创建一个Transfer实例，发送data
	tf:= &utils.Transfer{
		Conn : conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("转发消息failed=",err)
	}
} 