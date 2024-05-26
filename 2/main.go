package main

import "fmt"

type countCatchingAnimals int

type Cage struct {
	size      string
	hasAnimal bool
}

type Animal struct {
	cage    Cage
	name    string
	runAway bool
}

type Zookeeper struct {
	name        string
	catchAnimal countCatchingAnimals
}

func (z *Zookeeper) catch(a Animal) {
	a.cage.hasAnimal = true
	a.runAway = false
	z.catchAnimal++
	fmt.Printf("%s in cage \n", a.name)
	fmt.Printf("%s catch %d animals \n", z.name, z.catchAnimal)
}

func main() {

	tiger := Animal{
		cage: Cage{
			size:      "Big",
			hasAnimal: false,
		},
		name:    "Tiger",
		runAway: true,
	}

	zebra := Animal{
		cage: Cage{
			size:      "Big",
			hasAnimal: false,
		},
		name:    "Zebra",
		runAway: true,
	}

	elephant := Animal{
		cage: Cage{
			size:      "Big",
			hasAnimal: false,
		},
		name:    "Elephant",
		runAway: true,
	}

	monkey := Animal{
		cage: Cage{
			size:      "Big",
			hasAnimal: false,
		},
		name:    "Monkey",
		runAway: true,
	}

	giraffe := Animal{
		cage: Cage{
			size:      "Big",
			hasAnimal: false,
		},
		name:    "Giraffe",
		runAway: true,
	}

	z := Zookeeper{
		name: "John",
	}
	z.catch(tiger)
	z.catch(zebra)
	z.catch(elephant)
	z.catch(monkey)
	z.catch(giraffe)

}
