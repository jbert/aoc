package y2022

import (
	"fmt"
	"io"
	"sort"
	"strings"
	"unicode"

	"github.com/jbert/aoc"
	"github.com/jbert/aoc/fun"
	"github.com/jbert/aoc/num"
)

type Day13 struct{ Year }

func NewDay13() *Day13 {
	d := Day13{}
	return &d
}

func (d *Day13) Run(out io.Writer, lines []string) error {
	lgs := aoc.LineGroups(lines)
	pairs := fun.Map(lineGroupToPair, lgs)
	rightOrderSum := 0
	for i, p := range pairs {
		fmt.Printf("%s\n%s\n", p.a, p.b)
		fmt.Printf("Right order? %v\n", p.RightOrder())
		fmt.Printf("\n")
		if p.RightOrder() {
			rightOrderSum += i + 1
		}
	}
	fmt.Printf("Part 1: %d\n", rightOrderSum)

	packets := fun.Map(lineToPacket, fun.Filter(func(l string) bool { return len(l) > 0 }, lines))
	packets = append(packets, lineToPacket("[[2]]"))
	packets = append(packets, lineToPacket("[[6]]"))
	sort.Slice(packets, func(i, j int) bool {
		return packets[i].Cmp(packets[j]) == -1
	})

	div1 := 0
	div2 := 0
	for i, p := range packets {
		if p.String() == "[[2]]" {
			div1 = i + 1
		}
		if p.String() == "[[6]]" {
			div2 = i + 1
		}
	}
	fmt.Printf("Part 2: %d\n", div1*div2)

	return nil
}

type Pair struct {
	a, b *Node
}

func (p Pair) String() string {
	return fmt.Sprintf("PAIR:\n%s\n%s\n", p.a, p.b)
}

func (p Pair) RightOrder() bool {
	c := p.a.Cmp(p.b)
	if c == 0 {
		panic("same!")
	}
	return c == -1
}

func lineGroupToPair(lg []string) Pair {
	if len(lg) != 2 {
		panic("wtf")
	}
	return Pair{
		a: lineToPacket(lg[0]),
		b: lineToPacket(lg[1]),
	}
}

// Node is either/or. v is set iff children is nil
type Node struct {
	v        int
	children []*Node
}

func (n *Node) Cmp(m *Node) int {
	if n.children == nil && m.children == nil {
		if n.v < m.v {
			return -1
		}
		if n.v > m.v {
			return +1
		}
		return 0
	}
	if n.children == nil {
		nn := &Node{v: 0, children: []*Node{{v: n.v}}}
		return nn.Cmp(m)
	}
	if m.children == nil {
		mm := &Node{v: 0, children: []*Node{{v: m.v}}}
		return n.Cmp(mm)
	}
	// Both lists
	nc := n.children
	mc := m.children
	for {
		if len(nc) == 0 && len(mc) == 0 {
			return 0
		}
		if len(nc) == 0 {
			return -1
		}
		if len(mc) == 0 {
			return +1
		}
		c := nc[0].Cmp(mc[0])
		if c != 0 {
			return c
		}
		nc = nc[1:]
		mc = mc[1:]
	}
	panic("not reached")
}

func (n *Node) String() string {
	if n.children == nil {
		return fmt.Sprintf("%d", n.v)
	}
	return fmt.Sprintf("[%s]", strings.Join(fun.Map(func(n *Node) string { return n.String() }, n.children), ","))
}

func lineToPacket(l string) *Node {
	rest, node := parseList(l)
	if len(rest) != 0 {
		panic("wtf")
	}
	return node
}

func parseList(l string) (string, *Node) {
	if l[0] != '[' {
		panic(fmt.Sprintf("No opening brace: %v", l))
	}
	l = l[1:]

	current := &Node{children: make([]*Node, 0)}

NEXTCHAR:
	for {
		var node *Node
		switch l[0] {
		case '[':
			l, node = parseList(l)
			current.children = append(current.children, node)
		case ']':
			l = l[1:]
			break NEXTCHAR
		case ',':
			l = l[1:]
		default:
			//			fmt.Printf("digit?: %s\n", l)
			endDigit := -1
			for i, r := range l {
				endDigit = i
				if !unicode.IsDigit(r) {
					break
				}
			}
			//			fmt.Printf("number?: %d: %s\n", endDigit, l[:endDigit])
			n := num.MustAtoi(l[:endDigit])
			node = &Node{v: n, children: nil}
			current.children = append(current.children, node)
			l = l[endDigit:]
		}
	}
	return l, current
}
