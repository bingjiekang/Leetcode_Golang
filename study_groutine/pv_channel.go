package studygroutine

import (
	"fmt"
)

// 我们在循环体首尾都使用了阻塞独占模式，两个chan交替释放控制权，达到了安全的协程交互控制。
func Pv_channel() {
	var ch1, ch2 chan int = make(chan int), make(chan int)
	var exit chan bool = make(chan bool)

	// 生产者
	go func() {
		for i := 0; i < 50; i++ {
			// 重点
			ch1 <- i
			fmt.Println("协程1输出:A")
			<-ch2
		}
		exit <- true
	}()

	// 消费者
	go func() {
		for i := 0; i < 50; i++ {
			// 重点
			<-ch1
			fmt.Println("协程2输出:B")
			ch2 <- i
		}
	}()
	if <-exit {
		fmt.Println("携程结束,总进程结束....ending")
	}
}
