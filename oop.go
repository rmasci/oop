package oopdemo

import (
	"encoding/json"
)

// These are fields used for all humans
type Human struct {
	First string
	Last  string
	Age   string
}

// These are fields for Super Humans, it inherits the type and methods of Human
type SuperHuman struct {
	Human
	Superpower string
}

// This is for non-superhumans
type RegularHuman struct {
	Human
	Job string
}

// This is an interface that implements the JsonPrint function for both Regular Humans and Super Humans
type JsonOut interface {
	JsonPrint() (string, error)
}

// This function is passed with the first, last, age of your human.
func (h *Human) InitHuman(first, last, age string) {
	h.First = first
	h.Last = last
	h.Age = age
}

/// This initializes a superhuman pass in the first last and age
func (s *SuperHuman) InitSuperHuman(first, last, age string) {
	s.InitHuman(first, last, age)
}

// Give a superpower
func (s *SuperHuman) AssignSuperpower(superpower string) {
	s.Superpower = superpower
}

// Init a regular human
func (r *RegularHuman) InitRegularHuman(first, last, age string) {
	r.Human = Human{
		First: first,
		Last:  last,
		Age:   age,
	}
}

// Print out Regular Humans in JSON
func (r *RegularHuman) JsonPrint() (string, error) {
	jsonByte, err := json.MarshalIndent(r, "", "\t")
	return string(jsonByte), err
}

// Print out Superhumans in JSON
func (s *SuperHuman) JsonPrint() (string, error) {
	jsonByte, err := json.MarshalIndent(s, "", "\t")
	return string(jsonByte), err
}
