package internal

type FrequencyArr struct{
	Arr map[string]int
}

func FrequencyArrInit() FrequencyArr{
	return FrequencyArr{
		Arr: make(map[string]int),
	}
}