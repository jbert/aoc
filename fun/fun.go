package fun

import "golang.org/x/exp/constraints"

func Reverse[T any](l []T) []T {
	ll := len(l)
	rev := make([]T, ll)
	for i := range l {
		rev[ll-i-1] = l[i]
	}
	return rev
}

func Map[A any, B any](f func(A) B, as []A) []B {
	bs := make([]B, len(as))
	for i := range as {
		bs[i] = f(as[i])
	}
	return bs
}

func Filter[A any](pred func(A) bool, as []A) []A {
	fs := make([]A, 0)
	for _, a := range as {
		if pred(a) {
			fs = append(fs, a)
		}
	}
	return fs
}

type Number interface {
	constraints.Integer | constraints.Float
}

func Iota[A Number](start A, count int) []A {
	l := make([]A, count)
	current := start
	for i := range l {
		l[i] = current
		current += 1
	}
	return l
}

func Sum[A Number](as []A) A {
	var zero A
	return Foldl(func(a, b A) A { return a + b }, zero, as)
}

func Foldl[A any, B any](f func(A, B) B, acc B, as []A) B {
	for _, a := range as {
		acc = f(a, acc)
	}
	return acc
}

func SplitBy[A any](as []A, n int) [][]A {
	lol := make([][]A, len(as)/n)
	for i := range lol {
		lol[i] = as[i*n : (i+1)*n]
	}
	return lol
}
