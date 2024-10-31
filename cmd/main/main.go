package main

import (
	"fmt"
    "github.com/musooo/frequency-analysis/internal"
    "github.com/musooo/frequency-analysis/pkg"
)

func main() {
    str := pkg.ReadFileAndReturnString("file.txt")
    arr,count := internal.CharFrequency(&str,internal.FrequencyArrInit())
    internal.RelativeFrequency(&arr.Fs,count,&arr.RelativeF)
    internal.IsCalc(&arr.RelativeF,&arr.Is)
    entropyTot := internal.EntropyCalc(&arr.RelativeF,&arr.Is, &arr.Hs)
    fmt.Println("S F(s) f(s) H(s)")
    for key,value := range arr.Fs {
		fmt.Println(key, " ",value, arr.RelativeF[key], " ", arr.Is[key] , " ", arr.Hs[key])
	}
    fmt.Println("tot ", count, " 1 ", "/ ", entropyTot)

}