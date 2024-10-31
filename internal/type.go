package internal

type FrequencyArr struct{
	Fs map[string]int
	RelativeF map[string]float64
	Is map[string]int
	Hs map[string]float64
}

func FrequencyArrInit() FrequencyArr{
	return FrequencyArr{
		Fs: make(map[string]int),
		RelativeF: make(map[string]float64),
		Is: make(map[string]int),
		Hs: make(map[string]float64),
	}
}