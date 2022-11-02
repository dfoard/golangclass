/* Alta3 Research | RZFeeser
   Refer - no panic, refer or defer   */

package main

import(
    "fmt"
    "os"
    "io"
)
    
func CopyFile(dstName, srcName string) (written int64, err error) {
    fmt.Println("source:",srcName)
    fmt.Println("dest:", dstName)
    src, err := os.Open(srcName)
    if err != nil {
        return
    }

    dst, err := os.Create(dstName)
    if err != nil {
        fmt.Println("error:", err)
        return
    }

    written, err = io.Copy(dst, src)
    dst.Close()
    src.Close()
    return
}

func main() {
             // new name of our file    // file we are copying
    CopyFile("classnotes-copy1.txt", "classnotes.txt")
}
