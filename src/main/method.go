package main

import (
	"fmt"
)

type A struct {
	name string
}

type B struct {
	name string
}

//注意指针传递哦！
func (a *A) test() {
	a.name = "nazibo"
	fmt.Println("method A.test!")
}

func (b B) test() {
	b.name = "suredandan"
	fmt.Println("method B.test!")
}

type mt int

func (i *mt) test() {
	fmt.Println("method mt!!!")
}

func mainmethod() {
	a := A{}
	a.test()
	fmt.Println(a)

	b := B{}
	b.test()
	fmt.Println(b)

	var ii mt
	ii.test()

	//下面这个和先定义一个mt，再去调用 .test()是等价的
	//*和&只是为了指针传递
	(*mt).test(&ii)

	//以下是实现了基础int类型的一个increase方法作为示例
	var j mt
	j.increse(999)
	fmt.Println(j)
}

func (i *mt) increse(num int) {
	*i += mt(num) //注意这个地方，虽然num是int型，但要先转一次转成mt再来做+
}
