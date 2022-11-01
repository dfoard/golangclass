/* Alta3 Research | RZFeeser
   Methods - value and pointer receivers   */


package main
 
import (
    "fmt"
)
type Virtmach struct{
    ip string
    hostname string
    diskgb int
    ram int
}

// display hostname
func (v Virtmach)getHostName() (string) {
    return v.hostname
}

//change hostname
func (v *Virtmach)setHostName(name string){
    v.hostname = name
}

func (v Virtmach)display(){
    fmt.Println("Primary IP Address:", v.ip)    // primary IP address 
    fmt.Println("Hostname:", v.hostname)        // hostname 
    fmt.Println("Disk GB:", v.diskgb)           // diskgb 
    fmt.Println("RAM:", v.ram)                  // ram 
}


 
func main() {
    machine1 := Virtmach{
       ip: "1.1.1.1",
       hostname: "MagicServer",
       diskgb: 4048,
       ram: 64,
     }
    
     // display hostname
     fmt.Println("Cuurent hostname: ",machine1.getHostName()) 

     //change hostname
     machine1.setHostName("AwesomeServer")

     //display all values
     machine1.display()
}
