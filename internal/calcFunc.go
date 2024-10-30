package internal


func CharFrequency(text *string, arr FrequencyArr) FrequencyArr {
	for i := 0; i < len(*text); i++ {
		if (*text)[i] >= 65 && (*text)[i] <= 90 || (*text)[i] >= 97 && (*text)[i] <= 122 {
			arr.Arr[string((*text)[i])] += 1
		}
	}

	return arr
}