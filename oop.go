// Package oopdemo demonstrates some (but not all of) Go's ability to do OOP.

package oop

import (
	"encoding/json"
)

// In this you'll notice that you have a type (class) that is Human. Human is then inherited by both SuperHuman
// and by RegularHuman.  Human has a method InitHuman which is also inherited by RegularHuman and SuperHuman.
// This also shows an Interface that can be used as one function, to either RegularHuman or SuperHuman.

// Human These are fields used for all humans
type Human struct {
	First string
	Last  string
	Age   int
	alias string
}

// SuperHuman These are fields for Super Humans,
// it inherits the type and methods of Human
type SuperHuman struct {
	Human
	Superpower   string
	OldSuperhero string
	alias        string
}

// RegularHuman This is for non-superhuman
type RegularHuman struct {
	Human
	Job       string
	OldPerson string
}

// InitHuman This function is passed with the first, last, age of your human.
func (h *Human) InitHuman(first, last string, age int) {
	h = &Human{First: first, Last: last, Age: age}
}

// InitSuperHuman This initializes a superhuman pass in the first last and age
func (s *SuperHuman) InitSuperHuman(first, last, superpower string, age int) {
	// All variables are targeted for the GC to cleanup after this is run.
	// first, last, superpower, age -- will be gone once method loses scope.
	s.InitHuman(first, last, age)
	s.Superpower = superpower
}

// InitRegularHuman Init a regular human
func (r *RegularHuman) InitRegularHuman(first, last, job string, age int) {
	r.Human = Human{
		First: first,
		Last:  last,
		Age:   age,
	}
	r.Job = job
}

// JsonOut Polymorphism -- concept that you can access objects of different types through the same interface. This is
//an interface that implements the JsonPrint function for both Regular Humans and Super Humans
type JsonOut interface {
	JsonPrint() (string, error)
}

// JsonPrint Print out Regularhumans or Superhumans in JSON. This code is attached to an interface so that the same
// 'function' called JsonPrint can be used with a RegularHuman and SuperHuman type.  This saves us having to write some
// kind of if / then statement tha figures out which type the variable belongs to and printing it out accordingly.
func (r *RegularHuman) JsonPrint() (string, error) {
	if r.Age > 50 {
		r.OldPerson = "Has your human applied for AARP?"
	} else {
		r.OldPerson = "Your human is young"
	}
	jsonByte, err := json.MarshalIndent(r, "", "\t")
	return string(jsonByte), err
}

// JsonPrint Print out Regularhumans or Superhumans in JSON. This code is attached to an interface so that the same
// 'function' called JsonPrint can be used with a RegularHuman and SuperHuman type.  This saves us having to write some
// kind of if / then statement tha figures out which type the variable belongs to and printing it out accordingly.
func (s *SuperHuman) JsonPrint() (string, error) {
	if s.Age > 50 {
		s.OldSuperhero = "Super hero needs some Super Geritol!"
	} else {
		s.OldSuperhero = "Super hero is young"
	}
	jsonByte, err := json.MarshalIndent(s, "", "\t")
	return string(jsonByte), err
}
