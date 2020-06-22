package _0_06

import (
	"fmt"
	"testing"
)

func patternMatching(pattern string, value string) bool {
	s := [2]string{"", ""}
	return solve(s, pattern, 0, value, 0)
}

func solve(s [2]string, pattern string, index1 int, value string, index2 int) bool {
	//匹配完成
	if index1 == len(pattern) && index2 == len(value) {
		return true
	}

	if index1 >= len(pattern) || index2 > len(value) {
		return false
	}
	num := pattern[index1] - 'a'
	if s[num] == "null" {
		for i := index2; i <= len(value); i++ {
			s[num] = value[index2:i]
			if s[num] == "null" || s[num^1] == "null" || s[num] != s[num^1] && solve(s, pattern, index1+1, value, i) {
				return true
			}
		}
		s[num] = "null"
		return false
	} else {
		end := index2 + len(s[num])
		if end > len(value) || value[index2:end] != s[num] {
			return false
		}
		return solve(s, pattern, index1+1, value, end)
	}
}

// 官方解提
func patternMatching2(pattern string, value string) bool {
	countA, countB := 0, 0
	for i := 0; i < len(pattern); i++ {
		if pattern[i] == 'a' {
			countA++
		} else {
			countB++
		}
	}
	if countA < countB {
		countA, countB = countB, countA
		tmp := ""
		for i := 0; i < len(pattern); i++ {
			if pattern[i] == 'a' {
				tmp += "b"
			} else {
				tmp += "a"
			}
		}
		pattern = tmp
	}
	if len(value) == 0 {
		return countB == 0
	}
	if len(pattern) == 0 {
		return false
	}

	for lenA := 0; countA*lenA <= len(value); lenA++ {
		rest := len(value) - countA*lenA
		if (countB == 0 && rest == 0) || (countB != 0 && rest%countB == 0) {
			var lenB int
			if countB == 0 {
				lenB = 0
			} else {
				lenB = rest / countB
			}
			pos, correct := 0, true
			var valueA, valueB string
			for i := 0; i < len(pattern); i++ {
				if pattern[i] == 'a' {
					sub := value[pos : pos+lenA]
					if len(valueA) == 0 {
						valueA = sub
					} else if valueA != sub {
						correct = false
						break
					}
					pos += lenA
				} else {
					sub := value[pos : pos+lenB]
					if len(valueB) == 0 {
						valueB = sub
					} else if valueB != sub {
						correct = false
						break
					}
					pos += lenB
				}
			}
			if correct && valueA != valueB {
				return true
			}
		}
	}
	return false
}

func Test_patternMatching(t *testing.T) {
	pattern := "abba"
	value := "dogcatcatdog"
	fmt.Println(patternMatching(pattern, value))
}
