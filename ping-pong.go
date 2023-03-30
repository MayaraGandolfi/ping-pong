package main

import (
	"fmt"
	"time"
)

func ping(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func pong(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}
func imprimir(c chan string, c2 chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second)
		msg2 := <-c2
		fmt.Println(msg2)
		time.Sleep(time.Second)
	}
}

func main() {
	var c chan string = make(chan string)
	var c2 chan string = make(chan string)

	go ping(c)
	go pong(c2)
	go imprimir(c, c2)

	var entrada string
	fmt.Scanln(&entrada)
}
