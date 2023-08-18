package studygroutine

import "fmt"

// 有缓冲的生产者消费者,协程交替打印
func Pv_channel_cd() {
	var exit chan bool = make(chan bool)
	var ch1, ch2 chan bool = make(chan bool, 1), make(chan bool)

	// 启动,先写入
	ch1 <- true
	// 协程1
	go func() {
		for i := 0; i < 10; i++ {
			if <-ch1 {
				fmt.Println("协程1输出A")
				ch2 <- true
			}
		}

	}()

	// 携程2
	go func() {
		defer close(exit)
		for i := 0; i < 10; i++ {
			if <-ch2 {
				fmt.Println("协程2输出B")
				ch1 <- true
			}
		}

	}()

	tm, ok := <-exit
	fmt.Println("程序结束....ending", tm, ok)
}
