package utils

import (
	"fmt"
	"net"
	//"message"包
	"encoding/binary"
	"encoding/json"
)
//这里将这些方法关联到结构体

type Transfer struct {
	//分析它应该有那些字段
	Conn net.Conn
	Buf [8096]byte //这是传输时，使用的缓冲

}


func (this *Transfer)ReadPkg()(mes message.Message,err error){
	fmt.Println("读取客户端发送的数据..")
	//conn.Read 在conn没有被关闭的情况下，才会阻塞
	//如果客户端关闭了conn，则不会阻塞
	_,err = this.Conn.Read(this.Buf[:4])
	if  err != nil {
		//fmt.Println("conn.Read failed=",err)
		//err = erros.New("read pkg header error")
		return
	}
	//根据buf[:4] 转成一个uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[0:4])
	// pkgLen 读取消息内容
	n,err := this.Conn.Read(this.Buf[:pkgLen])
	if n!= int(pkgLen) || err != nil {
		//err = errors.New("read pkg body error")
		return
	}
	//把pkgLen 反序列化成 -> message.Message
	//技术就是一层窗户纸
	err = json.Unmarshal(this.Buf[:pkgLen],&mes)
	if err != nil {
		fmt.Println("json.Unmarshal failed=",err)
		return
	}
	return
}




func (this *Transfer) WritePkg(data []byte)(err error){

	//先发送一个长度给客户端
	var pkgLen uint32
	pkg = uint32(len(data)) 
	binary.BigEndian.PutUint32(this.Buf[0:4],pkgLen)
	//发送长度
	n,err := this.Conn.Write(this.Buf[:4])
	if n!=4 || err !=nil {
		fmt.Println("conn.Write(bytes) failed= ",err)
		return
	}

	//发送data消息本身
	n,err := this.Conn.Write(data)
	if n!=int(pkgLen) || err !=nil {
		fmt.Println("conn.Write(data) failed= ",err)
		return
	}
}