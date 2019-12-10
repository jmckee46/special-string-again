package main

// substrCount looks for a patterns of style aba, aabaa... or aaaa...
// if it finds aa...abaa...a it will calculate the aa..a part first, then the aba/aabaa/aaabaaa combinations
// and then start again on the middle/different letter to look again.
// example aaaabaaabaa
func substrCount(n int32, s string) int64 {
	var combinations int64
	var endOfFirst int32
	var firstLength int32
	var endOfLast int32
	var lastLength int32
	var start int32
	var continueFlag = true

	// add the individual letter combinations
	combinations = int64(n)

	for continueFlag {
		endOfFirst = checkForRepeater(start, s, n)
		firstLength = endOfFirst - start + 1
		// add the repeat combinations, which is aaaa on the first pass from the example above
		combinations += repeatCombinations(firstLength)

		// if repeated letter ends the string then break
		if endOfFirst == n-1 {
			break
		}

		// if a pattern of aba is detected
		if endOfFirst+2 < n && s[start] == s[endOfFirst+2] {
			endOfLast = checkForRepeater(endOfFirst+2, s, n)
			lastLength = endOfLast - endOfFirst - 1

			if firstLength < lastLength {
				combinations += int64(firstLength)
			} else {
				combinations += int64(lastLength)
			}
		}

		start = endOfFirst + 1
		if start >= n-1 {
			continueFlag = false
		}
	}

	return combinations
}

func checkForRepeater(start int32, s string, sLength int32) int32 {

	for i := int32(start); i < sLength; i++ {
		if s[start] != s[i] {
			return i - 1
		}
	}

	return sLength - 1
}

func repeatCombinations(repeatLength int32) int64 {
	var sum int64

	for i := int64(1); i <= int64(repeatLength-1); i++ {
		sum += i
	}

	return sum
}

func main() {}
