package main

import (
	"fmt"
)

func main() {

	var a bool = true

	if a  {
		fmt.Println("yes")
	}else {
		fmt.Println("no")
	}

	var b = "a"

	switch b {
	case "a":
		fmt.Println("is a")

		//fallthrough
	case "b":
		fmt.Println("is b")
	}

	//TODO
	var c interface{}

	switch c.(type) {
	case int:

	}

}