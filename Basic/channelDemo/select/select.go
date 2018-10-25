package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int{
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}
func main() {
	//var c1, c2  = make(chan int), make(chan int)
	var c1, c2  = generator(),generator()
	var worker = createWorker(0)

	var values []int
	tick := time.Tick(time.Second)
	tm := time.After(10 * time.Second)
	n := 0

	for{
		var activeWorker chan<- int
		var activeValue int

		if len(values) > 0{
			activeWorker = worker
			activeValue = values[0]

		}
		select {
		case n = <-c1:

			//fmt.Println("Received from c1:", n)
			values = append(values, n)
		case n = <-c2:

			values = append(values, n)
			//fmt.Println("Received from c2:", n)
		case activeWorker <- activeValue:
			values = values[1:]

		case <-time.After(800 * time.Millisecond):
			fmt.Println("timeout")
		case <-tick:
			fmt.Println("queue len =", len(values))
		case <-tm:
			fmt.Println("bye")
			return


		//default:
		//	fmt.Println("No value form c1,c2!")
		}
	}

}

func worker(id int, c chan int)  {
	for n := range c{
		time.Sleep(1 * time.Second)
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

func createWorker(id int) chan<- int{
	c := make(chan int)
	go worker(id ,c)
	return c

}






