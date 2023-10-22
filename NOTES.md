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
