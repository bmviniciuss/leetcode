package problem20

type Stack struct {
	values []rune
}

func NewStack() *Stack {
	return &Stack{
		values: []rune{},
	}
}

func (s *Stack) Insert(val rune) {
	s.values = append(s.values, val)
}

func (s *Stack) Peek() (*rune, bool) {
	if len(s.values) == 0 {
		return new(rune), false
	}
	idx := len(s.values) - 1
	val := &s.values[idx]
	return val, true
}

func (s *Stack) Pop() (*rune, bool) {
	if len(s.values) == 0 {
		return new(rune), false
	}
	idx := len(s.values) - 1
	val := s.values[idx]
	s.values = s.values[:idx]
	return &val, true
}

func (s *Stack) Len() int {
	return len(s.values)
}

var (
	openPara     = rune('(')
	openCurly    = rune('{')
	openBrackt   = rune('[')
	closePara    = rune(')')
	closeCurly   = rune('}')
	closeBrackts = rune(']')
	closingMap   = map[rune]rune{
		closePara:    openPara,
		closeCurly:   openCurly,
		closeBrackts: openBrackt,
	}
)

func isOpeningRune(r rune) bool {
	return r == openPara || r == openCurly || r == openBrackt
}

func isClosingRune(r rune) bool {
	return r == closePara || r == closeCurly || r == closeBrackts
}

func isValid(s string) bool {
	l := len(s)
	if l == 0 {
		return true
	}
	if l%2 != 0 {
		return false
	}

	stack := NewStack()
	for _, letterRune := range s {
		if isOpeningRune(letterRune) {
			stack.Insert(letterRune)
			continue
		}
		if isClosingRune(letterRune) {
			complement := closingMap[letterRune]
			top, ok := stack.Peek()
			if !ok {
				return false // got closing but stack is empty
			}
			if *top != complement {
				return false // does not match with stack top
			}
			stack.Pop()
		}
	}
	return stack.Len() == 0
}
