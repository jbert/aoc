package y2022

import (
	"fmt"
	"io"
	"maps"
	"sort"
	"strings"

	"github.com/jbert/aoc/graph"
	"github.com/jbert/aoc/num"
	"github.com/jbert/fun"
)

type Day16 struct{ Year }

func NewDay16() *Day16 {
	d := Day16{}
	return &d
}

type Valve struct {
	label    string
	flowRate int
	on       bool
}

type actionType int

const (
	TURN actionType = iota
	MOVE
)

type Action struct {
	typ   actionType
	label string
}

func (a Action) String() string {
	switch a.typ {
	case TURN:
		return fmt.Sprintf("You open valve %s", a.label)
	case MOVE:
		return fmt.Sprintf("You move to valve %s", a.label)
	default:
		panic("wtf")
	}
}

func (s state) next(a Action, valves map[string]Valve) state {
	//	fmt.Printf("%s: OVP: %d (%s)\n", s.toString(), s.openValvePressure(valves), a)
	newState := s
	newState.on = maps.Clone(s.on)
	switch a.typ {
	case TURN:
		newState.on[a.label] = true
	case MOVE:
		newState.location = a.label
	default:
		panic("wtf")
	}
	newState.bestPressure += s.openValvePressure(valves)
	//	fmt.Printf("%s: OVP: %d\n", newState.toString(), newState.openValvePressure(valves))
	return newState
}

type state struct {
	location     string
	on           map[string]bool
	bestPressure int
}

func (s state) toString() string {
	return s.location + "|" + strings.Join(s.openValveLabels(), "")
}

func (s state) possibleActions(g *graph.Graph[string], valves map[string]Valve) []Action {
	neighbours := g.Neighbours(s.location)
	actions := fun.Map(func(label string) Action { return Action{MOVE, label} }, neighbours)
	if !s.on[s.location] && valves[s.location].flowRate > 0 {
		actions = append(actions, Action{TURN, s.location})
	}
	return actions
}

func (s state) openPressure(valves map[string]Valve) int {
	labels := s.openValveLabels()
	return fun.Sum(fun.Map(func(label string) int { return valves[label].flowRate }, labels))
}

func (s state) openValvePressure(valves map[string]Valve) int {
	olabels := s.openValveLabels()
	fw := func(label string) int { return valves[label].flowRate }
	return fun.Sum(fun.Map(fw, olabels))
}

func (s state) openValveLabels() []string {
	var ovs []string
	for label, on := range s.on {
		if on {
			ovs = append(ovs, label)
		}
	}
	sort.Strings(ovs)
	return ovs
}

func (s state) openValveString(valves map[string]Valve) string {
	olabels := s.openValveLabels()
	switch len(olabels) {
	case 0:
		return "No valves are open."
	case 1:
		return fmt.Sprintf("Valve %s is open, releasing %d pressure.", olabels[0], s.openValvePressure(valves))
	default:
		commaStr := strings.Join(olabels[0:len(olabels)-1], ", ")
		return fmt.Sprintf("Valves %s and %s are open, releasing %d pressure.", commaStr, olabels[len(olabels)-1], s.openValvePressure(valves))
	}
}

func (d *Day16) Run(out io.Writer, lines []string) error {
	edges := []graph.Edge[string]{}
	valves := make(map[string]Valve)
	for _, l := range lines {
		v, lineEdges := parseLine(l)
		valves[v.label] = v
		edges = append(edges, lineEdges...)
	}
	g := graph.NewFromEdges(edges, true)
	fmt.Printf("G: %v\n", g)
	fmt.Printf("V: %v\n", valves)

	start := state{
		location:     "AA",
		on:           make(map[string]bool),
		bestPressure: 0,
	}
	for label := range valves {
		start.on[label] = false
	}

	states := []state{start}

	for minute := 1; minute <= 30; minute++ {
		fmt.Printf("== Minute %d ==\n", minute)
		nextStates := make(map[string]state)
		for _, s := range states {
			//			fmt.Printf("== State %d ==\n", i)
			possActions := s.possibleActions(g, valves)
			//action := possActions[len(possActions)-1]
			for _, action := range possActions {
				//				fmt.Printf("%s\n", action)
				//				fmt.Printf("%s\n", s.openValveString(valves))
				//				fmt.Printf("\n")
				nextState := s.next(action, valves)
				nextKey := nextState.toString()
				existingState, ok := nextStates[nextKey]
				if ok {
					nextState.bestPressure = max(nextState.bestPressure, existingState.bestPressure)
				}
				nextStates[nextState.toString()] = nextState // Collapse identical states, keeping best pressure
			}
		}
		states = []state{}
		for _, s := range nextStates {
			states = append(states, s)
		}
		fmt.Printf("%d states\n", len(states))
		/*
			maxPressure := fun.Max(fun.Map(func(s state) int {
				fmt.Printf("%s: BP: %d\n", s.toString(), s.bestPressure)
				return s.bestPressure
			}, states))
			fmt.Printf("Max released: %d\n", maxPressure)
		*/
	}
	maxPressure := fun.Max(fun.Map(func(s state) int { return s.bestPressure }, states))
	fmt.Printf("Max released: %d\n", maxPressure)
	return nil
}

func parseLine(l string) (Valve, []graph.Edge[string]) {
	v := Valve{}
	v.label = l[6:8]
	l = l[23:]
	i := strings.Index(l, ";")
	v.flowRate = num.MustAtoi(l[:i])
	l = l[i:]
	l = l[18:]
	i = strings.Index(l, " ")
	l = l[i+1:]
	fmt.Printf("[%s]\n", l)
	dests := strings.Split(l, ", ")
	edges := fun.Map(func(s string) graph.Edge[string] { return graph.Edge[string]{From: v.label, To: s} }, dests)
	return v, edges
}
