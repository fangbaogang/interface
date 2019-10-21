package main

import "fmt"

type Sayer interface {
	say()
}

type dog struct {
	name string
}

func (d *dog) say() {
	fmt.Println("wangwang")
}

func create_dog(name string) Sayer {

	d := &dog{
		name: name,
	}
	return d
}

func main() {
	d := create_dog("黑豹")
	e, ok := d.(*dog)
	//ok 是 true 表明 Sayer 存储的是 *dog 类型的值
	//如果需要区分多种类型，可以使用 switch 断言，更简单直接，这种断言方式只能在 switch 语句中使用
	if ok {
		fmt.Println(e.name)
		e.say()
	}
}
