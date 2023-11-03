package goiter

type Pair[A, B any] struct {
	First  A
	Second B
}

func (p Pair[A, B]) Values() (A, B) { return p.First, p.Second }

type Result[A any] Pair[A, error]

func Ok[A any](a A) Result[A] {
	return Result[A]{a, nil}
}

func Err[A any](err error) Result[A] {
	var a A
	return Result[A]{a, err}
}

func (r Result[A]) Unwrap() (A, error) { return r.First, r.Second }

type Option[A any] Pair[A, bool]

func Some[A any](a A) Option[A] {
	return Option[A]{a, true}
}

func None[A any]() Option[A] {
	var a A
	return Option[A]{a, false}
}

func (opt Option[A]) Unwrap() (A, bool) { return opt.First, opt.Second }
