package maybe

// TODO: waiting next version of Golang support Just[T] | Nothing
type Maybe[T any] interface{}

type Just[T any] struct {
	Value T
}

type Nothing struct{}

// IsJust function returns true if its argument is of the Just[T]
func IsJust[T any](m Maybe[T]) bool {
	switch m.(type) {
	case Just[T]:
		return true
	}
	return false
}

// IsNothing function returns true if its argument is Nothing
func IsNothing[T any](m Maybe[T]) bool {
	switch m.(type) {
	case Nothing:
		return true
	}
	return false
}

// FromMaybe function takes a default value and a Maybe value. if the Maybe is Nothing, it returns the default value; otherwise, it returns the value contained in the Maybe.
func FromMaybe[T any](init T, m Maybe[T]) T {
	switch m.(type) {
	case Just[T]:
		return m.(Just[T]).Value
	}
	return init
}

// Map function returns the result of applying f to m if it's not empty. Otherwise returns Nothing
func Map[I, O any](m Maybe[I], f func(I) O) Maybe[O] {
	switch m.(type) {
	case Just[I]:
		return Just[O]{Value: f(m.(Just[I]).Value)}
	}
	return Nothing{}
}

// FlatMap function returns the result of applying f function must return Maybe to m if it's not empty. Orhterwise returns Nothing
func FlatMap[I, O any](m Maybe[I], f func(I) Maybe[O]) Maybe[O] {
	switch m.(type) {
	case Just[I]:
		return f(m.(Just[I]).Value)
	}
	return Nothing{}
}
