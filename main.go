package main

import (
	"fmt"
	"github.com/cleysonlb/go-design-pattern/singleton"
)

func main() {
	fmt.Println("Go design patterns")

	//  Usando Singleton
	s := singleton.New()
	s.Name = "Luke Skywalker"
	fmt.Println(s.Name)
	s2 := singleton.New()
	fmt.Println(s2.Name)
}