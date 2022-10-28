package main

import "fmt"

// 用结构体定义动物的属性
type animals struct {
	name      string
	age       int
	height    float32
	weight    float32
	coat      string
	risklevel string
	food      foodeating
}

// 用结构体定义动物吃的食物，嵌套到动物属性中
type foodeating struct {
	meat      string
	vagetable string
	fruit     string
}

func main() {
	//用动物属性定义老虎，并为各属性赋值
	var a animals
	a.food.meat = "牛肉，羊肉，鸡肉..."
	a.food.vagetable = "无"
	a.food.fruit = "无"
	tiger := animals{
		name:      "大黄",         //老虎名字是大黄
		age:       7,            //年龄为7岁
		height:    2.1,          //高2.1m
		weight:    299.9,        //重299.9kg
		coat:      "黄色，白色，黑色混杂", //毛色
		risklevel: "极度危险",       //危险程度
	}
	//打印老虎的各个属性
	fmt.Printf("名字:%s\n年龄:%d岁\n身高:%fm\n体重:%fkg\n毛色:%s\n危险程度:%s\n食物:%s\n", tiger.name, tiger.age, tiger.height, tiger.weight, tiger.coat, tiger.risklevel, a.food)
}
