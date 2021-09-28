package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go say(i)
	}

	println(1 * time.Second)
	time.Sleep(1* time.Second)
	say(10086)
}

func say(s int) {
	time.Sleep(time.Duration(s) * time.Second)
	fmt.Println(s)

}