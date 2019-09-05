package chapter1

import "fmt"

func main() {
	//1.打印1-10
	forprac()
	// 2 使用goto 打印1 - 10
	gotofor()
	// 遍历array并打印
	forarray()
}
func forprac() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("for循环：%d \n", i)
	}
}
func gotofor() {
	i := 0
loop:
	fmt.Printf("goto: %d\n", i)
	if i < 10 {
		i++
		goto loop
	}

}
func forarray() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, v := range arr {
		fmt.Printf("array:%d\n", v)
	}
	fmt.Printf("%v", arr) // %v 打印默认格式的值，the value in a default format
}
