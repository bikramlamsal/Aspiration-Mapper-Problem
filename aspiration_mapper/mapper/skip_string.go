package mapper

import "unicode"

type SkipString struct {
	Data    []rune
	Skipper int
	idx     int
}

func (s *SkipString) TransformRune(pos int) {
	var x = s.Data[pos]
	if unicode.IsLetter(x) {
		if s.idx%s.Skipper == 0 {
			x = unicode.ToUpper(x)
		} else {
			x = unicode.ToLower(x)
		}
		s.idx++
	}
	s.Data[pos] = x
}

func (s *SkipString) GetValueAsRuneSlice() []rune {
	s.idx = 1
	return s.Data
}

func (s SkipString) String() string {
	return string(s.Data)
}
