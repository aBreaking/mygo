package main

import (
	"fmt"
	"time"
)

func main() {
	p1:=Person{
		name: "zhangsan",
		age: 18,
		birthday: time.Now(),
	}

	fmt.Println(p1)

	p2 := Person{"lisi",19,time.Now()}
	fmt.Println(p2)

	var p3 Person
	p3.name = "p3"
	fmt.Println(p3)

	changeStruct(p3)
	fmt.Println(p3)

	changePoint(&p3)
	fmt.Println(p3)

}

/**
go 都是值传递，虽然我传入了对象，修改了对象里的属性，但是在方法外面该对象的属性还是不变的
 */
func changeStruct(p Person) {
	p.name = "changeStruct"
	fmt.Println("cs",p)
}

/**
引用传递 必须是指针类型，这样修改了对象的属性，在方法外面也是生效的
 */
func changePoint(p *Person) {
	p.name = "pointer"
}


type Person struct{
	name string
	age int
	birthday time.Time
}