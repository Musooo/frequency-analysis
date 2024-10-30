package internal



func CharFrequency(text *string, arr FrequencyArr) FrequencyArr{
	for i :=0; i<len(*text); i++{
		if (*text)[i]>=101 && (*text)[i]<=132 || (*text)[i]>=141 && (*text)[i]<=172{
			arr.Arr[string((*text)[i])] += 1
		}
	}

	return arr
}