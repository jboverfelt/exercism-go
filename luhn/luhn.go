package luhn

import (
	"strconv"
	"unicode"
)

var fullWidthMap = map[rune]int{
	0xFF10: 0,
	0xFF11: 1,
	0xFF12: 2,
	0xFF13: 3,
	0xFF14: 4,
	0xFF15: 5,
	0xFF16: 6,
	0xFF17: 7,
	0xFF18: 8,
	0xFF19: 9,
}

var halfWidthMap = map[int]rune{
	0: 0xFF10,
	1: 0xFF11,
	2: 0xFF12,
	3: 0xFF13,
	4: 0xFF14,
	5: 0xFF15,
	6: 0xFF16,
	7: 0xFF17,
	8: 0xFF18,
	9: 0xFF19,
}

// FilterDigit returns a slice of all the integers
// in the string, converting to half width when necessary
// It also returns a boolean flag indicating whether or not
// there were full width integers present
func FilterDigit(str string) ([]int, bool) {
	var ret []int
	var containsFullWidth bool

	for _, r := range str {
		if unicode.IsDigit(r) {
			mappedInt, ok := fullWidthMap[r]
			if ok {
				ret = append(ret, mappedInt)
				containsFullWidth = true
			} else {

				digit, err := strconv.Atoi(string(r))

				if err != nil {
					panic(err)
				}

				ret = append(ret, digit)
			}
		}
	}

	return ret, containsFullWidth
}

func Valid(str string) bool {
	// we dont' care about full width integers
	// for validation purposes
	digits, _ := FilterDigit(str)

	if len(digits) < 2 {
		return false
	}

	sum := sumDigits(digits, false)

	return sum%10 == 0
}

func AddCheck(str string) string {
	digits, containsFullWidth := FilterDigit(str)

	sum := sumDigits(digits, true)

	checkDigit := (sum * 9) % 10

	if containsFullWidth {
		return str + string(halfWidthMap[checkDigit])
	}

	return str + strconv.Itoa(checkDigit)
}

func sumDigits(digits []int, everyOtherDigit bool) int {
	var sum int
	for i := (len(digits) - 1); i > -1; i-- {
		if !everyOtherDigit {
			sum += digits[i]
			everyOtherDigit = true
		} else {
			curDigit := digits[i] * 2
			if curDigit > 9 {
				curDigit -= 9
			}

			sum += curDigit

			everyOtherDigit = false
		}
	}

	return sum
}
