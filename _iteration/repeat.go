package iteration

func Repeat(char string, n int) string {
	repeated := ""
	for i := 0; i < n; i++ {
		repeated += char
	}
	return repeated
}
