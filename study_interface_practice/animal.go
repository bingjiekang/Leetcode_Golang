package main

import (
	"fmt"
)

type AnimalIN interface {
	// 吃
	Eat()
	// 喝
	Drink()
	// 睡
	Sleep()
}

type animal struct {
}

func (a animal) Eat() {
	fmt.Println("动物接口Eat的实现")
}

func (a animal) Drink() {
	fmt.Println("动物接口Drink的实现")
}

func (a animal) Sleep() {
	fmt.Println("动物接口Sleep的实现")
}

type cat struct {
	animal
}

// 猫方法的重载
func (c cat) Drink() {
	fmt.Println("猫咪喝存净水")
}

// 子方法的重写
func (c cat) Game() {
	fmt.Println("猫需要玩耍")
}

func main() {
	var tm animal
	var xiaobai cat
	var tmIN AnimalIN = tm
	var tminc AnimalIN = xiaobai
	tmIN.Eat()
	tmIN.Drink()
	tmIN.Sleep()
	tminc.Eat()
	tminc.Drink()
	tminc.Sleep()
	xiaobai.Game()

}
