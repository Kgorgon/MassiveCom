package process

import (
	"fmt"
	"encoding/json"
	//"引入message包"
	"massivecom/client/utils"
)


type SmsProcess struct {

}

//发送群聊的消息
func (this *SmsProcess)SendGroupMes(content string)(err error){
	//1.创建一个Mes
	var mes message.Message
	mes.Type = message.SmsMesType
	//2.创建一个SmsMes 实例
	var smsMes message.SmsMes
	smsMes.Content = content //内容
	smsMes.UserID = CurUser.UserID //
	smsMes.UserStatus = CurUser.UserStatus//
	//3.序列化smsMes
	data, err := json.Marshal(smsMes)
	if err != {
		fmt.Println("sms序列化failed=",err)
		return
	}
	mes.Data = string(data)
	//4.对mes再次序列化
	data, err = json.Marshal(mes)
	if err != {
		fmt.Println("mes序列化failed=",err)
		return
	}

	//5.将序列化后的mes发送给服务器
	tf := &utils.Trasfer {
		Conn : CurUser.Conn,
	}
	//6.发送
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("sms failed=",err)
		return
	}
	return
}