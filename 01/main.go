package main

import "fmt"

type Sayer interface {
	say()
}

type dog struct {
	name string
}

func (d dog) say() {
	fmt.Println("wangwang")

}

func create_dog(name string) *dog {

	o := &dog{
		name: name,
	}
	return o
}

func main() {

	d := create_dog("heibao")
	d.say()

}
