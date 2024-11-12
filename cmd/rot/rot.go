package rot

import (
	"fmt"
	"github.com/musooo/frequency-analysis/pkg"
)

func rotN(text *string, num int) string {
    var outS string
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
    return outS
}

func Rot(){
	str := pkg.ReadFileAndReturnString("file.txt")
	fmt.Println(rotN(&str,12))
	// for i:=1; i<27; i++{
	// 	fmt.Println(rotN(&str,i))
	// 	fmt.Println("-----------")
	// }
}