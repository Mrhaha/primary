package primary

import (
	"fmt"
	"net"
)

func StartServer() {
	listen,err := net.Listen("tcp","localhost:8888")
	if err != nil {
		fmt.Printf("the listen is err :%v",err)
		return
	}
	conn,err := listen.Accept()
	if err != nil {
		fmt.Println("the accept is err")
		return
	}
	buf := make([]byte,1024)
	n,err1 := conn.Read(buf)
	if err1 != nil {
		fmt.Println("err = ",err)
		return
	}
	fmt.Println("buf = ",string(buf[:n]))
	conn.Close()
	listen.Close()
}