package main

import (
	"fmt"
	"log"
)

type people struct {
	name string
	age  int
}

func (this *people) String() {
	log.Printf(fmt.Sprintf("<name is: %s, age is %d", this.name, this.age))
}

func (this *people) ToString() string {
	return fmt.Sprintf("<name is: %s, age is %d>", this.name, this.age)
}

type inter interface {
	ToString() string
}

func main() {
	var p = inter(&people{name: "test", age: 1})
	log.Println(p.(*people).name)
	log.Print(p.ToString())
}
