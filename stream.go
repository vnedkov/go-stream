// Package stream implements functionality similar to Java stream API
// It uses a generic type receiver with tree type parameters

package stream

// Stream is a generic structure. A pointer to it us used as a type receiver
// First type parameter defines the type of stream elements
// Second Type parameter defines the type of output elements of Map method. If Map is not used, "any" type can be used
// Third parameter defines the result type of Reduce method
type Stream[T any, U any, V any] struct {
	dt []T
}

// Streamer is an interface defining applicable operations that can be performed on a stream
type Streamer[T any, U any, V any] interface {
	// Collect returns all elements of the stream
	Collect() []T
	// ForEach executes a given function for all elements
	ForEach(f func(T)) Streamer[T, U, V]
	// Filter removes all the elements for which the given function returns false
	Filter(f func(T) bool) Streamer[T, U, V]
	// Join combines two streams in one
	Join(t ...Streamer[T, U, V]) Streamer[T, U, V]
	// Map uses the given function to transform all elements to a new type
	Map(f func(T) U) Streamer[U, T, V]
	// Reduce reduces all elements to a new type given a transformation function
	Reduce(f func(T, V) V, initialValue V) V
}
