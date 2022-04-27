package functor

type FunctorFmap[A any] func(A) interface{}

type Functor[A any] interface {
	Fmap(FunctorFmap[A]) Functor[A]
}

type FunctorList[A any] []A

func (f *FunctorList[A]) Fmap(fn FunctorFmap[A]) Functor[interface{}] {
	result := FunctorList[interface{}]{}

	for _, entry := range *f {
		result = append(result, fn(entry))
	}

	return &result
}
