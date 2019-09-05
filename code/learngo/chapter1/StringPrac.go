package chapter1

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// 1.打印A
	for i := 1; i <= 100; i++ {
		printA(i)
		println()
	}
	// 2.统计字符串在中字符的数量，同时输出字符串的字节数
	str := "asSASA ddd dsjkdsjs dk"
	printNumAndByte(str)
	// 3.replace
	replaceStr(str)
	// 4.reverse str
	reverseStr("footbar")
}

func printA(num int) {
	for i := 0; i < num; i++ {
		print("A")
	}
}
func printNumAndByte(str string) {
	println("字符数", len([]byte(str)))
	println("字节数", utf8.RuneCount([]byte(str)))
}

func replaceStr(str string) {
	r := []rune(str)
	copy(r[4:4+3], []rune("abc"))
	fmt.Printf("Before: %s\n", str)
	fmt.Printf("After: %s\n", string(r))

}

func reverseStr(str string) {
	r := []rune(str)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	fmt.Printf("%s\n", string(r))
}
