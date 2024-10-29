package pkgtest_test

import (
	"testing"
	"github.com/musooo/frequency-analysis/pkg"
)

func TestReadFileAndReturnString(t *testing.T){
	t.Run("testing file reading", func(t *testing.T) {
		if pkg.ReadFileAndReturnString("prova.txt")!="ciao"{
			t.Error("error reading file")
		}
	})
}