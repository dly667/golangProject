package main

import (
	"fmt"
	"time"
)

func createWorker() chan<- int{
	c := make(chan int)
	go func() {
		for{
			fmt.Printf("%d\n",<-c)
		}
	}()
	return c
}

func chanDemo(){
	cc := createWorker()
	cc <- 1000


	time.Sleep(time.Microsecond)
	//fmt.Println(c)
}

func bufferedChanDemo(){
	c := make(chan int,5)


	c <- 2
	c <- 3
	c <- 4
	c <- 5
	c <- 6
	c <- 7

	go func() {
		for{
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
	//chanDemo()
	fmt.Println("Buffered channel")
	//bufferedChanDemo()
	fmt.Println("Channel close and range")
	closedChanDemo()
}
