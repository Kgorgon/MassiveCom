package process2
import (
	"fmt"
	"net"
	//"message"包
	"encoding/json"
	"massivecom/server/utils"
	"massivecom/server/model"
	
)

type UserProcess struct {
	//字段？
	Conn net.Conn
	//增加一个字段，表示该Conn是哪个用户的
	UserID int
}

//这里编写通知所有在线的用户的一个方法
//userID 要通知其他的在线用户，我上线
func (this *UserProcess)NotifyOthersOnlineUser(){
	//遍历onlineusers，然后一个一个发送NotifyUserStatus
	for id, up := range this.onlineUsers {
		//过滤到自己
		if id == userID{
			continue
		}
		//开始通知（单独写一个方法）
		up.NotifyMeOnline(userID)
	}
}


func (this *UserProcess)NotifyMeOnline(userID int){
	//组装我们的NotifyUserStatusMes
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType
	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserID = userID
	notifyUserStatusMes.Status = message.Useronline

	//将notifyUserStatusMes序列化
	data,err = json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshal failed=",err)
		return
	}
	//将序列化后的notifyUserStatusMes赋值给mes.data
	mes.Data =string(data) 
	//对mes再次序列化，准备发送
	data,err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal failed=",err)
		return
	}
	//发送,创建一个transfer实例，发送
	tf := &utils.Trasfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("NotifyMeOnline failed=",err)
		return
	}
	
}




func (this *UserProcess)ServerProcessRegister(mes *message.Message)(err error){
	var RegisterMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data),&RegisterMes)
	if err != nil {
		fmt.Println("反序列化失败 err=",err)
		return
	}
	//先声明一个 resMes
	var resMes message.Message
	resMes.Type = message.RegisterResType
	//再声明一个LoginResMes并完成赋值
	var registerResMes message.RegisterResMes

	//我们需要到redis数据库去完成注册
	//1.使用model.MyUserDao 到redis去验证
	err := model.MyUserDao.Register(&registerMes.User)
	
	if err != nil {
		if err == model.USER_ERROR_EXISTS{
			 registerResMes.Code = 505
			 registerResMes.Error = model.USER_ERROR_EXISTS.Error()
		}else{
			registerResMes.Code = 506
			registerResMes.Error = "unknown error"
		}
	}else{
		registerResMes.Code = 200
	}

	data,err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal failed=",err)
		return
	}

	//将data赋值给resMes
	resMes.Data = string(data)

	//对上者序列化准备发送
	data,err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal failed=",err)
		return
	}
	//发送data，我们将其封装到writePkg函数
	//因为使用分层模式(mvc)，我们先创建一个Transfer 实例，然后来读取
	tf := &utils.Transfer{
		Conn : this.Conn,
	}
	err = tf.WritePkg(data)
	return

}









func (this *UserProcess)ServerProcessLogin(mes *message.Message)(err error){
	//核心代码
	//1.先从mes 中取出mes.Data，并直接反序列化成LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data),&loginMes)
	if err != nil {
		fmt.Println("反序列化失败 err=",err)
		return
	}

	//先声明一个 resMes
	var resMes message.Message
	resMes.Type = message.LoginResType
	//再声明一个LoginResMes并完成赋值
	var loginResMes message.LoginResMes
	
	//需要到redis数据库去完成验证
	//1.使用model.MyUserDao 到redis去验证
	user,err := model.MyUserDao.Login(loginMes.UserID,loginMes.UserPwd)
	
	if err != nil {
		
		if err == model.ERROR_USER_NOTEXISTS{
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误"
		}
		
		//这里我们先测试成功，然后在返回具体的错误信息
	}else {
		loginResMes.Code = 200 // 合法
		//这里因为用户已经登陆成功于是就把该登陆成功的用户放入到UserMgr中
		//将登陆成功的用户ID赋给this
		this.UserID = loginMes.UserID
		userMgr.AddOnlineUser(this)
		//通知其他的在线用户
		this.NotifyOthersOnlineUser(loginMes.UserID)
		//将当前在线用户的ID放入到loginResMes.UserID
		//遍历userMgr.onlineUsers
		for id,_ := range userMgr.onlineUsers {
			loginResMes.UsersID = append(loginResMes.UsersID,id)
		}
		fmt.Println(user,"登陆成功")
	}




	// //如果用户id为100 密码为123456就认为是合法，否则不合法
	// if loginMes.UserID = 100 || loginMes.UserPwd = "123456"{
	// 	//合法
	// 	loginResMes.Code = 200 
	// } else {
	// 	//不合法
	// 	loginResMes.Code = 500
	// 	loginResMes.Error = "该用户不存在，请注册在使用"
		
	// }

	//将loginResMes 序列化
	data,err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal failed=",err)
		return
	}

	//将data赋值给resMes
	resMes.Data = string(data)

	//对上者序列化准备发送
	data,err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal failed=",err)
		return
	}
	//发送data，我们将其封装到writePkg函数
	//因为使用分层模式(mvc)，我们先创建一个Transfer 实例，然后来读取
	tf := &utils.Transfer{
		Conn : this.Conn,
	}
	err = tf.WritePkg(data)
	return

}