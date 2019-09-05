package main

import (
	"bufio"
	"bytes"
	"container/list"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"sync"
)

func main(){
	var p *int //定义一个指针
	fmt.Printf("%v\n",p)

	var i int //定义一个整型变量
	p = &i //使p指向i

	//打印p（内存地址）
	fmt.Printf("%v\n",p)
	//从指针获取值，通过在指针变量前置‘*’实现

	//修改i的值
	*p = 8
	fmt.Printf("%v\n",*p)
	fmt.Printf("%v\n",i)

	//Go 的内存分配原语：make new
	//用new分配内存 new(T) 返回一个指针（指向新分配的类型T的零值）
	t := new(SyncedBuffer)
	var v SyncedBuffer
	fmt.Printf("%v-%v\n",t,v)
	//用make分配内存，内建函数make(T,args)与new(T)有着不同的功能，它只能创建slice、map和channel，并且返回一个有初始值（非零）的T类型，而不是*T
	// tips:导致这三个类型有所不同的原因是指向数据结构的引用必须被初始化
	//例如 make([]int,10,100)分配了100个整数的数组，然后用长度10和容量100创建了slice结构指向数组的前10个元素，区别是，new([]int)返回指向新分配的内存的指针，而零值填充的slice结构是指向nil的slice值

	//eg.例子
	var newp *[]int = new([]int) //分配slice结构内存，很少使用
	fmt.Printf("%v\n",newp)
	var makev []int = make([]int,100) // 指向一个新分配的有100个整数的数组
	fmt.Printf("%v\n",makev)

	/**
		不必要的例子
	var p *[]int = new([]int)
	*p = make([]int,100,100)
		常用
	v := make([]int,100)
	 */
	//切记：make仅用于map，slice，和channel，并且返回的不是指针，应当用new获得特定的指针

	//构造函数与复合声明
	//有时零值不能满足需求，必须要有一个用于初始化的构造函数

	//自定义类型

	nameage := new(NameAge)
	nameage.name = "小胡"
	nameage.age = 24
	fmt.Printf("%v\n%s（name）-%d（age）",nameage,nameage.name,nameage.age)

	//结构字段
	fmt.Printf("%v\n", )

	//转换，有时需要将一个类型转换为另一种类型
	/**
		From b   []byte |	i []int 	r []rune	| s string 	f float32 i int
		To
		[]byte X									[]byte(s)
		[]int			X							[]int(s)
		[]rune							X			[]rune(s)
		string	string(b) string(i)		string(r)	x
		float32 												x		float32(i)
		int 													int(f)		x
	 */
	//用户定义类型转换

	//组合 Go不是面向对象的语言，因此并没有继承

	//================================================练习题
	//Q18
	m := []e{1,2,3,4}
	s := []e{"a","b","v","d"}
	mf := Map(m,mult2)
	sf := Map(s,mult2)
	fmt.Printf("%v\n",mf)
	fmt.Printf("%v\n",sf)

	//Q19
	var p1 foo2 //p1的类型是foo2
	var p2 = new(foo2) //p2的类型是*foo
	fmt.Printf("%v\n",p1)
	fmt.Printf("%v\n",p2)

	//Q20 Linked list
	l := list.New()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(4)
	l.PushBack(6)

 	for e:=l.Front();e != nil; e = e.Next() {
 		fmt.Printf("%v\n",e.Value)
	}

 	//Q20.2 Doubly linked list
 	l1 := new(List)
	l1.Push(1)
	l1.Push(2)
	l1.Push(4)
 	for n:= l1.Front(); n != nil;n = n.Next() {
 		fmt.Printf("%v\n",n.Value)
	}
 	fmt.Println()
	for v,err := l1.Pop();err == nil ;v,err = l1.Pop() {
		fmt.Printf("%v\n",v)
	}
	//Q21 Cat
	flag.Parse()
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	}
	for i := 0; i < flag.NArg(); i++ {
		f,e := os.Open(flag.Arg(i))
		if e != nil {
			_, _ = fmt.Fprintf(os.Stderr, "%s:error reading from %s:%s\n", os.Args[0], flag.Arg(i), e.Error())
			continue
		}
		cat(bufio.NewReader(f))
	}
}
/**********
	A21 Begin
***************/
var numberFlag  = flag.Bool("n",false,"number each line")
func cat(r *bufio.Reader) {
	i := 1
	for {
		buf,e := r.ReadBytes('\n')
		if e == io.EOF {
			break
		}

		if *numberFlag {
			_, _ = fmt.Fprintf(os.Stdout, "%5d  %s", i, buf)
			i++
		} else {
			_, _ = fmt.Fprintf(os.Stdout, "%s", buf)
		}
	}
	return
}
/**********
	A21 End
***************/

/********
A20.2 begin
 */
type Value int //声明一个list要用的类型
type Node struct { //声明一个节点类型
	Value
	prev,next *Node
}
type List struct { //模拟container/list接口
	head,tail *Node
}

func (l *List) Front() *Node {
	return l.head
}

func (n *Node) Next() *Node {
	return n.next
}
func (l *List)Push(v Value) *List  { //When pushing ,create a new Node witeh the provided value
	n := &Node{Value :v}
	if l.head == nil { //if the list is empty put the new node at the head
		l.head = n
	} else { // otherwise put it at the tail
		l.tail.next = n
		n.prev = l.tail // make sure the new node points back to the previously existing one
	}
	l.tail = n // point tail to the newly inserted node
	return l
}
var errEmpty = errors.New("List is Empty")

func (l *List) Pop() (v Value,err error)  {
	if l.tail == nil {
		err = errEmpty // when popping,return an error if the list is empty
	} else {
		v = l.tail.Value
		l.tail = l.tail.prev
		if l.tail == nil {
			l.head = nil
		}
	}

	return v,err
}

/***
A20.2 end
 */
type SyncedBuffer struct{
	lock sync.Mutex
	buffer bytes.Buffer
}

func NewFile(fd int, name string) *os.File{
	if fd < 0 {
		return  nil
	}
	f := new(os.File)
	f.Fd()
	f.Name()

	//复合声明
	// f := File(fd,name,nil,0)
	// return &f 返回f的地址

	// 也可以 return &File{fd,name,nil,0} 所有字段都必须按顺序写上，当然也可以 指定字段return &File{fd:fd,name:name}
	//从复合声明中获取地址，意味着告诉编译器在堆中分配空间，而不是栈中
	//在特定的情况下，如果复合声明不包含任何字段，它创建特定类型的零值，表达式 new(File) 与&File{}是等价的

	return f
}

//自定义类型
type foo int

type foo2 struct {
}

type NameAge struct {
	name string
	age int
}

type e interface{}

func mult2(f e) e {
	switch f.(type) {
	case int:
		return f.(int) * 2
	case string:
		return f.(string) + f.(string) + f.(string) + f.(string)
	}
	return f
}

func Map(n []e,f func(e) e) []e {
	m := make([]e,len(n))
	for k,v := range n{
		m[k] = f(v)
	}
	return m
}

