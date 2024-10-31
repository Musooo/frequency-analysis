package internal

import (
	"math"
)

func CharFrequency(text *string, arr FrequencyArr) (FrequencyArr, int) {
	count := 0
	for i := 0; i < len(*text); i++ {
		if (*text)[i] >= 65 && (*text)[i] <= 90 || (*text)[i] >= 97 && (*text)[i] <= 122 {
			arr.Fs[string((*text)[i])] += 1
			count++
		}
	}

	return arr, count
}

func RelativeFrequency(Fs *map[string]int, count int, RelativeF *map[string]float64) {
	for key, value := range *Fs {
		(*RelativeF)[key] = float64(value) / float64(count)
	}
}

func IsCalc(RelativeF *map[string]float64, Is *map[string]int) {
	for key, value := range *RelativeF {
		(*Is)[key] = int(math.Log2(value))*-1
	}
}

func EntropyCalc(RelativeF *map[string]float64, Is *map[string]int, Hs *map[string]float64){
	for key := range *RelativeF {
		(*Hs)[key] = (*RelativeF)[key] * float64((*Is)[key])
	}
}