package main

import "fmt"

type Human struct {
		Name string
}

type Action struct {
	Human
	Action string
}

func main() {
	a := Action{
		Human:  Human{Name: "Artem"},
		Action: "coding",
	}
	fmt.Println(a.Name)
}
