package functor_test

import (
	"testing"

	"github.com/mrgleam/go-maybe/functor"
	"github.com/stretchr/testify/require"
)

func TestFunctorListFmapSameType(t *testing.T) {
	var l functor.FunctorList[int]
	l = []int{1, 2}

	add1 := func(a int) interface{} {
		return a + 1
	}

	result := l.Fmap(add1)
	require.Equal(t, result, &functor.FunctorList[interface{}]{2, 3})
}

func TestFunctorListFmapDiffType(t *testing.T) {
	var l functor.FunctorList[string]
	l = []string{"hello", "world!"}

	strLen := func(a string) interface{} {
		return len(a)
	}

	result := l.Fmap(strLen)
	require.Equal(t, result, &functor.FunctorList[interface{}]{5, 6})
}
