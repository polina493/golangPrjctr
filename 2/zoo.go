package main

import (
	"fmt"
)

type animal struct {
	name string
	age  int
}

type lion struct {
	animal
	maneLength int
}

type cage struct {
	kind          string
	generalAmount int
	amountInCage  int
}

func countAmount(a []lion) int {
	amountOfAnimals := 0
	for range a {
		amountOfAnimals++
	}
	return amountOfAnimals
}

func (c *cage) zookeeper() {
	for c.amountInCage < c.generalAmount {
		c.amountInCage++
	}
}

func printStruct(c cage) {
	fmt.Printf("Kind: %v, General amount of animals in the zoo: %v, Animals amount in the cage: %v.\n", c.kind, c.generalAmount, c.amountInCage)
}

func main() {
	var lions = []lion{
		{
			animal:     animal{name: "Lion1", age: 5},
			maneLength: 20,
		},
		{
			animal:     animal{name: "Lion2", age: 5},
			maneLength: 20,
		},
		{
			animal:     animal{name: "Lion3", age: 5},
			maneLength: 20,
		},
		{
			animal:     animal{name: "Lion4", age: 5},
			maneLength: 20,
		},
		{
			animal:     animal{name: "Lion5", age: 5},
			maneLength: 20,
		},
		{
			animal:     animal{name: "Lion6", age: 5},
			maneLength: 20,
		},
	}

	lionsAmount := countAmount(lions)

	cageLions := cage{
		kind:          "lions",
		generalAmount: lionsAmount,
		amountInCage:  1,
	}

	printStruct(cageLions)
	cageLions.zookeeper()
	printStruct(cageLions)
}
