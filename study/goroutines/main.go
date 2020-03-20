package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 使用goroutine和channel实现一个计算int64随机数各位数和的程序。
// 1. 开启一个goroutine循环生成int64类型的随机数，发送到jobChan
// 2. 开启24个goroutine从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
// 3. 主goroutine从resultChan取出结果并打印到终端输出

type job struct {
	value int64
}

type result struct {
	job *job
	sum int64
}

var jobChan chan *job
var resChan chan *result

func generateInt(jb chan<- *job) {
	for {
		v := rand.Int63()
		j := &job{
			value: v,
		}
		jb <- j
		time.Sleep(time.Second)
	}
}

func workTask(jb <-chan *job, res chan<- *result) {
	for {
		j := <-jb
		sum := int64(0)
		n := j.value
		for n > 0 {
			sum += n % 10
			n = n / 10
		}

		r := &result{
			job: j,
			sum: sum,
		}
		res <- r
	}
}

func main() {
	jobChan = make(chan *job)
	resChan = make(chan *result)
	go generateInt(jobChan)
	for i := 0; i < 24; i++ {
		go workTask(jobChan, resChan)
	}
	for res := range resChan {
		fmt.Println(res.job.value, res.sum)
	}
}
