package main

import (
	"fmt"
)

type human interface {
	walk()
	talk()
}

type person struct {
	firstname string
	lastname  string
}

type superhuman interface {
	human
	fly()
}

type superman struct {
	wearUnderwear bool
}

func (s superman) walk() {
	fmt.Printf("Superhuman can walk!!!\n")
}

func (s superman) talk() {
	fmt.Printf("Super human can talk!!\n")
}

func (s superman) fly() {
	fmt.Println("Super human waer Under wear outside :", s.wearUnderwear)
}

func (p person) walk() {
	fmt.Printf("%s %s walk!!!\n", p.firstname, p.lastname)
}

func (p person) talk() {
	fmt.Printf("%s talk!!\n", p.firstname)
}

type dog struct {
	name string
}

func (d dog) walk() {
	fmt.Printf("%s walk!!\n", d.name)
}

func (d dog) talk() {
	fmt.Printf("%s tallk\n", d.name)
}

func walkAndtalk(h human) {
	h.talk()
	h.walk()
}

func doeverything(s superhuman) {
	s.fly()
	s.talk()
	s.walk()
}
func doit(i interface{}) {
	fmt.Println("Value received: ", i)
}

func main() {
	p := person{
		firstname: "Dinesh",
		lastname:  "Kumar",
	}

	d := dog{
		name: "Tommy",
	}
	s := superman{
		wearUnderwear: true,
	}
	p.walk()
	p.talk()
	// d.talk()
	// d.walk()
	walkAndtalk(d)
	doit(4)
	doit("Ramesh")

	doeverything(s)
}
