package y2022

import (
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/jbert/aoc"
	"github.com/jbert/fun"
	"github.com/jbert/aoc/num"
)

type Day11 struct{ Year }

func NewDay11() *Day11 {
	d := Day11{}
	return &d
}

func (d *Day11) Run(out io.Writer, lines []string) error {
	lgs := aoc.LineGroups(lines)

	fmt.Printf("Part 1: %d\n", DoPart(lgs, 20, true))
	fmt.Printf("Part 2: %d\n", DoPart(lgs, 10000, false))
	return nil
}

func DoPart(lgs [][]string, rounds int, doDiv3 bool) int64 {
	monkeys := fun.Map(lineGroupToMonkey, lgs)
	//	fmt.Printf("%v\n", monkeys)

	lcm := fun.Prod(fun.Map(func(m *Monkey) int64 { return m.divisible }, monkeys))
	fmt.Printf("LCM: %d\n", lcm)

	inspected := make([]int64, len(monkeys))
	for round := 1; round < rounds+1; round++ {
		for j, m := range monkeys {
			numInspected := m.Turn(monkeys, doDiv3, lcm)
			inspected[j] += numInspected
		}
		if round%1000 == 0 {
			fmt.Printf("Round %d:\n", round)
			for i, m := range monkeys {
				fmt.Printf("Monkey: %d: %v\n", i, m.items)
			}
		}
	}
	for i, ins := range inspected {
		fmt.Printf("Monkey %d inspected items %d times.\n", i, ins)
	}
	//	sort.Ints(inspected)
	sort.Slice(inspected, func(i, j int) bool { return inspected[i] < inspected[j] })
	inspected = fun.Reverse(inspected)
	return inspected[0] * inspected[1]
}

type Monkey struct {
	id    int
	items []int64

	op       string
	arg      int64
	argIsOld bool

	divisible int64
	ifTrue    int
	ifFalse   int
}

func (m *Monkey) Turn(monkeys []*Monkey, doDiv3 bool, lcm int64) int64 {
	for i := range m.items {
		m.items[i] = m.inspect(m.items[i])
		if doDiv3 {
			m.items[i] /= 3
		}
		if m.items[i]%m.divisible == 0 {
			monkeys[m.ifTrue].items = append(monkeys[m.ifTrue].items, m.items[i]%lcm)
		} else {
			monkeys[m.ifFalse].items = append(monkeys[m.ifFalse].items, m.items[i]%lcm)
		}
	}
	l := len(m.items)
	m.items = nil
	return int64(l)
}

func (m *Monkey) inspect(old int64) int64 {
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
	m.items = fun.Map(num.MustAtoi64, strings.Split(lg[1][i:], ", "))

	i = strings.Index(lg[2], "new = old ") + 10
	bits := strings.Split(lg[2][i:], " ")
	fmt.Printf("B: %v\n", bits)
	m.op = bits[0]
	if bits[1] == "old" {
		m.argIsOld = true
	} else {
		m.arg = num.MustAtoi64(bits[1])
	}

	i = strings.Index(lg[3], "divisible by ") + 13
	m.divisible = num.MustAtoi64(lg[3][i:])

	i = strings.Index(lg[4], "monkey ") + 7
	m.ifTrue = num.MustAtoi(lg[4][i:])

	i = strings.Index(lg[5], "monkey ") + 7
	m.ifFalse = num.MustAtoi(lg[5][i:])
	return &m
}
