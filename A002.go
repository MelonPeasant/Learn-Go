package main

import "fmt"

type Animal_A2 interface {
	Speak()
}
type Dog_A2 struct {
	Name string
}

func (d Dog_A2) Speak() {
	fmt.Printf("I'm Dog, My Name is %s\n", d.Name)
}

type Cat_A2 struct {
	Name string
}

func (c Cat_A2) Speak() {
	fmt.Printf("I'm Dog, My Name is %s\n", c.Name)
}

func A002() {
	d := Dog_A2{Name: "Wang Cai"}
	d.Speak()

	c := Cat_A2{Name: "Mi Mi"}
	c.Speak()

	var an_A2 Animal_A2
	an_A2 = d
	an_A2.Speak()

}
