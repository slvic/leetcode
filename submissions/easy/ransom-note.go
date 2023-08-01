package easy

func canConstruct(ransomNote string, magazine string) bool {
	magazineAlphabet := make(map[rune]int)

	for _, rune := range magazine {
		magazineAlphabet[rune]++
	}

	for _, rune := range ransomNote {
		magazineAlphabet[rune]--
		if magazineAlphabet[rune] < 0 {
			return false
		}
	}

	return true
}
