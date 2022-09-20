/*
oop is an attempt at showing Go's Object Oriented abilities. 
* Encapsulation -- Data in a class can be hidden from other classes. This is done in Go but. Inside a package, if you define a struct or method that doesn't start with a capital letter, its not 'exported' to the rest of the code. It's encapsulated. 
  type Human struct {
    First string // Public
    Last string // Public
    Age int // Public
    birthday string // Private
   }
   Benefit is that when you're looking at code and you're 500 lines down, when you see 'birthday' you know that it is a private or unexported variable that is not publically accessible. 
* Inheritance -- Go doesn't make the Inheritance mistake, but uses 'Composition'. Inheritance: 'is-a-relationship' where Composition: "has-a-relationship"
  type RegularHuman struct {
    Human  // This is public, and 'inherits' Firt, Last, Age from Human.
    Job string // public. Gives our regular human something to do.
   }
 * Abstraction -- Interfaces can be used to create common abstraction that can be used by multiple types.
 * Polymorphism -- Access objects of different types through the same interface. Both RegularHumans and SuperHumans can be processed by JsonPrint. 

The code is pretty simple, it creates a Regular Human called "Peter Parker" and "Clark Kent", then it creates two Superheros Spiderman and Superman. Both RegularHuman and SuperHuman inherit Human. Then prints out the results in JSON.
*/
