package routine

import (
	"fmt"
	"sync"
)

/*
*
协程简单调用，利用WaiteGroup携程启动多个任务
*/

var wg sync.WaitGroup

var a int = 0

// 互斥锁的实现
var lock sync.Mutex

func RoutineMain2() {
	//由于没有加互斥锁，所以最终结果不为0
	fmt.Println("********************")
	wg.Add(2)
	go add()
	go sub()
	wg.Wait()
	fmt.Println("result a =: ", a)
}

func add() {
	lock.Lock()
	for i := 0; i < 10000; i++ {
		a++
	}
	lock.Unlock()
	wg.Done()
}

func sub() {
	lock.Lock()
	for i := 0; i < 10000; i++ {
		a--
	}
	lock.Unlock()
	wg.Done()
}
