/* CHALLENGE 01 - Create a program that creates a struct called Astro. 
   It should track the name (string), age (int), mission (string), and if they are needed on the next mission, isneeded (bool). 
   Create 3 structs, and populate them with fictitious data. 
  When you are done, print your records to the screen.

*/

package main

import "fmt"

type Astro struct {
    name string
    age  int
    mission string
    isneeded bool
}

// this is our new struct
type nasaMission struct {
        people []Astro
        number int
        message string
}

func main() {
	ast1 := Astro{name: "Goofy", age:  10, mission: "Playing around", isneeded: false}
	ast2 := Astro{name: "Donald", age: 90, mission: "Movies", isneeded: true}
	ast3 := Astro{name: "Minnie", age: 90, mission: "Support Donald", isneeded: true}

	fmt.Println(ast1)
	fmt.Println(ast2)
	fmt.Println(ast3)
	// slice named people made up of Astro
    	p := []Astro{ast1, ast2, ast3}

    	// display the slice
    	fmt.Println(p)
	
    	// select data from the struct
    	fmt.Println(p[2].mission)	

	// initialize a nasaMission struct, "s"
    	s := nasaMission{p, 3, "success"}
	    
    	// display "s" without fields
    	fmt.Println(s)
    	
    	// display "s" with fields
	    fmt.Printf("%+v\n", s)
}
