package stream

// New creates a new Stream from a given slice
func Of[T any, U any, V any](d []T) Streamer[T, U, V] {
	return &Stream[T, U, V]{
		dt: d,
	}
}

func (s *Stream[T, U, V]) Collect() []T {
	return s.dt
}

// Join appends the given slice to the stream
func (s *Stream[T, U, V]) Join(t Streamer[T, U, V]) Streamer[T, U, V] {
	s.dt = append(s.dt, t.Collect()...)
	return s
}

// // Filter returns a new Stream with elements that satisfy the given predicate
func (s *Stream[T, U, V]) ForEach(f func(T)) Streamer[T, U, V] {
	for _, v := range s.dt {
		f(v)
	}
	return s
}

// // Filter returns a new Stream with elements that satisfy the given predicate
func (s *Stream[T, U, V]) Filter(f func(T) bool) Streamer[T, U, V] {
	var result []T
	for _, v := range s.dt {
		if f(v) {
			result = append(result, v)
		}
	}
	return &Stream[T, U, V]{
		dt: result,
	}
}

// Map returns a new Stream with elements that are the result of applying the given function to each element
func (s *Stream[T, U, V]) Map(f func(T) U) Streamer[U, T, V] {
	var result []U
	for _, v := range s.dt {
		result = append(result, f(v))
	}
	return &Stream[U, T, V]{
		dt: result,
	}
}

func (s *Stream[T, U, V]) Reduce(f func(T, V) V, initialValue V) V {
	result := initialValue
	for _, v := range s.dt {
		result = f(v, result)
	}
	return result
}
