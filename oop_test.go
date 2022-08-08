package oop_test

import (
	"fmt"

	"github.com/rmasci/oop"
)

func ExampleHuman_InitHuman() {
	oop.Human{
		First: "Clark",
		Last:  "Kent",
		Age:   32,
	}
}

func ExampleRegularHuman_JsonPrint() {
	var person1 oop.RegularHuman
	var person2 oop.SuperHuman
	person1.InitRegularHuman("Peter", "Parker", "Photographer", 23)
	person2.InitSuperHuman("Spider", "Man", "Spider dude", 23)
	out, _ := person1.JsonPrint()
	fmt.Println(out)
	out, _ = person2.JsonPrint()
	fmt.Println(out)
	/*
		{
		        "First": "Peter",
		        "Last": "Parker",
		        "Age": 23,
		        "Job": "Photographer",
		        "OldPerson": "Your human is young"
		}
		{
		        "First": "Spider",
		        "Last": "Man",
		        "Age": 23,
		        "Superpower": "Spider dude",
		        "OldSuperhero": "Super hero is young"
		}

	*/
}

func ExampleSuperHuman_InitSuperHuman() {
	var person1 oop.RegularHuman
	var person2 oop.SuperHuman
	person1.InitRegularHuman("Peter", "Parker", "Photographer", 23)
	person2.InitSuperHuman("Spider", "Man", "Spider dude", 23)
	out, _ := person1.JsonPrint()
	fmt.Println(out)
	out, _ = person2.JsonPrint()
	fmt.Println(out)
	/*
		{
		        "First": "Peter",
		        "Last": "Parker",
		        "Age": 23,
		        "Job": "Photographer",
		        "OldPerson": "Your human is young"
		}
		{
		        "First": "Spider",
		        "Last": "Man",
		        "Age": 23,
		        "Superpower": "Spider dude",
		        "OldSuperhero": "Super hero is young"
		}

	*/
}
