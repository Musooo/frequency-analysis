package main

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/musooo/frequency-analysis/internal"
	"github.com/musooo/frequency-analysis/pkg"
)

func main() {
    str := pkg.ReadFileAndReturnString("file.txt")
    arr,count := internal.CharFrequency(&str,internal.FrequencyArrInit())
    internal.RelativeFrequency(&arr.Fs,count,&arr.RelativeF)
    internal.IsCalc(&arr.RelativeF,&arr.Is)
    entropyTot := internal.EntropyCalc(&arr.RelativeF,&arr.Is, &arr.Hs)
    t := table.NewWriter()
    t.SetOutputMirror(os.Stdout)
    t.AppendHeader(table.Row{"S", "F(s)", "'f'(s)","I(s)", "H(s)"})
    for key,value := range arr.Fs {
        t.AppendRow(table.Row{key,value,arr.RelativeF[key], arr.Is[key], arr.Hs[key]})
		//fmt.Println(key, " ",value, arr.RelativeF[key], " ", arr.Is[key] , " ", arr.Hs[key])
	}
    //fmt.Println("tot ", count, " 1 ", "/ ", entropyTot)
    t.AppendRow(table.Row{"tot",count,"1","/",entropyTot})

    t.SetStyle(table.StyleLight)

    t.Render()

}