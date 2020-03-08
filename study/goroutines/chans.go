package goroutines

import (
	"fmt"
	"sync"
)

var chI1 chan int
var chI2 chan int
var wg sync.WaitGroup
var once sync.Once

func intGenerate(ch1 chan int) {
	defer wg.Done()
	for i := 1; i <= 100; i++ {
		ch1 <- i
	}
	close(ch1)
}

func intSquare(ch1, ch2 chan int) {
	defer wg.Done()
	for {
		v, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- v * v
	}
	once.Do(func() { close(ch2) })
}

func main1() {
	wg.Add(3)
	chI1 = make(chan int, 100)
	chI2 = make(chan int, 100)

	go intGenerate(chI1)
	go intSquare(chI1, chI2)
	go intSquare(chI1, chI2)
	wg.Wait()
	for v := range chI2 {
		fmt.Println(v)
	}
}
