package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}

	charFlag := make(chan struct{})
	numFlag := make(chan struct{})
	wg.Add(1)

	go func() {
		i := 1
		for {
			<-numFlag
			fmt.Printf("%d%d", i, i+1)
			i += 2
			charFlag <- struct{}{}
		}
	}()

	go func(wg *sync.WaitGroup) {
		ch := 'A'
		for {
			<-charFlag
			if ch > 'Z' {
				wg.Done()
				return
			}
			str := string(ch) + string(ch+1)
			fmt.Printf("%s", str)
			ch += 2
			numFlag <- struct{}{}
		}
	}(wg)

	numFlag <- struct{}{}
	wg.Wait()
}
