package stream

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOf(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		new := Of[int, any, any]([]int{1, 2, 3})
		assert.Equal(t, new.Collect(), []int{1, 2, 3})
	})
	t.Run("string", func(t *testing.T) {
		new := Of[string, any, any]([]string{"one", "two", "three"})
		assert.Equal(t, new.Collect(), []string{"one", "two", "three"})
	})
}

func TestJoin(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		stream1 := Of[int, any, any]([]int{1, 2, 3})
		stream2 := Of[int, any, any]([]int{4, 5, 6})
		stream3 := Of[int, any, any]([]int{7, 8})
		assert.Equal(t, stream1.Join(stream2, stream3).Collect(), []int{1, 2, 3, 4, 5, 6, 7, 8})
	})

	t.Run("string", func(t *testing.T) {
		stream4 := Of[string, any, any]([]string{"foo", "bar"})
		stream5 := Of[string, any, any]([]string{"baz", "qux"})
		assert.Equal(t, stream4.Join(stream5).Collect(), []string{"foo", "bar", "baz", "qux"})
	})
}

func TestMap(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		stream1 := Of[int, string, any]([]int{1, 2, 3})
		assert.Equal(t, stream1.
			Map(func(x int) string { return strconv.Itoa(x) }).
			Collect(), []string{"1", "2", "3"})
	})

	t.Run("string", func(t *testing.T) {
		stream2 := Of[string, int, any]([]string{"foo", "bar"})
		f2 := func(x string) int { return len(x) }
		assert.Equal(t, stream2.Map(f2).Collect(), []int{3, 3})
	})
}

func TestFilter(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		stream1 := Of[int, any, any]([]int{1, 2, 3, 4, 5})
		assert.Equal(t, stream1.
			Filter(func(x int) bool { return x%2 == 0 }).
			Collect(), []int{2, 4})
	})

	t.Run("string", func(t *testing.T) {
		stream2 := Of[string, any, any]([]string{"foo", "bar", "baz", "foobar"})
		assert.Equal(t, stream2.
			Filter(func(x string) bool { return len(x) > 3 }).
			Collect(), []string{"foobar"})
	})
}

func TestForEach(t *testing.T) {
	var sum int
	Of[int, any, any]([]int{1, 2, 3, 4, 5}).
		ForEach(func(i int) { sum += i })
	assert.Equal(t, 15, sum)
}

func TestReduce(t *testing.T) {
	sum := Of[int, any, int]([]int{1, 2, 3, 4, 5}).
		Reduce(func(i1, i2 int) int { return i1 + i2 }, 0)
	assert.Equal(t, 15, sum)
}

func TestMapReduce(t *testing.T) {
	letterCount := Of[string, int, int]([]string{"foo", "bar"}).
		Map(func(s string) int { return len(s) }).
		Reduce(func(i1, i2 int) int { return i1 + i2 }, 0)
	assert.Equal(t, 6, letterCount)
}
