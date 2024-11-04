package main

import (
	"time"
)

type TimeStack []time.Time

func (s *TimeStack) Push(value time.Time) {
	*s = append(*s, value)
}

func (s *TimeStack) Pop() (time.Time, bool) {
	if len(*s) == 0 {
		return time.Time{}, false
	}
	value := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return value, true
}

/* func main() {
	stack := TimeStack{}

	stack.Push(time.Now())
	time.Sleep(1 * time.Second)
	stack.Push(time.Now())

	for len(stack) > 0 {
		t, ok := stack.Pop()
		if ok {
			fmt.Println(t)
		}
	}
}
*/
