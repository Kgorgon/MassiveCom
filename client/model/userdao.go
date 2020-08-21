package model

import (
	"fmt"
	"github.com/redis"
	"encoding/json"
	//"引入message包"
)

//定义一个UserDao结构体
//完成对User结构体的各种操作：添加等

//我们在服务器启动后就初始化一个userDao的实例
//把它做成全局的，在需要和redis操作时，就直接使用即可

var (
	MyUserDao *UserDao
)





type UserDao struct {
	pool *redis.Pool
}

//使用工厂模式，创建一个UserDao的实例
func NewUserDao(pool *redis.Pool)(userDao *UserDao){
	userDao = &UserDao{
		pool,pool,
	}
	return
}




//UserDao应该提供哪些方法？
//1.根据一个用户ID返回一个User实例+error
func (this *UserDao)getUserById(conn redis.Conn, id int)(user *User,err error){

	//通过给定id 去 redis查询这个用户
	res,err := redis.String(conn.Do("HGet","users",id))
	if err != nil {
		//错误!
		if err == redis.ErrNil{//表示在users哈希中没有找到对应的ID
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	
	user = &User {}
	//这里我们需要把res反序列化成User实例
	err = json.Unmarshal([]byte(res),user)
	if err != nil{
		fmt.Println("反序列化 failed=",err)
		return
	}
	return
}



//完成登录的校验
//1. Login 完成对用户的验证
//2. 如果用户的id和pwd都正确，则返回一个user实例
//3. 如果有用户id pwd有错误，则返回一个对应的错误信息
func (this *UserDao) Login(userID int,UserPwd string)(user *User,err error){
	//先从userdao的连接池取出一根连接
	conn := this.pool.Get()
	defer conn.Close()
	user,err = this.getUserById(conn,userID)
	if err != nil {
		return
	}
	//这是证明这个用户是获取到
	if user.UserPWd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}


func (this *UserDao) Register(user *message.User)(err error){
	//先从userdao的连接池取出一根连接
	conn := this.pool.Get()
	defer conn.Close()
	user,err = this.getUserById(conn,user.UserID)
	if err == nil {
		err = ERROR_USER_EXISTS
	}
	//这时说明id不在redis里面，则可以完成注册
	data,err := json.Marshal(user)//序列化
	if err != nil {
		return
	}
	//入库


	_,err = conn.Do("Hset","users",user.UserID,string(data))
	if err != nil {
		fmt.Println("保存注册用户failed=",err)
		return
	}
	return

}