package process

import (
	"fmt"
	"encoding/json"
	//引入message包
)

func outputGroupMes(mes *message.Message){
	//显示即可
	//1.反序列化me
	var smsMes messgae.SmsMes
	err := json.Unmarshal([]byte(mes.beta),&smsMes)
	if err != nil {
		fmt.Println("反序列化failed=",err)
		return
	}
	//显示信息
	info := fmt.Sprintf("用户ID=\t%d 对大家说:\t%s",smsMes.UserID,smsMes.Content)
	fmt.Println(info)
	fmt.Println()
		
	

}