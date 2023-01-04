package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	ch := make(chan struct{}, 1)
	go func() {
		ch <- struct{}{}
		fmt.Println("出现")
		<-ch
		wg.Done()
	}()
	wg.Wait()
}
