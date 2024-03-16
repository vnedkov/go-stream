# go-stream

This package provides functionality similar to Java Stream API in Go

## Receiver type

Because Go does not support type parameters in methods yet [issue 49085](https://github.com/golang/go/issues/49085), all type parameters must be declared in the receiver type. For now three type parameters need to be defined:
* The type of the stream element
* The type of the result of `Map` method. If you are not planning to transform one type into another, the type parameter can be `any`
* The result type of `Reduce` method. If Reduce is not used, the type parameter can be `any`

Simple example when `Map` and `Reduce are not used 
```
	sum := stream.Of[int, any, any]([]int{1, 2, 3, 4, 5}).
		ForEach(func(i int) { sum += i })
	assert.Equal(t, 15, sum)
```

Example of using `Map` and `Reduce`
```
	letterCount := Of[string, int, int]([]string{"foo", "bar"}).
		Map(func(s string) int {return len(s)}).
		Reduce(func(i1, i2 int) int {return i1 + i2}, 0)
	assert.Equal(t, 6, letterCount)        
```
