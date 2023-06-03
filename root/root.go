package root

func TenTake(value string) []string {
	arr := make([]string, 10)
	for i := range arr {
		arr[i] = value
	}
	return arr
}
