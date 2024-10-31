package main

import (
	"fmt"
    "github.com/musooo/frequency-analysis/internal"
    "github.com/musooo/frequency-analysis/pkg"
)

func main() {
    str := pkg.ReadFileAndReturnString("file.txt")
    fmt.Println(str)
    arr,count := internal.CharFrequency(&str,internal.FrequencyArrInit())
    for key, value := range arr.Fs {
        fmt.Printf("Chiave: %s, Valore: %d\n", key, value)
    }

    internal.RelativeFrequency(&arr.Fs,count,&arr.RelativeF)
    for key, value := range arr.RelativeF {
        fmt.Printf("Chiave: %s, Valore: %f\n", key, value)
    }

    internal.IsCalc(&arr.RelativeF,&arr.Is)
    for key, value := range arr.Is {
        fmt.Printf("Chiave: %s, Valore: %d\n", key, value)
    }

    internal.EntropyCalc(&arr.RelativeF,&arr.Is, &arr.Hs)
    for key, value := range arr.Hs {
        fmt.Printf("Chiave: %s, Valore: %f\n", key, value)
    }
}