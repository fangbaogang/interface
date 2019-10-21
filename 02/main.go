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
	if ok {

		fmt.Println(e.name)
		e.say()
	}
}
