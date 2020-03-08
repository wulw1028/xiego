package goroutines

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func waitGroup(i int) {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Second * time.Duration(rand.Intn(3)))
	fmt.Println(i)
}

func runmain() {
	runtime.GOMAXPROCS(2)
	fmt.Println("cpu", runtime.NumCPU())
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go waitGroup(i)
	}
	wg.Wait()
}
