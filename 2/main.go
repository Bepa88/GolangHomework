package main

import "fmt"

type countCatchingAnimals int

type Animal struct {
	name string
}
type Cage struct {
	animal    Animal
	size      string
	hasAnimal bool
}

type Zookeeper struct {
	name        string
	catchAnimal countCatchingAnimals
}

func (z *Zookeeper) catching(a Animal) {
	var cage = Cage{
		animal:    a,
		size:      "Big",
		hasAnimal: true,
	}
	z.catchAnimal++
	fmt.Printf("%s in cage \n", cage.animal.name)
	fmt.Printf("%s catch %d animals \n", z.name, z.catchAnimal)
}

func main() {

	tiger := Animal{
		name: "Tiger",
	}

	zebra := Animal{
		name: "Zebra",
	}

	elephant := Animal{
		name: "Elephant",
	}

	monkey := Animal{
		name: "Monkey",
	}

	giraffe := Animal{
		name: "Giraffe",
	}

	z := Zookeeper{
		name: "John",
	}
	z.catching(tiger)
	z.catching(zebra)
	z.catching(elephant)
	z.catching(monkey)
	z.catching(giraffe)

}
