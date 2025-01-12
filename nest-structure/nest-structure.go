package main

import "fmt"

type Person struct {
}

func (p *Person) PrintHello() {
	fmt.Println("我是一个普通人")
}

func (p *Person) Speech() {
	fmt.Println("我是一个读书的学生")
}

type Student struct {
	Person
}

func (s *Student) PrintHello() {
	fmt.Println("我首先是个学生")
}

func (s *Student) private() {
	fmt.Println("我是一个私有的函数")
}

type CollageStudent struct {
	Student
}

func (c *CollageStudent) PrintHello() {
	fmt.Println("我是个大学生")
}

func main() {
	collage := CollageStudent{}
	collage.PrintHello()
	collage.Speech()
	collage.private()
}
