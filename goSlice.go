package main

import "fmt"
/**
这里有个公式，对于底层数组容量是 k 的切片 slice[i:j] 来说：
长度: j-i
容量: k-i
 */
func main() {
	var array = []int{0,1,2,3,4,5,6,7,8}
	s1 := array[:]
	print(s1)

	s2 := array[2:4]
	print(s2)

	var s3 = append(s2,11,22,33)
	print(s3)

	/*for _,s := range s3{
		fmt.Println(s)
	}*/

	s4 := make([]int,10)
	copy(s4,s3)
	print(s4)
}

func print(s []int) {
	fmt.Println(s,len(s),cap(s))
}
