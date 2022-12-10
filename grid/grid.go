package grid

// Grid is a zero-indexed rectangular grid with Get/Set
type Grid[A any] [][]A

func NewFromRows[A any](rows [][]A) Grid[A] {
	return rows
}

func New[A any](w, h int) Grid[A] {
	g := make([][]A, h)
	for y := range g {
		g[y] = make([]A, w)
	}
	return g
}

func (g Grid[A]) Get(x, y int) A {
	return g[y][x]
}

func (g Grid[A]) Set(x, y int, v A) {
	g[y][x] = v
}

func (g Grid[A]) Width() int {
	if len(g) == 0 {
		return 0
	}
	return len(g[0])
}

func (g Grid[A]) Height() int {
	return len(g)
}

func (g Grid[A]) ForEachVal(f func(A)) {
	g.ForEach(func(i, j int) {
		f(g.Get(i, j))
	})
}

func (g Grid[A]) ForEach(f func(int, int)) {
	for j, row := range g {
		for i := range row {
			f(i, j)
		}
	}
}

func (g Grid[A]) Combine(h Grid[A], f func(A, A) A) Grid[A] {
	r := New[A](g.Width(), g.Height())
	r.ForEach(func(i, j int) {
		gVal := g.Get(i, j)
		hVal := h.Get(i, j)
		r.Set(i, j, f(gVal, hVal))
	})
	return r
}