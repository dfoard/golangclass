/* RZFeeser | Alta3 Research
   if and else statements     */

package main

import "fmt"
import "math/rand"
import "time"

func main() {

rand.Seed(time.Now().UnixNano())

var age int = 42
var randage int = rand.Intn(101)


fmt.Println("given that age is", age) // show age

if age%2 == 0 {                       // if age mod 2 is 0
fmt.Println("age is even")
} else {                              // in all other cases
fmt.Println("age is odd")
    }

fmt.Println("given that age is", randage) // show age

if randage%2 == 0 {                       // if age mod 2 is 0
fmt.Println("new age is even")
} else {                              // in all other cases
fmt.Println("new age is odd")
    }
}
