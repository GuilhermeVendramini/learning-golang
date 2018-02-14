package main

import "fmt"

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

func specs(a interface{}) {
	fmt.Println(a)
}

func main() {
	fido := dog{animal{"Toto", "woof"}, true}
	fifi := cat{animal{"Nina", "meow"}, true}
	specs(fido)
	specs(fifi)
}
