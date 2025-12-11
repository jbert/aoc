package pts

import "fmt"

type Rect struct {
	bl P2
	tr P2
}

func NewRect(p P2, q P2) Rect {
	l := min(p.X, q.X)
	r := max(p.X, q.X)
	b := min(p.Y, q.Y)
	t := max(p.Y, q.Y)
	return Rect{bl: P2{l, b}, tr: P2{r, t}}
}

func (r Rect) String() string {
	return fmt.Sprintf("BL %s TR %s", r.bl.String(), r.tr.String())
}

func CmpRect(a Rect, b Rect) int {
	c := CmpP2(a.bl, b.bl)
	if c != 0 {
		return c
	}
	return CmpP2(a.tr, b.tr)
}

// Inclusive area, so Rect{(3,3),(3,3)} has area 1
func (r Rect) Area() int {
	w := r.tr.X - r.bl.X + 1
	h := r.tr.Y - r.bl.Y + 1
	return w * h
}

func (r Rect) BL() P2 {
	return r.bl
}

func (r Rect) TL() P2 {
	l := r.bl.X
	t := r.tr.Y
	return P2{X: l, Y: t}
}

func (r Rect) TR() P2 {
	return r.tr
}

func (r Rect) BR() P2 {
	rr := r.tr.X
	b := r.bl.Y
	return P2{X: rr, Y: b}
}

// All pts in rect (including edges)
func (r Rect) AllPts() []P2 {
	// ps = make([]pts.P2, r.Width()+1*r.Height()+1)
	var ps []P2
	for i := r.bl.X; i <= r.tr.X; i++ {
		for j := r.bl.Y; j <= r.tr.Y; j++ {
			ps = append(ps, P2{i, j})
		}
	}
	// fmt.Printf("JB rect: %s\nps %v\n", r, ps)
	return ps
}
