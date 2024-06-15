package main

import "fmt"

func main() {
	gorutinenum := 5
	var chanslice []chan struct{}
	exitchan := make(chan struct{})

	for i := 0; i < gorutinenum; i++ {
		chanslice = append(chanslice, make(chan struct{}))
	}

	res := 1
	next := 0
	for i := 0; i < gorutinenum; i++ {
		go func(i int) {
			for {
				<-chanslice[i]
				if res > 100 {
					exitchan <- struct{}{}
					break
				}

				fmt.Printf("gorutine %d: %d.\n", i, res)
				res++

				if i == gorutinenum-1 {
					next = 0
				} else {
					next = i + 1
				}
				chanslice[next] <- struct{}{}
			}
		}(i)
	}
	chanslice[0] <- struct{}{}

	select {
	case <-exitchan:
		fmt.Println("end")
	}
}
