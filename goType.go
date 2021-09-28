package main

import "fmt"

func main() {

	var a int = -1
	fmt.Println(&a)
	var b = 2
	a = a+b
	fmt.Println(&a)

	var ptr = &a

	fmt.Printf("%p -> %d \n",ptr,*ptr)

	var p = &ptr
	fmt.Println(**p)

	println("这是分割符")

	var s = "abc"
	for i := range s {
		println(s[i])
	}


}

