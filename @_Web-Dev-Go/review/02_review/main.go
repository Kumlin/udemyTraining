package main

import "fmt"

type person struct {
	name    string
	message string
}

func (p person) sayName() {
	fmt.Println(p.name)
}

func (p person) pSpeak() {
	fmt.Println(p.name, " ", p.message)
}

type secretAgent struct {
	person
	secretMessage string
}

func (s secretAgent) sayName() {
	fmt.Println(s.name)
}

func (s secretAgent) sSpeak() {
	fmt.Println(s.name, " ", s.secretMessage)
}

type human interface {
	sayName()
}

func nameHuman(h human) {
	h.sayName()
}

func main() {
	norm := person{"norman", "I am normal norman"}
	bond := secretAgent{person{"Bond", "I am Bond"}, "This is the secret message"}

	norm.pSpeak()
	bond.sSpeak()
	bond.person.pSpeak()
	nameHuman(bond)
	nameHuman(norm)
}
