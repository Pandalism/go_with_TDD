package iteration

// Repeat is a function that repeats the input character a given amount of times
func Repeat(character string, iterations int) string {

	var repeated string
	for i := 0; i < iterations; i++ {
		repeated += character
	}
	return repeated
}
