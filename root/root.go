package root

func HundredTake(value string) []string {
	arr := make([]string, 100)
	for i := range arr {
		arr[i] = value
	}
	return arr
}

