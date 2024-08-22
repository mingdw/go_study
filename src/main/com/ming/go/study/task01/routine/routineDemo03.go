package routine

import (
	"fmt"
	"sync"
	"time"
)

/*
协程简单调用，读写锁的实现
*/
var wg2 sync.WaitGroup
var lock2 sync.RWMutex

func RoutineMain3() {
	wg2.Add(6)
	for i := 1; i <= 5; i++ {
		go read()
	}
	go write()
	wg2.Wait()
}

func read() {
	defer wg2.Done()
	lock2.RLock()
	fmt.Println("读进程开始执行...")
	time.Sleep(time.Second)
	fmt.Println("读进程执行完毕...")
	lock2.RUnlock()
}

func write() {
	defer wg2.Done()
	lock2.Lock()
	fmt.Println("写进程开始执行...")
	time.Sleep(time.Second * 10)
	fmt.Println("写进程执行完毕...")
	lock2.Unlock()
}
