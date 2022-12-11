package y2022

import (
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/jbert/aoc"
	"github.com/jbert/aoc/fun"
)

type Day11 struct{ Year }

func NewDay11() *Day11 {
	d := Day11{}
	return &d
}

func (d *Day11) Run(out io.Writer, lines []string) error {
	lgs := aoc.LineGroups(lines)
	monkeys := fun.Map(lineGroupToMonkey, lgs)
	fmt.Printf("%v\n", monkeys)

	rounds := 20
	inspected := make([]int, len(monkeys))
	for round := 1; round < rounds+1; round++ {
		for j, m := range monkeys {
			numInspected := m.Turn(monkeys)
			inspected[j] += numInspected
		}
		fmt.Printf("Round %d:\n", round)
		for i, m := range monkeys {
			fmt.Printf("Monkey: %d: %v\n", i, m.items)
		}
	}
	for i, ins := range inspected {
		fmt.Printf("Monkey %d inspected items %d times.\n", i, ins)
	}
	sort.Ints(inspected)
	inspected = fun.Reverse(inspected)
	fmt.Printf("Part 1: %d\n", inspected[0]*inspected[1])
	return nil
}

type Monkey struct {
	id    int
	items []int

	op       string
	arg      int
	argIsOld bool

	divisible int
	ifTrue    int
	ifFalse   int
}

func (m *Monkey) Turn(monkeys []*Monkey) int {
	for i := range m.items {
		m.items[i] = m.inspect(m.items[i])
		m.items[i] /= 3
		if m.items[i]%m.divisible == 0 {
			monkeys[m.ifTrue].items = append(monkeys[m.ifTrue].items, m.items[i])
		} else {
			monkeys[m.ifFalse].items = append(monkeys[m.ifFalse].items, m.items[i])
		}
	}
	l := len(m.items)
	m.items = nil
	return l
}

func (m *Monkey) inspect(old int) int {
	a := old
	b := m.arg
	if m.argIsOld {
		b = old
	}
	switch m.op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("wtf")
	}
}

func (m *Monkey) String() string {
	b := &strings.Builder{}
	fmt.Fprintf(b, "Monkey %d:\n", m.id)
	fmt.Fprintf(b, "  Starting itens: %v\n", m.items)
	argStr := "old"
	if !m.argIsOld {
		argStr = fmt.Sprintf("%d", m.arg)
	}
	fmt.Fprintf(b, "  Operation: new = old %s %s\n", m.op, argStr)
	fmt.Fprintf(b, "  Test: divisible by %d\n", m.divisible)
	fmt.Fprintf(b, "    If true: throw to monkey %d\n", m.ifTrue)
	fmt.Fprintf(b, "    If false: throw to monkey %d\n", m.ifFalse)
	return b.String()
}

func lineGroupToMonkey(lg []string) *Monkey {
	m := Monkey{}

	i := strings.Index(lg[1], ": ") + 2
	m.items = fun.Map(aoc.MustAtoi, strings.Split(lg[1][i:], ", "))

	i = strings.Index(lg[2], "new = old ") + 10
	bits := strings.Split(lg[2][i:], " ")
	fmt.Printf("B: %v\n", bits)
	m.op = bits[0]
	if bits[1] == "old" {
		m.argIsOld = true
	} else {
		m.arg = aoc.MustAtoi(bits[1])
	}

	i = strings.Index(lg[3], "divisible by ") + 13
	m.divisible = aoc.MustAtoi(lg[3][i:])

	i = strings.Index(lg[4], "monkey ") + 7
	m.ifTrue = aoc.MustAtoi(lg[4][i:])

	i = strings.Index(lg[5], "monkey ") + 7
	m.ifFalse = aoc.MustAtoi(lg[5][i:])
	return &m
}
