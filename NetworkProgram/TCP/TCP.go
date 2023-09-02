package TCP

import (
	"log"
	"net"
	"sync"
	"time"
)

// 因为Listen中的network支持不同的网络类型
// 所以要定义一个常量的tcp
const tcp = "tcp"

// 服务端
func TcpServer() {
	// A:基于某个地址寄哪里监听
	// 服务端地址
	//address := "127.0.0.1:5678" //IPV4
	address := ":5678"
	listener, err := net.Listen(tcp, address)
	if err != nil {
		log.Fatalln(err)
	}
	// 关闭监听操作
	defer listener.Close()
	log.Printf("%s server is listening on %s\n", tcp, listener.Addr())

	//B:接受连接请求
	//不止一个Client请求,所以要循环接收请求
	for {
		// 阻塞接收当前请求
		connection, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}

		// 处理连接,读写操作
		// 连接的远程地址 (client addr)
		log.Printf("accept from %s\n", connection.RemoteAddr())
	}
}

func TcpClient() {
	// tcp服务器端地址
	address := "127.0.0.1:5678"
	// 模拟多客户端
	// 并发的客户端请求
	num := 10
	wg := sync.WaitGroup{}
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			//一.建立连接
			connection, err := net.DialTimeout(tcp, address, 1*time.Millisecond)
			if err != nil {
				log.Fatalln(err)
				return
			}
			// 保证关闭
			defer connection.Close()
			log.Printf("connection is establish ,client addr is %s\n", connection.LocalAddr())
		}(&wg)
	}
	wg.Wait()
}
