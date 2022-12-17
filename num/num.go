package num

import (
	"fmt"
	"strconv"
)

func Sign(x int) int {
	if x < 0 {
		return -1
	} else if x > 0 {
		return +1
	} else {
		return 0
	}
}

func IntAbs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func MustAtoi64(l string) int64 {
	n, err := strconv.ParseInt(l, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Line [%s] failed to be a number: %s", l, err))
	}
	return n
}

func MustAtoi(l string) int {
	n, err := strconv.Atoi(l)
	if err != nil {
		panic(fmt.Sprintf("Line [%s] failed to be a number: %s", l, err))
	}
	return n
}
