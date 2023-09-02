package TCP

import (
	"fmt"
	"log"
	"net"
	"time"
)

func Server() {
	// 要监听的地址和端口
	address := ":5678"

	//一:基于某个地址监听
	listener, err := net.Listen(tcp, address)
	if err != nil {
		log.Fatalln(err)
	}
	//defer 关闭监听
	defer listener.Close()
	fmt.Println("服务器正在监听:", listener.Addr())

	// 二:接收连接请求
	// 连接请求不是一个,所以要循环接收
	for true {
		//接受请求
		connection, err1 := listener.Accept()
		if err1 != nil {
			log.Fatalln(err1)
		}
		// 处理连接，读写
		go func(conn net.Conn) {
			// 日志连接的远程地址（client addr）
			log.Printf("accept from %s\n", conn.RemoteAddr())
			time.Sleep(time.Second)
		}(connection)
	}
}

func ServerWrite() {
	// 要监听的地址和端口
	address := ":5678"

	//一:基于某个地址监听
	listener, err := net.Listen(tcp, address)
	if err != nil {
		log.Fatalln(err)
	}
	//defer 关闭监听
	defer listener.Close()
	fmt.Println("服务器正在监听:", listener.Addr())

	// 二:接收连接请求
	// 连接请求不是一个,所以要循环接收
	for true {
		//接受请求
		connection, err1 := listener.Accept()
		if err1 != nil {
			log.Fatalln(err1)
		}
		// 调用处理每个连接的函数
		HandleConn(connection)
	}
}

// 处理每个连接
func HandleConn(conn net.Conn) {
	// 打印客户端地址
	fmt.Println("接收到了客户端连接,客户端地址为:", conn.RemoteAddr())

	// 1.保证连接关闭
	defer conn.Close()

	// 2.向客户端发送数据
	write, err := conn.Write([]byte("Server data" + "\n"))
	if err != nil {
		log.Println(err)
	}
	fmt.Println("服务端发送数据的长度 :", write)

	// 3.从客户端接收数据
	buf := make([]byte, 1024) // 构建 buffer 缓冲
	n, err2 := conn.Read(buf)
	if err2 != nil {
		fmt.Println("读取失败")
	}
	fmt.Println("读取客户端发送的内容为: ", string(buf[:n])) //截取前n个字节
}
