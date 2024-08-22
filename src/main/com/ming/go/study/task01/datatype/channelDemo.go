package datatype

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 管道相关的练习
func ChannelMain() {
	//定义管道
	var schannel chan string
	//初始化长度
	schannel = make(chan string, 3)

	fmt.Println("管道的类型: ", schannel)

	//管道存入数据
	schannel <- "a"
	schannel <- "b"
	schannel <- "c"
	fmt.Println("当前管道的实际大小： ", len(schannel), "; 管道容量是: ", cap(schannel))
	//如果管道数量已满，在存入会报错
	//schannel <- "d" 报错

	//管道中读取数据
	str001 := <-schannel
	fmt.Println("管道读取的数据： ", str001)
	fmt.Println("管道数据剩余数据: ", len(schannel))

	var schannel2 chan string
	//初始化长度
	schannel2 = make(chan string, 100)

	for i := 0; i < 50; i++ {
		schannel2 <- strconv.Itoa(i)
	}

	//管道关闭后写入报错,读不受影响
	close(schannel2)
	//schannel2 <- "asd"
	fmt.Println("管道关闭读取元素: ", <-schannel2)

	//管道的遍历
	for i := 0; i < 20; i++ {
		fmt.Println("管道遍历", strconv.Itoa(i), ": ", <-schannel2)
	}
	//如果for range遍历，需要close管道

	for v := range schannel2 {
		fmt.Println("管道遍历", v)
	}

	wg.Add(2)
	var strch chan string
	strch = make(chan string, 50)
	go writeCh(strch)
	go readCh(strch)
	wg.Wait()
}

func writeCh(ch chan string) {
	defer wg.Done()
	for i := 0; i < 50; i++ {
		ch <- strconv.Itoa(i)
		fmt.Println("写入的数据是: ", strconv.Itoa(i))
		time.Sleep(time.Second)
	}
	close(ch)
}

func readCh(ch chan string) {
	defer wg.Done()
	for v := range ch {
		fmt.Println("读取的数据是: ", v)
		time.Sleep(time.Second)
	}
}
