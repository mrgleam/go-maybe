package maybe_test

import (
	"testing"

	"github.com/mrgleam/go-maybe/maybe"
	"github.com/stretchr/testify/require"
)

func TestIsJust(t *testing.T) {
	j := maybe.IsJust[int](maybe.Just[int]{Value: 5})
	n := maybe.IsJust[int](maybe.Nothing{})
	require.True(t, j)
	require.False(t, n)
}

func TestIsNothing(t *testing.T) {
	j := maybe.IsNothing[int](maybe.Just[int]{Value: 5})
	n := maybe.IsNothing[interface{}](maybe.Nothing{})
	require.False(t, j)
	require.True(t, n)
}

func TestFromMaybeNothing(t *testing.T) {
	nothing := func() maybe.Maybe[interface{}] {
		return maybe.Nothing{}
	}()

	actual := maybe.FromMaybe("", nothing)

	require.Equal(t, "", actual)
}

func TestFromMaybeJust(t *testing.T) {
	just := func() maybe.Maybe[string] {
		return maybe.Just[string]{Value: "Hello, world"}
	}()

	actual := maybe.FromMaybe("", just)

	require.Equal(t, "Hello, world", actual)
}

func TestMapNothing(t *testing.T) {
	nothing := func() maybe.Maybe[int] {
		return maybe.Nothing{}
	}()

	double := func(a int) int {
		return a * 2
	}

	actual := maybe.Map(nothing, double)

	require.Equal(t, maybe.Nothing{}, actual)
}

func TestMapJust(t *testing.T) {
	just := func() maybe.Maybe[int] {
		return maybe.Just[int]{Value: 5}
	}()

	double := func(a int) int {
		return a * 2
	}

	actual := maybe.Map(just, double)

	require.Equal(t, maybe.Just[int]{Value: 10}, actual)
}

func TestFlatMapNothing(t *testing.T) {
	nothing := func() maybe.Maybe[int] {
		return maybe.Nothing{}
	}()

	double := func(a int) maybe.Maybe[int] {
		return maybe.Just[int]{Value: a * 2}
	}

	actual := maybe.FlatMap(nothing, double)

	require.Equal(t, maybe.Nothing{}, actual)
}

func TestFlatMapJust(t *testing.T) {
	just := func() maybe.Maybe[int] {
		return maybe.Just[int]{Value: 5}
	}()

	double := func(a int) maybe.Maybe[int] {
		return maybe.Just[int]{Value: a * 2}
	}

	actual := maybe.FlatMap(just, double)

	require.Equal(t, maybe.Just[int]{Value: 10}, actual)
}
