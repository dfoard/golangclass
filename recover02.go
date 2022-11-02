
/* Alta3 Research | RZFeeser
   Refer - using the defer statement  */

package main

import(
    "os"
    "io"
    "fmt"
)
    
func CopyFile(dstName, srcName string) (written int64, err error) {
    fmt.Println("source:", srcName)
    fmt.Println("dest:", dstName)
    src, err := os.Open(srcName)
    if err != nil {
        fmt.Println("Error:",err)
        return
    }
    defer src.Close()

    dst, err := os.Create(dstName)
    if err != nil {
        fmt.Println("Error:",err)
        return
    }
    defer dst.Close()

    return io.Copy(dst, src)
}

func main() {
    CopyFile("classnotes-copy2.txt", "classnotes-copy1.txt")
}
