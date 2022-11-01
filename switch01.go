/* RZFeeser | Alta3 Research
   switch - case and default  */

package main

import (
    "fmt"
    "runtime"
    "strings"
)

func main() {

           // init gover                
    var gove string = runtime.Version()
    var gover string =  "go18.3"
    switch {
    case strings.HasPrefix(gover,"go1.18"):                 // if matches "go1.18", do the code below then STOP
        fmt.Println("You are using the latest version of GoLang:",gover)
    case strings.HasPrefix(gover,"go1.17"):
        fmt.Println("This version of Go is fine:",gover)
    case strings.HasPrefix(gover,"go1.16"),strings.HasPrefix(gover,"go1.15"):       // if matches "go1.16", "go1.16.5", OR "go1.15" 
        fmt.Println("You are using an older, but acceptable version of GoLang:",gover)
    default:                       // in all other cases run the code below
        fmt.Println("I cannot make a recommendation.")
        fmt.Println("You're using version:",gover)
    }
   fmt.Println("Current version:",gove)
}
