package routine

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

/*
两个协程，一个写数据一个读数据
*/

// 默认情况下，管道为双向的，可读可写
var rwChan chan string
var wg3 sync.WaitGroup
var lock3 sync.RWMutex

func ChannelReadAndWriteMain() {
	/*wg3.Add(2)
	go write()
	go read()
	wg3.Wait()*/

	syncChan()
}

func writeChan(str string) {
	defer wg3.Done()
	fmt.Println("start write chan!")
	for i := 0; i < 50; i++ {
		lock3.Lock()
		rwChan <- str
		lock3.Unlock()
		//time.Sleep(time.Second)

	}
	fmt.Println("end write chan!")
}

func readChan(str string) {
	defer wg3.Done()
	fmt.Println("start write chan!")
	for i := 0; i < 50; i++ {
		lock3.RUnlock()
		fmt.Println("读取管道的值为：", <-rwChan)
		lock3.RUnlock()
		time.Sleep(time.Second)
	}
	fmt.Println("end write chan!")
}

// 管道的阻塞案例
func syncChan() {
	var chStr chan string
	chStr = make(chan string, 100)
	wg3.Add(2)
	go syncWrite(chStr)
	go syncReade(chStr)
	wg3.Wait()

}

func syncWrite(chStr chan string) {
	defer wg3.Done()
	for i := 0; i < 100; i++ {
		chStr <- strconv.Itoa(i)
		fmt.Println("写入数据： ", i)
		time.Sleep(time.Second)
	}
}
func syncReade(chStr chan string) {
	defer wg3.Done()
	for {
		fmt.Println("读取数据： ", <-chStr)
	}
}
