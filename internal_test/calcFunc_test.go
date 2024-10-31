package internaltest_test

import (
	"testing"
	"github.com/musooo/frequency-analysis/internal"
)

func TestCharFrequency(t *testing.T){
	t.Run("testing word counting",func(t *testing.T) {
		str :="ciao ,.,.,.,.<><>><>ciao\n"
		arr, count := internal.CharFrequency(&str,internal.FrequencyArrInit())
		if len(arr.Fs)!=4{
			t.Error("errore nel incremento struct")
		}
		if count!=8{
			t.Error("errore nel calcolo len")
		}
	})
}