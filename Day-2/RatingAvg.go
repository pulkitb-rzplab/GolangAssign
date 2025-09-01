package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func rate(ch chan<- int) {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	t := r.Intn(5)

	time.Sleep(time.Duration(t) * time.Second)
	x := r.Intn(6)
	ch <- x
	fmt.Println(x, "   time taken: ", t)

}

func main() {
	ch := make(chan int, 1)
	var wg sync.WaitGroup

	for range 200 {
		wg.Go(func() {
			rate(ch)
		})
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	sum := 0
	for i := range ch {
		sum += i

	}
	avg := float32(sum / 200)

	fmt.Println("The average rating is: ", avg)

}
