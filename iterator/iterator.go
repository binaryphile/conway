package iterator

type OfString = Iter[string]

type Iter[T any] func() (T, Iter[T])

func FromSlice[T any](ts []T) Iter[T] {
	if len(ts) == 0 {
		return nil
	}

	i := 0

	var next Iter[T]
	next = func() (T, Iter[T]) {
		if i == len(ts)-1 {
			return ts[i], nil
		}

		t := ts[i]
		i += 1
		return t, next
	}

	return next
}
