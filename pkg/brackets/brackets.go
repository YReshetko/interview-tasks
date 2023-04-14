package brackets

import (
	"errors"
)

type Num interface {
	byte | int | rune
}

type stack[T Num] []T

func newStack[T Num]() *stack[T] {
	s := stack[T](make([]T, 0, 0))
	return &s
}

func (s *stack[T]) push(v T) {
	ns := append(*s, v)
	*s = ns
}

func (s *stack[T]) pop() (T, error) {
	ps := *s
	if len(ps) == 0 {
		return 0, errors.New("invalid state")
	}
	v, ps := ps[len(ps)-1], ps[:len(ps)-1]
	*s = ps
	return v, nil
}

func (s *stack[T]) peek() (T, error) {
	ps := *s
	if len(ps) == 0 {
		return 0, errors.New("invalid state")
	}
	return ps[len(ps)-1], nil
}

func (s *stack[T]) isEmpty() bool {
	return len(*s) == 0
}

func IsValidWithStack(s string) bool {
	var openBrackets = map[int32]struct{}{'(': {}, '{': {}, '[': {}, '<': {}}
	var closeBrackets = map[int32]int32{')': '(', '}': '{', ']': '[', '>': '<'}
	st := newStack[rune]()
	for _, v := range s {
		if _, ok := openBrackets[v]; ok {
			st.push(v)
			continue
		}

		if b, ok := closeBrackets[v]; ok {
			ob, err := st.peek()
			if err != nil || ob != b {
				return false
			}
			_, _ = st.pop()
		}
	}
	return st.isEmpty()
}

func IsValidWithCounter(s string) bool {
	count := 0
	for _, v := range s {
		switch v {
		case '(':
			count++
		case ')':
			count--
			if count < 0 {
				return false
			}
		}
	}
	return count == 0
}
