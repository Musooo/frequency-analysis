package internaltest_test

import (
	"testing"
	"github.com/musooo/frequency-analysis/internal"
)

func TestCharFrequency(t *testing.T){
	t.Run("testing word counting",func(t *testing.T) {
		str :="ciao ,.,.,.,.<><>><>ciao\n"
		arr := internal.CharFrequency(&str,internal.FrequencyArrInit())
		if len(arr.Arr)!=52{
			t.Error("errore nel incremento struct")
		}
	})
}