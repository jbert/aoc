package y2022

import (
	"fmt"
	"io"
	"os"
	"slices"
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
}

type actionType int

const (
	TURN actionType = iota
	MOVE
	NONE
	ELETURN
	ELEMOVE
)

type vID int

const INVALID_VID = vID(-1)

type Action struct {
	typ        actionType
	valveID    vID
	eleTyp     actionType
	eleValveID vID
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
	newState.open = slices.Clone(s.open)
	switch a.typ {
	case TURN:
		newState.open[a.valveID] = true
	case MOVE:
		newState.location = a.valveID
	case NONE:
		panic("can't have none on non-elephant")
	default:
		panic("wtf")
	}
	switch a.eleTyp {
	case ELETURN:
		newState.open[a.eleValveID] = true
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
	location     vID
	elephant     vID
	open         []bool
	bestPressure int
}

func (s state) doingElephant() bool {
	return s.elephant != INVALID_VID
}

func (s state) toString() string {
	buf := make([]byte, len(s.open))
	for i, open := range s.open {
		buf[i] = '0'
		if open {
			buf[i] = '1'
		}
	}
	return fmt.Sprintf("%d|%d|%s", s.location, s.elephant, string(buf))
}

func (s state) possibleActions() []Action {
	neighbours := valveNeighbours[s.location]
	actions := fun.Map(func(loc vID) Action {
		return Action{typ: MOVE, valveID: loc, eleTyp: NONE, eleValveID: INVALID_VID}
	}, neighbours)
	if !s.open[s.location] && valves[s.location].flowRate > 0 {
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
			if !s.open[s.elephant] && valves[s.elephant].flowRate > 0 {
				action := baseAction // shallow copy ok
				action.eleTyp = ELETURN
				action.eleValveID = s.elephant
				actions = append(actions, action)
			}
		}
	}
	return actions
}

func (s state) openValvePressure() int {
	pressure := 0
	for valveID, open := range s.open {
		if open {
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
var labelTovID map[string]vID
var vIDToLabel map[vID]string
var valveNeighbours [][]vID

func valveName(id vID) string {
	return vIDToLabel[id]
}
func labelID(label string) vID {
	return labelTovID[label]
}

func (d *Day16) Run(out io.Writer, lines []string) error {
	edges := []graph.Edge[string]{}
	labelTovID = make(map[string]vID)
	vIDToLabel = make(map[vID]string)

	valveID := 0
	for _, l := range lines {
		v, lineEdges := parseLine(l)
		valves = append(valves, v)
		labelTovID[v.label] = vID(valveID)
		vIDToLabel[vID(valveID)] = v.label
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
	valveNeighbours = make([][]vID, len(vertices))
	for _, valveLabel := range vertices {
		vID := labelID(valveLabel)
		valveNeighbours[vID] = fun.Map(labelID, g.Neighbours(valveLabel))
	}

	if err := d.run(false); err != nil {
		return fmt.Errorf("Part1: %w", err)
	}
	fmt.Printf("---------------------------------------\n")
	//	if err := d.run(true); err != nil {
	//		return fmt.Errorf("Part2: %w", err)
	//	}
	return nil
}

func (d *Day16) run(useElephant bool) error {

	start := state{
		location:     labelID("AA"),
		open:         make([]bool, len(valves)),
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

	states := []state{start}

	for minute := 1; minute <= maxMinutes; minute++ {
		fmt.Printf("== Minute %d ==\n", minute)
		nextStates := make(map[string]state)
		for _, s := range states {
			//			fmt.Printf("== State %d ==\n", i)
			possActions := s.possibleActions()
			//action := possActions[len(possActions)-1]
			for _, action := range possActions {
				//	fmt.Printf("%s\n", action)
				//				fmt.Printf("%s\n", s.openValveString(valves))
				//	fmt.Printf("\n")
				nextState := s.next(action)
				nextState.bestPressure += s.openValvePressure()

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
