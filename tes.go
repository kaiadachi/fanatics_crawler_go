package main

import (
	"fmt"
	"sync"
	"time"
)

func a(i int, wg *sync.WaitGroup, ch chan bool){
	fmt.Println(i)
	time.Sleep(5*time.Second)
	fmt.Println("ok")
	wg.Done()
	<-ch
}

func main() {
	ch := make(chan bool, 3)
	wg := &sync.WaitGroup{}
	for i := 0; i <=10; i++{
		ch <-true
		wg.Add(1)
		go a(i, wg, ch)
		time.Sleep(1*time.Millisecond)
	}

	wg.Wait()
	fmt.Println("fin")
}
