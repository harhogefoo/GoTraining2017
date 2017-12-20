package ex12


func isAnagram(str1, str2 string) bool {
	dict1 := map[rune]int{}
	dict2 := map[rune]int{}

	for _, r := range str1 {
		dict1[r] += 1
	}

	for _, r := range str2 {
		dict2[r] += 1
	}

	for r, count := range dict1 {
		if dict2[r] != count {
			return false
		}
	}

	for r, count := range dict2 {
		if dict1[r] != count {
			return false
		}
	}

	return true
}