package main

import (
	"fmt"
)

func main() {
	var myFloat32 float32 = 4.5
	var myComplex64 complex64 = 4.5
	fmt .Println(myFloat32)
	fmt.Println(myComplex64)
	
	type AltaString string

	//this block of code will not work because og Go's Highly Type framework
	/*
	const myString string = "Hello"
	var zstring AltaString = mySttring
	*/
	
	//This will work as myUntypeString is a untyped variable
	const myUntypedString = "Alta3 Research"
	var uts AltaString = myUntypedString
	
	fmt.Println(uts)
	
	// Typed constants will NOT work
    	// the const 'typedInt' can ONLY be used with type int
    	/*
    	const typedInt int             = 1
    	var myFloat64 float64          = typedInt      // compiler error
    	*/
}
