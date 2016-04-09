package stringset

import (
	"fmt"
	"strings"
)

const testVersion = 3

type Set map[string]struct{}

func (s Set) String() string {
	str := "{"
	slc := s.Slice()

	var strElems []string

	for _, val := range slc {
		strElems = append(strElems, fmt.Sprintf("\"%v\"", val))
	}

	str += strings.Join(strElems, ", ")

	str += "}"

	return str
}

func New() Set {
	return make(Set)
}

func NewFromSlice(slc []string) Set {
	s := make(Set)
	for _, str := range slc {
		s[str] = struct{}{}
	}

	return s
}

func (s Set) Add(item string) {
	s[item] = struct{}{}
}

func (s Set) Delete(item string) {
	delete(s, item)
}

func (s Set) Has(item string) bool {
	_, ok := s[item]
	return ok
}

func (s Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Set) Len() int {
	return len(s)
}

func (s Set) Slice() []string {
	var slc []string
	for key := range s {
		slc = append(slc, key)
	}

	return slc
}

func Equal(s1, s2 Set) bool {
	if s1.Len() != s2.Len() {
		return false
	}

	for key := range s1 {
		if !s2.Has(key) {
			return false
		}
	}

	return true
}

func Subset(s1, s2 Set) bool {
	for key := range s1 {
		if !s2.Has(key) {
			return false
		}
	}

	return true
}

func Disjoint(s1, s2 Set) bool {
	for key := range s1 {
		if s2.Has(key) {
			return false
		}
	}

	return true
}

func Intersection(s1, s2 Set) Set {
	s := make(Set)

	for key := range s1 {
		if s2.Has(key) {
			s[key] = struct{}{}
		}
	}

	return s
}

func Union(s1, s2 Set) Set {
	s := make(Set)

	for key := range s1 {
		s[key] = struct{}{}
	}

	for key := range s2 {
		s[key] = struct{}{}
	}

	return s
}

func Difference(s1, s2 Set) Set {
	s := make(Set)

	for key := range s1 {
		if !s2.Has(key) {
			s[key] = struct{}{}
		}
	}

	return s
}

func SymmetricDifference(s1, s2 Set) Set {
	return Union(Difference(s1, s2), Difference(s2, s1))
}