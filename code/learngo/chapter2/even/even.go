package even

//大写字母开头的函数可导出
func Even(i int) bool {
	return i%2 == 0
}

//
func odd(i int) bool {
	return i%2 == 1
}
