package main

import "fmt"

type animals interface {
	getName() string
}

type animal struct {
	name  string
	sound string
}

type dog struct {
	animal
	friendly bool
}

type cat struct {
	animal
	annoying bool
}

func specs(a animals) {
	fmt.Println(a)
	fmt.Println(a.getName())
}

func (a cat) getName() string {
	return fmt.Sprintln("Name:", a.name)
}

func (a dog) getName() string {
	return fmt.Sprintln("Name:", a.name)
}

func main() {
	fido := dog{animal{"Toto", "woof"}, true}
	fifi := cat{animal{"Nina", "meow"}, true}
	specs(fido)
	specs(fifi)
	fmt.Println("-------------")
	fmt.Println(fido.getName())
	fmt.Println(fifi.getName())
}
