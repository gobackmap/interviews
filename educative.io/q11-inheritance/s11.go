package main

import (
	"fmt"
	"time"
)

type Animal struct {
	Weight    float32
	Height    float32
	Food      string
	SleepTime float64
}

func (a Animal) Eat() {
	fmt.Println("I'm eating", a.Food)
}

func (a Animal) Sleep() {
	fmt.Println("I'm sleep", a.SleepTime, "Hours")
}

type Dog struct {
	Animal
}

func main() {
	dog := new(Dog)
	dog.Food = "meat"
	dog.Weight = 32.65
	dog.Height = 65
	dog.SleepTime = (11 * time.Hour).Hours()
	dog.Sleep()
	dog.Eat()
}
