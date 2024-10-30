package main

import (
	"fmt"
    "github.com/musooo/frequency-analysis/internal"
    "github.com/musooo/frequency-analysis/pkg"
)

func main() {
    str := pkg.ReadFileAndReturnString("file.txt")
    fmt.Println(str)
    arr := internal.CharFrequency(&str,internal.FrequencyArrInit())
    fmt.Print(arr.Arr)
}