package main

import "fmt"

func main() {
	a1 := []int{1,2,3,4,5} // 这个就相当于初始化了一个slice
	fmt.Println(a1)

	var a2 = [3]int{2,3,4}
	fmt.Println(a2)

	for i := 0; i < len(a1); i++ {
		fmt.Println(a1[i])
	}

	changeArray1(a1) // a1: [1 2 3 4 5]
	fmt.Println(a1) // a1 : [10086 2 3 4 5]

	changeArray2(a2) // a2： [2 3 4]
	fmt.Println(a2) //a2 : [2 3 4]


}
// Go 语言中对未指定长度数组的处理，一般采用 切片 的方式，切片包含对底层数组内容的引用，作为函数参数时，类似于 指针传递，函数中的修改对调用者可见
func changeArray1(a []int) {
	a[0] = 10086
}

//Go 语言的数组是值，其长度是其类型的一部分，作为函数参数时，是 值传递，函数中的修改对调用者不可见
func changeArray2(a [3]int) {
	a[0] = 10085
}