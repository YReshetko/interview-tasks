package brackets_test

import (
	"fmt"
	"github.com/YReshetko/interview-tasks/pkg/brackets"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValidWithStack(t *testing.T) {
	toTest := []struct {
		value    string
		expected bool
	}{
		{")(", false},
		{"({}){[]<>}", true},
		{"(()(()()()))(((()))", false},
		{"((((()(()()()))(((()))))))))", false},
		{"(((((((((((((((((((((((((((()(()()()))(((())))))))))))))))))))))))))))))))", false},
	}

	for i, test := range toTest {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, test.expected, brackets.IsValidWithStack(test.value))
		})
	}
}

func TestIsValidWithCounter(t *testing.T) {
	toTest := []struct {
		value    string
		expected bool
	}{
		{")(", false},
		{"()", true},
		{"(()(()()()))(((()))", false},
		{"((((()(()()()))(((()))))))))", false},
		{"(((((((((((((((((((((((((((()(()()()))(((())))))))))))))))))))))))))))))))", false},
	}

	for i, test := range toTest {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, test.expected, brackets.IsValidWithCounter(test.value))
		})
	}
}
