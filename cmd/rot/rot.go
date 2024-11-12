package rot

import (
	"fmt"
	"github.com/musooo/frequency-analysis/pkg"
)

func rotN(text *string, num int, ) string { //wg *sync.WaitGroup
    //defer wg.Done()
    var outS string
    outS = fmt.Sprintf("rot:%d\n", num)
    for _, char := range *text {
        if char >= 'A' && char <= 'Z' {
            shifted := ((int(char)-'A')+num)%26 + 'A'
            outS += string(rune(shifted))
        } else if char >= 'a' && char <= 'z' {
            shifted := ((int(char)-'a')+num)%26 + 'a'
            outS += string(rune(shifted))
        } else {
            outS += string(char)
        }
    }
    outS += ("\n------------\n")

    return outS
}

func Rot(){
	str := pkg.ReadFileAndReturnString("file.txt")
	//fmt.Println(rotN(&str,12))
    // var wg sync.WaitGroup
	// for i:=1; i<27; i++{
    //     wg.Add(1)
    //     go rotN(&str,i, &wg)
	// }

    // wg.Wait()
    for i:=1; i<27; i++{
        fmt.Println(rotN(&str,i))
    }
}