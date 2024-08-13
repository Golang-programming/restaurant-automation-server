package main

import "fmt"

// Define an interface for Animal
type Animal interface {
	Speak() string
}

// Define a struct for Dog that implements the Animal interface
type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

// Define a struct for Cat that implements the Animal interface
type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}

func main() {
	// Create instances of Dog and Cat
	var dog Animal = Dog{}
	var cat Animal = Cat{}

	// Call the Speak method on both
	fmt.Println(dog.Speak()) // Output: Woof!
	fmt.Println(cat.Speak()) // Output: Meow!
}
