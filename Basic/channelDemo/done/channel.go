package main

import (
	"fmt"
	"sync"
	"time"
)

func createWorker(id int,wg *sync.WaitGroup) worker{
	w := worker{
		in : make(chan int),
		done: func() {
			wg.Done()
		},
	}

	go doWorker(id, w.in,w.done)
	return w
}

func doWorker(id int,c chan int, done func()){
	for n := range c{

		fmt.Printf("Worker %d received %c\n",id,n)
		done()
		//go func() {
		//
		//
		//}()
	}

}
type worker struct {
	in chan int
	done func()
}
func chanDemo(){
	var wg sync.WaitGroup
	var workers [10]worker
	for i:=0;i<10;i++{
		workers[i] = createWorker(i,&wg)
	}

	// wait for all of them
	//go func() {
	//		for _,worker :=range workers{
	//			<- worker.done
	//			<- worker.done
	//		}
	//}()

	// 使用WatiGroup来等待goroutine的结束

	wg.Add(20)
	for i := 0; i< 10;i++{
		workers[i].in <- 'a' + i
	}

	for i := 0; i< 10;i++{
		workers[i].in <- 'A' + i
	}
	wg.Wait()

	//time.Sleep(time.Microsecond)
	//fmt.Println(c)
}

func bufferedChanDemo(){
	c := make(chan int,5)


	c <- 2
	c <- 3
	c <- 4
	c <- 5
	c <- 6
	//c <- 7
	go func() {
		for{
			fmt.Println("b")
			fmt.Printf("%d\n",<-c)
		}
	}()
	go func(c chan int) {
		c <-9
	}(c)
	go func(c chan int) {
		c <-10
	}(c)
	go func() {
		for{
			fmt.Println("a")
			fmt.Printf("%d\n",<-c)
		}
	}()

	time.Sleep(time.Microsecond)
}

func closedChanDemo(){
	c := make(chan int)

	go func() {
		for{
			//close 下
			// 第一种遍历方式
			//n, ok := <-c
			//if !ok{
			//	break
			//}
			//fmt.Printf("%d %v\n",n,ok)
			// 第二种遍历方式
			for n:=range c{
				fmt.Printf("%d\n",n)
			}


		}
	}()
	c <- 2

	c <- 3

	c <- 4

	c <- 5

	c <- 6

	c <- 7
	c <- 8

	close(c)
	time.Sleep(time.Microsecond)
	//fmt.Println(c)
}

func main()  {
	fmt.Println("Channel as first-class citizen")
	chanDemo()
	//fmt.Println("Buffered channel")
	//bufferedChanDemo()
	//fmt.Println("Channel close and range")
	//closedChanDemo()

}
