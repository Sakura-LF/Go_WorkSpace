package TCP

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

func Client() {
	//服务器地址
	ServerAddress := "127.0.0.1:5678"

	//创造多个客户端的连接请求
	num := 1000
	wg := sync.WaitGroup{}
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func() {
			defer wg.Done()
			//一: 建立连接
			connection, err := net.DialTimeout(tcp, ServerAddress, 1*time.Millisecond)
			if err != nil {
				log.Fatalln(err)
			}
			//defer 关闭连接
			defer connection.Close()

			//处理连接
			fmt.Println("连接被接受,服务器地址为", connection.LocalAddr())
		}()

		//10毫秒一个请求
		time.Sleep(10 * time.Millisecond)
	}
	wg.Wait()
}

func ClientRead() {
	//服务器地址
	ServerAddress := "127.0.0.1:5678"

	// 一: 建立连接
	connection, err := net.DialTimeout(tcp, ServerAddress, 1*time.Millisecond)
	if err != nil {
		log.Fatalln("客户端连接失败")
	}

	// 二:保证关闭
	defer connection.Close()
	fmt.Println("连接完成,服务器地址为:", connection.LocalAddr())

	// 三:读取服务端发送的数据
	buf := make([]byte, 1024) // 构建 buffer 缓冲
	n, err2 := connection.Read(buf)
	if err2 != nil {
		fmt.Println("读取失败")
	}
	fmt.Println("读取服务端发送的内容为: ", string(buf[:n])) //截取前n个字节

	// 四:向服务端发送数据
	write, err := connection.Write([]byte("Client date" + "\n"))
	if err != nil {
		log.Println(err)
	}
	fmt.Println("客户端发送数据长度 :", write)
}
