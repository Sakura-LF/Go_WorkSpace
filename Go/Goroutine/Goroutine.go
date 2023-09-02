package Goroutine

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"log"
	"runtime"
	"sync"
	"time"
)

// 输出奇数的函数
//func PrintOdd() {
//	//并发执行结束后计数器-1
//	defer
//	for i := 1; i <= 10; i += 2 {
//		fmt.Println(i)
//		time.Sleep(100 * time.Millisecond)
//	}
//}
//
//// 输出偶数的函数
//func PrintEven() {
//	//并发执行结束后计数器-1
//	for i := 2; i <= 10; i += 2 {
//		fmt.Println(i)
//		time.Sleep(100 * time.Millisecond)
//	}
//}

func Goroutine() {
	//一: 初始化WaitGroup
	wg := sync.WaitGroup{}

	//二: 累加WG的计数器
	wg.Add(2)
	go func() {
		//并发执行结束后计数器-1
		defer wg.Done()
		for i := 1; i <= 10; i += 2 {
			fmt.Println(i)
			time.Sleep(100 * time.Millisecond)
		}
	}()
	go func() {
		//并发执行结束后计数器-1
		defer wg.Done()
		for i := 2; i <= 10; i += 2 {
			fmt.Println(i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	//主 goroutine等待
	wg.Wait()
	//执行到Wait会被阻塞,需要等到所有goroutine结束
	fmt.Println("其他goroutine结束")
}

func GoroutineRandom() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(n int) {
			defer wg.Done()
			fmt.Println(n)
		}(i)
	}
	wg.Wait()
}

func GoroutineNum() {
	//1.统计当前存在的Goroutine的数量
	go func() {
		for {
			fmt.Println("NumGoroutine:", runtime.NumGoroutine())
			//每个半秒钟输出Goroutine的数量
			time.Sleep(500 * time.Millisecond)
		}
	}()

	//2.启动大量的goroutine
	for {
		go func() {
			time.Sleep(100 * time.Second)
		}()
	}
}

func GoroutineAnts() {
	//1.统计Goroutine数量
	go func() {
		for {
			fmt.Println("Goroutine数量: ", runtime.NumGoroutine())
			time.Sleep(500 * time.Millisecond)
		}
	}()

	//2.初始化协程池 goroutine pool
	Size := 1024
	pool, err := ants.NewPool(Size)
	if err != nil {
		log.Fatalln(err)
	}
	//保证 pool 被关闭
	defer pool.Release()

	//3.利用goroutine pool ,调度需要并发的大量goroutine
	for {
		//想 pool 提交一个执行的goroutine
		pool.Submit(func() {
			v := make([]int, 1024)
			_ = v
			fmt.Println("in goroutine")
			time.Sleep(100 * time.Second)
		})
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func CPU() {
	fmt.Println(runtime.GOMAXPROCS(runtime.NumCPU()))
	fmt.Println(runtime.NumCPU())
}

func GoroutineSched() {

	wg := sync.WaitGroup{}
	wg.Add(2)

	//设置一个P在调度G
	runtime.GOMAXPROCS(1) //单线程模式

	// 输出奇数
	max := 100
	go func() {
		defer wg.Done()
		for i := 1; i < max; i += 2 {
			fmt.Print(i, " ")

			//主动让出
			runtime.Gosched()

			//增加执行时间
			//time.Sleep(1 * time.Millisecond)
		}
	}()

	//输出偶数
	go func() {
		defer wg.Done()
		for i := 2; i < max; i += 2 {
			fmt.Print(i, " ")

			// 主动让出
			runtime.Gosched()

			//增加执行时间
			//time.Sleep(1 * time.Millisecond)
		}
	}()

	wg.Wait()
}
