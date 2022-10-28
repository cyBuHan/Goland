package main

import "fmt"

// 定义一个doginter接口
type DogInter interface {
	Dog()
}

// 定义Dog属性结构体，包括名字，狗的食物，狗奔跑的速度以及狗的行为
type Dog struct {
	Name     string
	food     string
	speed    int
	behavior string
}

// 定义Dog函数
func (a *Dog) Dog() {
	fmt.Println(a.Name + "会" + a.behavior)
}

func main() {
	var b DogInter = &Dog{
		Name:     "小狗",
		behavior: "汪汪叫",
	}
	//直接使用接口调用函数和结构体
	b.Dog()
}
