package process2
import (
	"fmt"
)

//因为UserMgr实例再服务器端有且只有一个
//因为再很多的地方都会使用到，因此弄成全局变量
var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

//完成对userMgr初始化工作
func init(){
	userMgr = &UserMgr {
		onlineUsers :make(map[int]*UserProcess,1024)
	}
}
//完成对onlineusers的添加
func (this *UserMgr)AddOnlineUser(up *UserProcess){
	this.onlineUsers[up.UserID] = up
}
//删除
func (this *UserMgr)DelOnlineUser(up *UserProcess){
	delete(this.onlineUsers,userID)
}
//返回当前所有在线用户
func (this *UserMgr)GetAllOnlineUser()map[int]*UserProcess{
	return this.onlineUsers
}

//根据ID返回一个对应的值
func(this *UserMgr)GetOnlineUserByID(userID int)(up *UserProcess, err error){
	//如何从map中取出一个值
	up,ok := this.onlineUsers[userID]
	if !ok {//说明你要查找的这个用户当前不在线
		err = fmt.Errorf("用户%d不存在",userID)
		return
	} 
	return
}
