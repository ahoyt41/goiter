package goiter

type MapFunc[A, B any] func(A) B

func Map[A, B any](iter Iterator[A], fn MapFunc[A, B]) Iterator[B] {
	return &iterator[B]{
		next: func() (B, bool) {
			var b B
			a, ok := iter.Next()
			if !ok {
				return b, false
			}
			return fn(a), true
		},
	}
}

type Predicate[A any] func(A) bool

func Filter[A any](iter Iterator[A], fn Predicate[A]) Iterator[A] {
	return &iterator[A]{
		next: func() (A, bool) {
			for {
				a, ok := iter.Next()
				if !ok {
					return a, false
				}
				if fn(a) {
					return a, true
				}
			}
		},
	}
}

type Reducer[A, B any] func(B, A) B

func Reduce[A, B any](iter Iterator[A], fn Reducer[A, B], init B) B {
	for {
		a, ok := iter.Next()
		if !ok {
			break
		}
		init = fn(init, a)
	}
	return init
}

func Collect[A any](iter Iterator[A]) []A {
	slice := []A{}
	for {
		item, ok := iter.Next()
		if !ok {
			break
		}
		slice = append(slice, item)
	}
	return slice
}

func ForEach[A any](iter Iterator[A], fn func(A) error) error {
	for {
		item, ok := iter.Next()
		if !ok {
			break
		}
		if err := fn(item); err != nil {
			return err
		}
	}
	return nil
}
