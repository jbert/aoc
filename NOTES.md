## Day 17

Lots of off-by-one errors. Having the visualisation is very useful.

Just realised I need to check left and right mobility as well as down.

Is there a good abstraction here?
    - yep, check the overlap bitmap against the equivalent bitmap

Need to check all rows/cols of the piece, going L and R and down.

## Day 16

Stopped on this at the time. Did a brute force search, but tried to cull the
multiplicity over time by noting that once the same set of valves were open,
the only state which mattered was the amount of pressure release so far.

Had an annoying bug due to shallow copy of the state - maps.Clone() for the
win.

Part 2 needs optimisation (or a better algorithm).

Baseline (part 1 only):
19.85user 1.42system 0:13.42elapsed 158%CPU (0avgtext+0avgdata 2352868maxresident)k

Profile (and thinking) suggests we're doing lots of map lookup and copying.
Lets use numbers instead of string labels (and correspondingly, arrays instead
of maps)

After using integers for most:
17.42user 0.56system 0:17.23elapsed 104%CPU (0avgtext+0avgdata 381424maxresident)k

Profile suggests action.toString is the bottlneck

Yes! 5-6x speedup:
3.37user 0.32system 0:03.24elapsed 114%CPU (0avgtext+0avgdata 247980maxresident)k

....and nowhere near good enough. Algorithmic improvement needed.


Try structural recursion on graph?
    - how to build recursion?
    - if have optimal soln on sub-graph, add in a single edge?
    - add in a new vertex (and set of connecting edges)
        - yes?
        - how do we manage minutes?
            - exhaustively?
                - choose to turn or not turn new vert (and only -- the minute
                  count for sub graph)
                - or choose to turn new graph, then move (and -= 2 the minute
                  count for sub graph)
                - gah - what about start node?

Tried using bitfield for []bool, using uint64 so struct is comparable and can
be used as a map key.

Further 3x speedup:
1.12user 0.27system 0:01.11elapsed 124%CPU (0avgtext+0avgdata 134472maxresident)k

System is almost fast enough, but OOMs...

Phew. Optimised mem use _juuuuust_ enough. Used ~20GB in the end.

853.81user 31.99system 14:30.66elapsed 101%CPU (0avgtext+0avgdata 20916112maxresident)k
63368inputs+15928outputs (461major+18540893minor)pagefaults 0swaps

Score another for brute force.

## Day 13

Tried a couple of things to make the parsing easier, but ultimately "char at a
time switch, recursively handle lists" gets me there with least friction.

## Day 12

Gah! Two big mistakes.

Firstly, I was using byte subtraction to check I was only moving up one step.
But golang bytes are unsigned! So moving down one step was up 254. D'oh.

Secondly, I was trying and trying to debug my astar, before I realised two
days later that the actual problem doesn't have the starting S in (0,0).
That's just mean.

## Day 11

### Part 2

So, I thought I just needed to go to int64. I was wrong. This wasn't obvious
as I wasn't dumping the intermediate values, just the results.

The bolded text to 'hint' that we needed to avoid big numbers led me to int64
but also didn't seem to quite fit. That - and the background concept of
divisibility - was enough to make me think of modular arithmetic.

I started thinking I could just have a modulus per monkey (with it's divisor)
but realised the passing between would break that, hence the LCM.

## Day 10

I do *love* implementing a virtual machine. The main thing to get right here
was ensuring we could inspect at ticks "between" instructions. I started off
thinking I'd use a goro to provide the instruction stream and drive the cpu
with a function call per tick, but that required keeping more state than I
wanted in the CPU (current instruction, cycles consumed etc) and the 'monitor'
callback suggested itself as a nicer alternative.

I also find myself thinking ahead to what the part 2 might be, and trying to
ensure I don't have a too-convoluted codebase/abstraction so that I'm well
prepared.

## Day 9

A fairly straightforward part 1, but I was flummoxed by part 2 until I did the
work to display intermediate stages exactly as in the text. This showed the
exact step which was failing, which - together with the 'be careful' hint that
more types of motion are possible - was enough.

## Day 8

A bit unsatisfying. I didn't see much opportunity for code sharing between the
different directions. I guess some kind of directional looping abstraction
might work?

## Day 5

This looked like it was going to be tricky to parse, but shortcuts worked
well, together with a nice 'stack' abstraction.

## Day 4

Handling the endpoints of the ranges can get quite crunchy - it is easy to
make mistakes on the different cases. Judicious use of reversing to consider
a.Covers(b) and b.Covers(a) simplifies things pretty well.

## Day 3

The 'set' abstraction makes this super easy.

## Day 2

### Part 1

Good bug. I hit the "Bad round" panic but it printed OK!

So I (sensibly) wrote some tests. Which all failed.

I had the RPS logic line as:

```
func (rps RPS) Beats(other RPS) bool {
	return (rps == 'R' && other == 'S') || (rps == 'P' && other == 'R') || (rps == 'S' && other == 'P')
}
```

Do you see the problem?

It may be obvious to you, but this compiles + runs and doesn't work. This is
the reason...

```
type RPS int

const (
	R RPS = iota
	P
	S
)
```
So...the logic above is testing the RPS values against char values, 'R' != R.

Silent autocasting between integer types? What is this, C?

(Real answer - I guess 'R' is a numeric literal, which are typeless in golang.
This allows you to do things like:
```
func DigToInt(digit byte) int {
    return c - '0'
}
```
but at what cost....

### Part 2

Yuck. Although this was a nice, straightforward way to do it, it was pretty
painful. With hindsight, something like mod-3 indexing into the string "RPS"
would be a lot nicer. Oh, just use the numbers `0,1,2` (meaning 'R', 'P', 'S') 
and then note (all arithmetic mod 3):

x beats x-1

Gah. That's a lot nicer.
