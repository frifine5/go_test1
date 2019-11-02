package main

import (
	"sync"
	"time"
	"fmt"
	"strconv"
)


var wg sync.WaitGroup


//func say(s string){
//	for i:=0; i<3;i++{
//		println(s)
//	}
//	wg.Done()
//}


func say(s string, c chan string){
	for i:=0; i<3;i++{
		c <- s
	}

	wg.Done()
}

func proc1(){
	time.Sleep(time.Second * 1)
	wg.Add(2)

	ch := make(chan string)

	go say("Hello", ch)
	go say("World", ch)

	i := 1
	for{
		str := <- ch
		println(str)
		if i>=6{
			close(ch)
			break
		}
		i++
	}

	//time.Sleep(time.Second * 1)
	wg.Wait()

}

/*

 */

func proc2(){
	ch1 := make(chan int)
	ch2 := make(chan string)
	go pump1(ch1)
	go pump2(ch2)
	go suck(ch1, ch2)
	time.Sleep(time.Duration(time.Second * 30))

}


func pump1(ch chan int){
	for i:=0;;i++{
		ch <- i *2
		time.Sleep(time.Duration(time.Second))
	}
}

func pump2(ch chan string){
	for i:= 0; ; i++{
		ch <- strconv.Itoa(i + 5)
		time.Sleep(time.Duration(time.Second))
	}
}

func suck(ch1 chan int, ch2 chan string ){
	chRate := time.Tick(time.Duration(time.Second * 5))
	for{
		select{
		case v := <- ch1:
			fmt.Printf("Recv on channel 1: %d\n", v)
		case v := <- ch2:
			fmt.Printf("Recv on channel 2: %s\n", v)
		case <- chRate:
			fmt.Printf("Log log...\n")
		}
	}

}







func main(){
	//proc1()

	proc2()
}


