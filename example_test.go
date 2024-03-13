package stream_test

import (
	"fmt"

	"stream"
)

func ExampleStream() {
	numberOfAllLettersInThreeLetterWords := stream.
		Of[string, int, int]([]string{"foo", "bar", "baz", "foobar"}).
		Join(stream.Of[string, int, int]([]string{"bat", "cat", "dog", "elephant"})).
		Map(func(s string) int { return len(s) }).
		Filter(func(i int) bool { return i == 3 }).
		Reduce(func(i int, j int) int {
			return i + j
		}, 0)
	fmt.Println(numberOfAllLettersInThreeLetterWords)
	// Output: 18
}
