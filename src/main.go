package main

import (
	"fmt"
	"strconv"
)

//Dog dsfkn nnff
type Dog struct {
	Animal
}

//Animal sd
type Animal struct {
	Age int
}

//Move sdf
func (a *Animal) Move() {
	fmt.Println("Animal moved")
}

// SayAge df
func (a *Animal) SayAge() {
	fmt.Printf("Animal age: %d\n", a.Age)
}

type Amount int

func (a *Amount) Add(add Amount) {
	*a += add
}

func main() {
	d := Dog{}
	d.Age = 3
	d.Move()
	d.SayAge()

	var a Amount
	a = 1
	a.Add(2)
	fmt.Println(a)
	a.Add(4)
	fmt.Println(a)

	i := 10
	t := "variable"

	fmt.Println(i)
	fmt.Println(t)

	strVar := "100"
	intVar, _ := strconv.Atoi(strVar)
	fmt.Println(intVar)
	strVar1 := "-52541"
	intVar1, _ := strconv.ParseInt(strVar1, 10, 32)
	fmt.Println(intVar1)
	strVar2 := "101010101010101010"
	intVar2, _ := strconv.ParseInt(strVar2, 10, 64)
	fmt.Println(intVar2)

}
