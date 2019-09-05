package chapter1

/**
编写一个程序，打印1到100的数字，当是三的倍数的时候打印Fizz ，当是5的倍数的时候打印 Buzz,当同时是3和5的倍数的时候打印FizzBuzz
*/
func main() {
	fizzBuzz()
}

func fizzBuzz() {
	for i := 1; i <= 100; i++ {
		if (i%3 == 0) && (i%5 == 0) {
			println("FizzBuzz")
		} else if i%3 == 0 {
			println("Fizz")
		} else if i%5 == 0 {
			println("Buzz")
		} else {
			println(i)
		}
	}
}
