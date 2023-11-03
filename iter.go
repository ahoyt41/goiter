package goiter

type Iterator[T any] interface {
	Next() (T, bool)
}

type iterator[T any] struct {
	next func() (T, bool)
}

func (iter *iterator[T]) Next() (T, bool) { return iter.next() }

type channelIter[T any] struct {
	itemChan chan T
}

func (iter *channelIter[T]) Next() (T, bool) {
	var t T
	select {
	case t = <-iter.itemChan:
		return t, true
	default:
		return t, false
	}
}

func SliceIter[T any](slice []T) Iterator[T] {
	i := 0
	return &iterator[T]{
		next: func() (T, bool) {
			var t T
			if i >= len(slice) {
				return t, false
			}
			t = slice[i]
			i++
			return t, true
		},
	}
}

func Range(start, end, step int) Iterator[int] {
	current := start
	return &iterator[int]{
		next: func() (int, bool) {
			if current >= end {
				return current, false
			}
			var val int
			val, current = current, current+step
			return val, true
		},
	}
}
