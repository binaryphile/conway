package iterator

type OfString = Iterator[string]

type Iterator[T any] func() (T, Iterator[T])

func FromSlice[T any](ts []T) Iterator[T] {
	if len(ts) == 0 {
		return nil
	}

	i := 0

	var next Iterator[T]
	next = func() (T, Iterator[T]) {
		if i == len(ts)-1 {
			return ts[i], nil
		}

		t := ts[i]
		i += 1
		return t, next
	}

	return next
}
