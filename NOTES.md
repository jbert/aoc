## Day 2

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
