package brackets

import "strings"

const testVersion = 3

var openCloseMappings = map[rune]rune{
	'{': '}',
	'[': ']',
	'(': ')',
}

type Stack []rune

func (s *Stack) Put(r rune) {
	*s = append((*s), r)
}

func (s *Stack) Clear() {
	*s = Stack{}
}

func (s *Stack) Pop() rune {
	if len(*s) == 0 {
		return 0
	}

	d := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return d
}

func (s Stack) Empty() bool {
	return len(s) == 0
}

type BracketType int

const (
	Open BracketType = iota
	Close
	Invalid
)

func Bracket(input string) (bool, error) {
	var stack Stack
	var popping bool

	for _, r := range input {
		switch getBracketType(r) {
		case Close:
			popping = true
		case Open:
			popping = false
			stack.Put(r)
		default:
			continue
		}

		if popping {
			poppedRune := stack.Pop()
			expected, ok := openCloseMappings[poppedRune]

			if !ok || expected != r {
				return false, nil
			}
		}
	}

	if stack.Empty() {
		return true, nil
	}

	return false, nil
}

func getBracketType(r rune) BracketType {
	closeBrackets := "}])"
	openBrackets := "{[("

	if strings.ContainsRune(closeBrackets, r) {
		return Close
	} else if strings.ContainsRune(openBrackets, r) {
		return Open
	}

	return Invalid
}
