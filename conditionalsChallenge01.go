package main

import (
	"fmt"
	"math/rand"
	"time"
)


func main() {
 
    rand.Seed(time.Now().Unix())
    names := []string{
    "David",
    "Donna",
    "Kevin",
    "Adam",
    "Andre",
    }


   if val := names[rand.Intn(len(names))]; val == "Andre" {
	fmt.Println("Andre is son number 3")
   } else if val == "Donna" {
	fmt.Println("Donna is my wife")
   } else {
        fmt.Println("Not in my family")
   }
}
