package Channel

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func ChannelOperate() {
	// 初始化
	ch := make(chan int) //无缓冲channel

	//发送值
	go func() {
		ch <- 100
	}()

	//接收值
	go func() {
		value := <-ch
		fmt.Println("receive from ch : ", value)
	}()
	time.Sleep(time.Second)

	//关闭ch
	close(ch)

}

func ChannelFor() {
	// 一，初始化部分数据
	ch := make(chan int) // 无缓冲的channel
	wg := sync.WaitGroup{}

	// 二，持续发送
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			// random send value
			ch <- rand.Intn(10)
		}
		// 关闭
		//close(ch)
	}()

	// 三，持续接收，for range
	wg.Add(1)
	go func() {
		defer wg.Done()
		// 持续接收
		for e := range ch {
			println("received from ch, element is ", e)
		}
	}()

	wg.Wait()
}

// 通过比对时间来判断接受和发送的时间是否是相互的
func NoBufferChannel() {
	//一: 创建无缓冲channel
	ch := make(chan int)
	wg := sync.WaitGroup{}

	//二: 间隔发送
	wg.Add(1)
	go func() {
		wg.Done()
		for i := 0; i < 5; i++ {
			ch <- i
			// 格式化时间输出
			fmt.Println("发送数据: ", i, " 时间:", time.Now().Format("15:04:05.999999"))
			//间隔时间
			time.Sleep(1 * time.Second)
		}
		//关闭 channel
		close(ch)
	}()

	//三: 间隔接收
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range ch {
			fmt.Println("接收数据: ", i, " 时间:", time.Now().Format("15:04:05.999999"))

			// 间隔时间
			time.Sleep(3 * time.Second)
		}
	}()

	wg.Wait()
}

func BufferChannel() {
	//一: 创建缓冲channel
	ch := make(chan int, 5)
	wg := sync.WaitGroup{}

	//二: 间隔发送
	wg.Add(1)
	go func() {
		wg.Done()
		for i := 0; i < 5; i++ {
			ch <- i
			// 格式化时间输出
			fmt.Println("发送数据: ", i, " 时间:", time.Now().Format("15:04:05.999999"), "Len: ", len(ch), "Cap: ", cap(ch))
			//间隔时间
			time.Sleep(1 * time.Second)
		}
		//关闭 channel
		close(ch)
	}()

	//三: 间隔接收
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range ch {
			fmt.Println("接收数据: ", i, " 时间:", time.Now().Format("15:04:05.999999"), "Len: ", len(ch), "Cap: ", cap(ch))

			// 间隔时间
			time.Sleep(3 * time.Second)
		}
	}()

	wg.Wait()
}

func ChannelGroutineNumCtl() {
	// 1.独立的 goroutine 输出 goroutine数量
	go func() {
		for {
			fmt.Println("NumGroutine", runtime.NumGoroutine())
			time.Sleep(500 * time.Millisecond)
		}
	}()

	// 2.初始化channel,设置缓存区大小(并发规模)
	const size = 1024
	ch := make(chan struct{}, size)

	// 3.并发的goroutine
	for {
		//一: 启动goroutine前,执行 ch send
		// 当ch的缓冲已满时,阻塞
		ch <- struct{}{}
		go func() {
			time.Sleep(10 * time.Second)
			//二: goroutine结束时,接收一个ch中的元素
			<-ch
		}()
	}
}

func ChannelDirection() {
	// 一: 初始化数据
	ch := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	// 使用双向channel为单向channel赋值
	go getElement(ch, wg)
	go setElement(ch, 45, wg)

	wg.Wait()

}

// 仅接收的channel
func getElement(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("接收的数据为:", <-ch)
}

// 仅发送的channel
func setElement(ch chan<- int, value int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- value

	fmt.Println("发送的数据为:", value)
}

func SelectTest() {
	//1.声明需要的变量
	var array [3]int
	var c1, c2, c3, c4 = make(chan int), make(chan int), make(chan int), make(chan int)
	var i1, i2 int

	//2.用于select的goroutine
	go func() {
		c1 <- 10
	}()
	go func() {
		<-c2
	}()
	go func() {
		close(c3)
	}()
	go func() {
		c4 <- 40
	}()

	//3.多路监听的select
	go func() {
		select {
		//监听是否可以从c1中获取值
		case i1 = <-c1:
			fmt.Println("接收到了C1的值:", i1)
		//监听对c2的写操作
		case c2 <- i2:
			fmt.Println("向c2写入了:", i2)
		//c3是否被关闭
		case i3, ok := <-c3:
			if ok {
				fmt.Println("从接收到了:", i3)
			} else {
				fmt.Println("c3已被关闭")
			}
		//测试左值表达式的执行实际
		case array[f()] = <-c4:
			fmt.Println("从c4收到", array[f()])
		//默认
		default:
			fmt.Println("没有channel被操作")
		}
	}()

	//简单阻塞
	time.Sleep(500 * time.Millisecond)
}

func f() int {
	fmt.Println("f()执行")
	return 2
}

func ForSelect() {
	// 定义channel
	ch := make(chan int)

	// 向 channel 发送数据
	go func() {
		//实际工作中,ch可能来自缓存,网络,数据库
		for {
			ch <- rand.Intn(1000)
			time.Sleep(200 * time.Millisecond)
		}
	}()

	// 持续监听 channel
	go func() {
		for {
			select {
			case value := <-ch:
				fmt.Println("接收到了来之channel的数据:", value)
			}
		}
	}()

	time.Sleep(3000 * time.Millisecond)
}

func BlockChannel() {
	// 空select阻塞
	fmt.Println("before select")
	//select {}
	//fmt.Println("after select")

	//nil channel
	var channel chan int //nil channel 不能读写
	go func() {
		channel <- rand.Intn(1000)
	}()
	fmt.Println("before select")
	select {
	case <-channel:
	case channel <- 100:
	}
	fmt.Println("after select")
}

func NilChannel() {

	// 一:初始化channel
	channel := make(chan int)

	// 二:操作 channel 的 goroutine
	go func() {
		//向 channel 中写入随机数
		for {
			channel <- rand.Intn(1000)
			time.Sleep(400 * time.Millisecond)
		}
	}()

	// 三:select 处理内部的 goroutine
	go func() {
		// 设置定时器
		after := time.After(3 * time.Second)
		sum := 0
		// 持续监听 channel
		for {
			select {
			case value := <-channel:
				fmt.Println("接收到channel的数据:", value)
				sum++
			// 定时器时间到了之后,就可以从里面读出内容
			case <-after:
				// 将channel设置为nil channel
				channel = nil
				fmt.Println("channel已设置为nil,sum为:", sum)
			}
		}
	}()
	// 让主 goroutine 阻塞
	time.Sleep(5 * time.Second)
}

func GuessChannel() {
	// 初始化数据
	People := 10             // 参与人数
	max := 10                // 猜数最大范围
	answer := rand.Intn(max) // 正确答案
	wg := sync.WaitGroup{}

	fmt.Println("正确答案是:", answer)
	fmt.Println("----------------------------------------")

	// 存有正确答案的channel
	channel := make(chan int, People)

	wg.Add(People)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			//每个人随机猜一个数
			result := rand.Intn(max)
			fmt.Println("参赛者猜数是:", result)
			//如果猜正确了,就直接发到channel里面
			if answer == result {
				channel <- result
			}
		}()
	}
	wg.Wait()

	fmt.Println("-----------------------------")
	select {
	case value := <-channel:
		fmt.Println("有人猜中了value:", value)
	default:
		fmt.Println("没有任何人猜到数据")
	}

}

func SelectNonBlock() {
	// 初始化数据
	counter := 10 // 参与人数
	max := 20     // [0, 19] // 最大范围
	rand.Seed(time.Now().UnixMilli())
	answer := rand.Intn(max) // 随机答案
	println("The answer is ", answer)
	println("------------------------------")

	// 正确答案channel
	bingoCh := make(chan int, counter)
	// wg
	wg := sync.WaitGroup{}
	wg.Add(counter)
	for i := 0; i < counter; i++ {
		// 每个goroutine代表一个猜数字的人
		go func() {
			defer wg.Done()
			result := rand.Intn(max)
			println("someone guess ", result)
			// 答案争取，写入channel
			if result == answer {
				bingoCh <- result
			}
		}()
	}
	wg.Wait()

	println("------------------------------")
	// 是否有人发送了正确结果
	// 可以是0或多个人
	// 核心问题是是否有人猜中，而不是几个人
	select {
	case result := <-bingoCh:
		println("some one hint the answer ", result)
	default:
		println("no one hint the answer")
	}
}

func RaceTest() {
	// 初始化数据
	type Rows struct {
		// 数据字段

		// 索引标识
		Index int
	}

	// 定义查询器数量
	const QueryNum = 8

	// 数据通信channel 和 停止信号channel
	ResultChannel := make(chan Rows, 1)
	StopChannel := [QueryNum]chan struct{}{}
	for key := range StopChannel {
		StopChannel[key] = make(chan struct{})
	}

	//wg,rand
	wg := sync.WaitGroup{}
	//rand1 := rand.Intn(100)

	// 模拟查询 , 等待第一个结果反馈
	// 基于定义的查询器的数量,开启goroutine
	wg.Add(QueryNum)
	for i := 0; i < QueryNum; i++ {
		// 开启一个查询器
		go func(i int) {
			defer wg.Done()
			// 模拟执行时间
			RandTime := rand.Intn(1000)
			fmt.Println("查询器", i, "获取数据,需要 ", RandTime, " ms 执行")

			// 每个goroutine的结果channel
			chrst := make(chan Rows, 1)

			// 模拟执行查询
			go func() {
				//模拟时长
				time.Sleep(time.Duration(RandTime) * time.Millisecond)
				chrst <- Rows{
					Index: i,
				}
			}()

			// 监听查询结果和停止信号channel
			select {
			case result := <-chrst:
				fmt.Println("查询器 ", i, " 获得数据")
				// 保证没有其他结果写入
				if len(ResultChannel) == 0 {
					ResultChannel <- result
				}

			//停止信号
			case <-StopChannel[i]:
				fmt.Println("查询器 ", i, " 停止")
				return
			}
		}(i)
	}

	// 三:等待第一个查询结果的反馈
	wg.Add(1)
	go func() {
		defer wg.Done()
		//等待ch中传递的结果
		select {
		case value := <-ResultChannel:
			fmt.Println("得到了结果", value.Index, "停止其他查询")
			// 循环遍历,通知其他查询器结束
			for key := range StopChannel {
				if key == value.Index {
					continue
				}
				StopChannel[key] <- struct{}{}
			}
		//计划一个超时时间
		case <-time.After(5 * time.Second):
			fmt.Println("所有查询器停止")
			for key := range StopChannel {
				StopChannel[key] <- struct{}{}
			}
		}
	}()

	wg.Wait()
}
