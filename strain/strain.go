package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(pred func(int) bool) Ints {
	var newInts Ints

	for _, val := range i {
		if pred(val) {
			newInts = append(newInts, val)
		}
	}

	return newInts
}

func (i Ints) Discard(pred func(int) bool) Ints {
	var newInts Ints

	for _, val := range i {
		if !pred(val) {
			newInts = append(newInts, val)
		}
	}

	return newInts
}

func (l Lists) Keep(pred func([]int) bool) Lists {
	var newLists Lists

	for _, val := range l {
		if pred(val) {
			newLists = append(newLists, val)
		}
	}

	return newLists
}

func (s Strings) Keep(pred func(string) bool) Strings {
	var newStrings Strings

	for _, val := range s {
		if pred(val) {
			newStrings = append(newStrings, val)
		}
	}

	return newStrings
}