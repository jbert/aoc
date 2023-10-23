package y2022

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/jbert/aoc/bitfield"
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
	flowRate int16
}

type actionType int

const (
	TURN actionType = iota
	MOVE
	NONE
	ELETURN
	ELEMOVE
)

const INVALID_VID = -1

type Action struct {
	typ        actionType
	valveID    int8
	eleTyp     actionType
	eleValveID int8
}

func (a Action) String() string {
	b := &strings.Builder{}
	switch a.typ {
	case TURN:
		fmt.Fprintf(b, "You open valve %s", valveName(a.valveID))
	case MOVE:
		fmt.Fprintf(b, "You move to valve %s", valveName(a.valveID))
	case NONE:
		panic("can't have none on non-elephant")
	default:
		panic("wtf")
	}
	switch a.eleTyp {
	case ELETURN:
		fmt.Fprintf(b, "The elephant opens valve %s", valveName(a.eleValveID))
	case ELEMOVE:
		fmt.Fprintf(b, "The elephant moves to valve %s", valveName(a.eleValveID))
	case NONE:
		// Nothing
	default:
		panic("wtf")
	}
	return b.String()
}

func (s state) next(a Action) state {
	//	fmt.Printf("%s: OVP: %d (%s)\n", s.toString(), s.openValvePressure(valves), a)
	newState := s
	//	newState.open = slices.Clone(s.open)
	switch a.typ {
	case TURN:
		newState.open.Set(a.valveID)
	case MOVE:
		newState.location = a.valveID
	case NONE:
		panic("can't have none on non-elephant")
	default:
		panic("wtf")
	}
	switch a.eleTyp {
	case ELETURN:
		newState.open.Set(a.eleValveID)
	case ELEMOVE:
		newState.elephant = a.eleValveID
	case NONE:
		// Nothing
	default:
		panic("wtf")
	}
	return newState
}

type state struct {
	location     int8
	elephant     int8
	open         bitfield.B64
	bestPressure int16
}

func (s state) doingElephant() bool {
	return s.elephant != INVALID_VID
}

func (s state) toString() string {
	return fmt.Sprintf("%d|%d|%s", s.location, s.elephant, string(s.open))
}

func (s state) possibleActions() []Action {
	neighbours := valveNeighbours[s.location]
	actions := fun.Map(func(loc int8) Action {
		return Action{typ: MOVE, valveID: loc, eleTyp: NONE, eleValveID: INVALID_VID}
	}, neighbours)
	if !s.open.Get(s.location) && valves[s.location].flowRate > 0 {
		actions = append(actions, Action{typ: TURN, valveID: s.location, eleTyp: NONE, eleValveID: INVALID_VID})
	}

	if s.doingElephant() {
		baseActions := slices.Clone(actions)
		actions = nil
		eleNeighbours := valveNeighbours[s.elephant]
		for _, baseAction := range baseActions {
			for _, loc := range eleNeighbours {
				action := baseAction // shallow copy ok
				action.eleTyp = ELEMOVE
				action.eleValveID = loc
				actions = append(actions, action)
			}
			if !s.open.Get(s.elephant) && valves[s.elephant].flowRate > 0 && !(baseAction.typ == TURN && baseAction.valveID == s.elephant) {
				action := baseAction // shallow copy ok
				action.eleTyp = ELETURN
				action.eleValveID = s.elephant
				actions = append(actions, action)
			}
		}
	}
	return actions
}

func (s state) openValvePressure() int16 {
	pressure := int16(0)
	size := int8(len(valves))
	for valveID := int8(0); valveID < size; valveID++ {
		if s.open.Get(valveID) {
			pressure += valves[valveID].flowRate
		}
	}
	return pressure
}

/*
func (s state) openValveLabels() []string {
	var ovs []string
	for label, on := range s.open {
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
*/

var valves []Valve
var labelTovID map[string]int8
var vIDToLabel map[int8]string
var valveNeighbours [][]int8

func valveName(id int8) string {
	return vIDToLabel[id]
}
func labelID(label string) int8 {
	return labelTovID[label]
}

func (d *Day16) Run(out io.Writer, lines []string) error {
	edges := []graph.Edge[string]{}
	labelTovID = make(map[string]int8)
	vIDToLabel = make(map[int8]string)

	valveID := int8(0)
	for _, l := range lines {
		v, lineEdges := parseLine(l)
		valves = append(valves, v)
		labelTovID[v.label] = valveID
		vIDToLabel[valveID] = v.label
		valveID++

		edges = append(edges, lineEdges...)
	}

	g := graph.NewFromEdges(edges, true)
	fmt.Printf("G: %v\n", g)
	fmt.Printf("V: %v\n", valves)

	f, err := os.Create("tt.dot")
	if err != nil {
		return fmt.Errorf("Can't open png file: %w", err)
	}
	defer f.Close()
	g.ToDot(f, "tt")

	fmt.Printf("V: %v\n", valves)

	vertices := g.Vertices()
	valveNeighbours = make([][]int8, len(vertices))
	for _, valveLabel := range vertices {
		vID := labelID(valveLabel)
		valveNeighbours[vID] = fun.Map(labelID, g.Neighbours(valveLabel))
	}

	if err := d.run(false); err != nil {
		return fmt.Errorf("Part1: %w", err)
	}
	fmt.Printf("---------------------------------------\n")
	if err := d.run(true); err != nil {
		return fmt.Errorf("Part2: %w", err)
	}
	return nil
}

func (d *Day16) run(useElephant bool) error {

	start := state{
		location: labelID("AA"),
		//		open:         bitfield.New(len(valves)),
		elephant:     INVALID_VID,
		bestPressure: 0,
	}
	//	for label := range valves {
	//		start.open[label] = false
	//	}
	maxMinutes := 30
	if useElephant {
		start.elephant = labelID("AA")
		maxMinutes = 26
	}

	currentBests := make(map[state]int16)
	currentBests[start] = 0

	for minute := 1; minute <= maxMinutes; minute++ {
		start := time.Now()
		fmt.Printf("== Minute %d ==\n", minute)
		newCurrentBests := make(map[state]int16)
		for s, bestP := range currentBests {
			delete(currentBests, s)
			s.bestPressure = bestP
			//			fmt.Printf("== State %d ==\n", i)
			possActions := s.possibleActions()
			//action := possActions[len(possActions)-1]
			for _, action := range possActions {
				//	fmt.Printf("%s\n", action)
				//				fmt.Printf("%s\n", s.openValveString(valves))
				//	fmt.Printf("\n")
				nextState := s.next(action)
				nextState.bestPressure += s.openValvePressure()

				//				nextKey := nextState.toString()
				nextKey := nextState
				nextKey.bestPressure = 0

				existingBest := newCurrentBests[nextKey]
				newCurrentBests[nextKey] = max(existingBest, nextState.bestPressure)
			}
		}
		if len(currentBests) != 0 {
			panic("wtf")
		}
		currentBests = newCurrentBests
		fmt.Printf("%d states - %s\n", len(currentBests), time.Since(start))
		/*
			maxPressure := fun.Max(fun.Map(func(s state) int {
				fmt.Printf("%s: BP: %d\n", s.toString(), s.bestPressure)
				return s.bestPressure
			}, states))
			fmt.Printf("Max released: %d\n", maxPressure)
		*/
	}
	//	maxPressure := fun.Max(fun.Map(func(s state) int { return s.bestPressure }, states))
	maxPressure := int16(0)
	for _, bestP := range currentBests {
		if bestP > maxPressure {
			maxPressure = bestP
		}
	}
	fmt.Printf("Max released: %d\n", maxPressure)
	return nil
}

func parseLine(l string) (Valve, []graph.Edge[string]) {
	v := Valve{}
	v.label = l[6:8]
	l = l[23:]
	i := strings.Index(l, ";")
	v.flowRate = int16(num.MustAtoi(l[:i]))
	l = l[i:]
	l = l[18:]
	i = strings.Index(l, " ")
	l = l[i+1:]
	fmt.Printf("[%s]\n", l)
	dests := strings.Split(l, ", ")
	edges := fun.Map(func(s string) graph.Edge[string] { return graph.Edge[string]{From: v.label, To: s} }, dests)
	return v, edges
}
