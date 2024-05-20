package main

import "fmt"

var firstName string = "Vira"
var age int = 35
var city string = "Kyiv"
var hobby string = "drawing"

func aboutMe(name string, age int, city string, hobby string) {

	fmt.Printf("My name is %s. I'm %d years old. I'm from %s. I like %s.", name, age, city, hobby)
}

func main() {
	aboutMe(firstName, age, city, hobby)
}
