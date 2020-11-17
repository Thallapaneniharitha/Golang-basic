1.//Create file
package main
import (  
    "fmt"
    "os"
)
func main() {  
    f, err := os.Create("myfile.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    l, err := f.WriteString("Hello World")
    if err != nil {
        fmt.Println(err)
        f.Close()
        return
    }
    fmt.Println(l, "bytes written successfully")
    err = f.Close()
    if err != nil {
        fmt.Println(err)
        return
    }
}
output:
11 bytes written successfully