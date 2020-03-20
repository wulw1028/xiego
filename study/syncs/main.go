package main

import (
	"sync"
	"time"
	"fmt"
)

var(
	x int = 0
	wg sync.WaitGroup
	lock sync.Mutex
	rwLock sync.RWMutex
)

func read(){
	rwLock.RLock()
	fmt.Println(x)
	time.Sleep(time.Millisecond * 5)
	rwLock.RUnlock()
	wg.Done()
}

func write(){
	rwLock.Lock()
	x = x +1
	time.Sleep(time.Millisecond * 10)
	rwLock.Unlock()
	wg.Done()
}

func main()  {
	n := time.Now()
	for i:=0; i<10;i++{
		wg.Add(1)
		go write()
	}
	
	for i:=0;i<100;i++{
		wg.Add(1)
		go read()
	}

	wg.Wait()
	fmt.Println(time.Now().Sub(n))
}