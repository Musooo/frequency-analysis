package internal

type FrequencyArr struct{
	Arr map[string]int
}

func FrequencyArrInit() FrequencyArr{
	mYarr := FrequencyArr{Arr: make(map[string]int)}
	for i := 'A'; i <= 'Z'; i++ {
		mYarr.Arr[string(i)] = 0
    }

	for i := 'a'; i <= 'z'; i++ {
		mYarr.Arr[string(i)] = 0
    }

	return mYarr
}