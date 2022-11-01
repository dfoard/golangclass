package main

import (  
    "fmt"
)

func main() {  
    
    // create a map type for programming languages
    programmingLang := map[string]string{
        "Python": ".py",
        "Golang": ".go",
        "Java": ".java",
        "Ansible": ".yml",
        "C++": ".cpp",
    }
    fmt.Println(programmingLang)

    //remove C++ key
    delete(programmingLang,"C++")

    fmt.Println(programmingLang)

    //add Julia key
    programmingLang["Julia"] = "j.1"

    fmt.Println(programmingLang)
}
