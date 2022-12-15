package y2022

import (
	"fmt"
	"io"
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
	for _, p := range pairs {
		fmt.Printf("%s\n%s\n\n", p.a, p.b)
	}
	return nil
}

type Pair struct {
	a, b *Node
}

func (p Pair) String() string {
	return fmt.Sprintf("PAIR:\n%s\n%s\n", p.a, p.b)
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

	current := &Node{}

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
