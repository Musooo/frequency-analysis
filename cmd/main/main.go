package main
 
import (
    "fmt"
    "os"
    "bufio"
    //"github.com/musooo/frequency-analysis/internal"
)
//space = 32
func main() {
    file, err := os.Open("file.txt")
    var f string 
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()
 
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        f += scanner.Text()
        
    }
    
    fmt.Println(f[4])

    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }
}