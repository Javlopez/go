package iteration

const RepeatCount = 5

func Repeat(character string, times int) string {
	var repeated string
	counter := RepeatCount
	if times == 0 {
		counter = times
	}
	for i := 0; i < counter; i++ {
		repeated = repeated + character
	}

	return repeated
}
