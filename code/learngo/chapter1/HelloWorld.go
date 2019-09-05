package chapter1 // 1.go文件都必须在一个package中

import (
	"fmt"
) // 实现格式化I/O

func main() {
	fmt.Println("Hello World,の")
	// 变量定义
	// = 定义
	var a int = 1
	var b bool = false
	println("变量定义", b, a)
	// 用：= 定义
	c, d := 1, 2 // 通过值进行推导出来变量的类型
	println(":=定义变量", c, d)
	// 特殊的变量名 _ 任何赋值给他的值都会被抛弃
	_, e := 12, false
	println("特殊变量_", e)
	// Go编译器 对声明却未使用的变量会报错
	// 布尔类型 bool
	// 数字类型 int 在32位机器上 是 32 位，在64位系统上是 64位
	// 完整数字类型：int8, int16 int32 int64 byte uint8 uint16 uint32 uint64 byte是uint8的别名
	// 浮点型：float32 和 float64 64位的整数和浮点数总是64位的，即使是在32位的机器上
	// 这些类型都是独立的，混合着用会引起编译器错误
	// 常量 constant 编译时被创建，只能是数字、字符串或布尔值 (第一个iota表示0,再次新一行使用的时候值增加1)
	const (
		consta = iota
		constb = iota
	)
	println("常量", consta, constb)
	// 也可以省略重复的iota
	const (
		consta1 = iota
		constb1
		constc1
	)
	println("常量2", consta1, constb1, constc1)

	// 字符串.字符串在Go语言中是UTF-8字符由双引号("")包裹的字符序列，单引号（‘’）包裹的是字符，不是字符串
	s := "Hello World"
	println("字符串", s)

	s1 := "hello"

	carr := []rune(s1)
	// 替换数组第一个元素
	carr[0] = 'c'

	s2 := string(carr)
	fmt.Printf("字符串替换%s\n测试 %d \n", s2, 2)

	// 多行字符串
	mulStr := "test String1 String2"
	println("多行字符串", mulStr)

	// 复数 Go语言原生支持复数 变量类型 complex128(64位虚数部分)，还要小一些的 complex64(32位虚数部分)
	var complexnum complex64 = 5 + 5i
	fmt.Printf("复数Value is : %v\n", complexnum)
	// Go 语言错误 的内建类型 error
	//var e error = nil;
	//testPtr(a)
	getSumAndSub(1, 3)

	// 运算符和内建函数
	/**
	* / % << >> & &^
	+ - | ^
	== != < <= > >=
	<-
	&&
	||
	*/
	// GO不支持运算符重载（或者方法重载）而一些内建运算符却支持重载，比例 + 可以用于整数，浮点数、复数 和 字符串（字符串相加表示串联他们）

	fmt.Printf("$- %s\n:GO关键字", "关键字")
	/**
	break default func interface select
	case defer go map struct
	chan else goto package switch
	const fallthrough if range type
	continue for import return var
	go 用于并行
	select 用于选择不同类型的通讯
	struct 用于抽象数据类型
	*/
	fmt.Printf("$- %s\n:", "控制结构")
	/**
	在Go语言中只有很少的几个控制结构，没有 do 或者while循环，只有for 有灵活的switch 和if
	switch 接受像for那样可选的初始化语句，还有叫做类型选择和多路通讯转接器的select
	*/
	fmt.Printf("$- %s\n", "for")
	//Go语言有三种形式
	/**
	//和C的for一样
	for init; condition ; post {}
	和whild一样
	for condition {}
	死循环
	for {}
	*/

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		println(sum, i)
	}

	// 由于Go中没有逗号表达式，++ 和 -- 是语句而不是表达式，如果想在for循环中执行多个变量，应当使用平行赋值
	// Reverse s
	var arr = []rune(s)
	print("Before Reverse")
	println(string(arr))
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	println("After Reverse", string(arr))

	//break 和 continue
	for i := 0; i < 10; i++ {
		if i < 5 {
			break
		}
		println("break", i)
	}
	// 循环嵌套时，可以在break后面指定标签，用标签决定哪个循环被终止
J:
	for j := 0; j < 10; j++ {
	I:
		for i := 0; i < 5; i++ {
			println("嵌套", j, i)
			if i == 3 {
				break J
			}
			if i == 4 {
				break I
			}
		}
	}
	// continue 进入下一轮迭代
	for i := 0; i < 10; i++ {
		if i > 5 {
			continue
		}
		println("continue", i)
	}
	// range
	/**
	range 是一个迭代器，可用于循环，可以在slice array string map 和 channel 中
	当被调用时，从它循环的内容重返回一个键值对，基于不同的内容，range返回不同的值

	当对slice或者 array 迭代时，range返回序列号作为键，这个序号对于应的内容作为值
	*/
	list := []string{"a", "b", "c", "d", "e", "f"}
	for k, v := range list {
		println("range", k, v)
	}
	for pos, char := range "aのΦ中b" {
		fmt.Printf("character '%c' starts at byte position %d\n", char, pos)
	}

	// switch
	/**
	go的switch非常灵活，表达式不必是常量或者整数，执行过程从上至下，直到找到匹配项
	如果switch没有表达式，会匹配true。
	*/
	println(unhex('B'))
	// switch 不会在匹配失败后自动向下匹配，但是可以通过fallthrough使其这样做
	var i int = 0
	//不会调用到unhex
	switch i {
	case 0:
	case 1:
		unhex('C')
	}
	//会调用到unhex
	switch i {
	case 0:
		fallthrough
	case 1:
		unhex('C')
	}

	//用default可以指定当其他所有分支都不匹配的时候的行为
	switch i {
	case 0:
	case 1:
		unhex('C')
	default:
		ifcase()
	}
	// 分支可以使用逗号分隔的列表
	println(shouldEscape('c'))

	// Go 的内建函数，意味着不用引入任何包就可以使用
	/**
	close : 用于channel 通讯，使用它来关闭channel
	delete：用于在map中删除实例
	len 和 cap ：可用于不同类型，len用于返回字符串/slice和数组的长度
	new : 用于各种类型的内存分配
	make: 用于内建模型（map、slice和channel）的内存分配
	copy: 用于复制slice
	append： 用于追加splice
	panic 和 recover 用于异常处理机制
	print 和 println ： 底层打印函数，可以在不引入fmt包的情况下使用，主要用于调试
	complex/real和imag: 全部用于处理复数
	*/

	// array slice map
	var array1 [10]int
	// TODO 大小是类型的一部分，不同的大小是不同的类型，因此不能改变大小。数组他同样是值类型的，将一个数组赋值给另外一个数组，会复制所有的元素，尤其是向一个函数传递一个数组的时候，它会得到数组的一个副本，而不是一个指针

	array1[0] = 1
	array1[1] = 2
	fmt.Printf("The first element is %d\n", array1[0])
	//可以利用array在列表中进行多个值的排序，或者采用更加灵活的：slice
	//字典或者哈希类型在 Go中称为：map

	//1 声明一个数组
	var arra [3]int
	println(arra[0], arra[1], arra[2]) //会用0初始化
	//2 复合声明 Tips:复合声明允许你直接将值复制给array,slice,map
	arrb := [3]int{1, 2, 3}            // 或者 := [...]int{1,2,3,4,5}Go会自动统计元素个数
	println(arrb[0], arrb[1], arrb[2]) //初始化每个元素
	// 3 多维数组
	arrc := [3][2]int{[2]int{1, 2}, [2]int{3, 4}, [2]int{5, 6}}
	println(arrc[0][1], arrc[1][0], arrc[2][1]) // 2 3 6
	// Tips array slice map 的复合声明变得更加简单，使用复合声明的array slice map 元素复合声明的类型与外部类型一致时，可省略
	// eg. arrc := [3][3]int{{1,2},{3,4},{5,6}}

	//slice slice与array接近，但是在新的元素加入的时候可以增加长度，slice总是指向底层的一个array。slice是一个指向array的指针，slice是一个引用类型，这意味着当赋值某个slice到另外一个变量，两个引用会指向同一个array
	// 例如：如果一个函数需要一个slice参数，在内对slice的元素的修改也会体现在函数调用者中，和底层传递array指针类似
	// Tips 引用类型使用make创建
	sl := make([]int, 10)
	sl[0] = 2
	println(sl[0])
	// 给定一个array 或者 slice，一个新的slice通过a[I:J]的方式创建。这会创建一个新的slice 指向变量a，从序号I开始，结束在J之前，长度为J-I
	arrd := [...]int{1, 2, 3, 4, 5}
	sl1 := arrd[:]
	println(len(sl1))
	sl2 := arrd[1:4]
	println(len(sl2))

	// map 可以认为是一个用字符串做索引的数组：map[<from type>]<to type>

	monthdays := map[string]int{
		"Jan": 31, "Feb": 28, "Mar": 31,
		"Apr": 30, "May": 31, "Jun": 30,
		"Jul": 31, "Aug": 31, "Sep": 30,
		"Oct": 31, "Nov": 30, "Dec": 31, //末尾的，是必须的
	}
	//当只需要声明一个map时，可以使用make monthdays := make(map[string]int)

	//在map中搜索时，使用方括号
	println(monthdays["Jun"])
	year := 0
	for _, days := range monthdays { // 键没有使用，所以用 _,days
		year += days
	}
	fmt.Printf("Number of days in a year : %d \n", year)
	//向map中添加元素
	monthdays["Undecim"] = 30 // 添加一个月
	monthdays["Feb"] = 29     // 闰年的时候重写这个元素

	// 检查元素是或否存在
	var value int
	var preset bool
	value, preset = monthdays["Feb"]
	println(value, preset)
	// 另一种写法：
	v, ok := monthdays["Feb"]
	println(v, ok)
	//从map中移除元素
	delete(monthdays, "Feb")
	println(monthdays["Feb"])

}

// Go语言指针
func testPtr(num *int) {
	*num = 20
}

// 返回多个值
func getSumAndSub(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2

	return sum, sub
}

// go语言的创新，切片+延时
func ifcase() {
	x := 10
	if x > 0 {

	} else {

	}
}

func learngoto() {
	i := 0
Here: //标签
	println(i)
	i++
	goto There //跳转

There:
	println("there")
	goto There
	goto Here
}

func unhex(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	return 0

}
func shouldEscape(c byte) bool {
	switch c {
	case ' ', '?', '=', '#', '+': // , as 'or'
		return true
	}
	return false
}

// 使用两个switch 对字节数组进行比较
func Compare(a, b []byte) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
	}
	switch {
	case len(a) < len(b):
		return -1
	case len(a) > len(b):
		return 1
	}

	return 0
}
