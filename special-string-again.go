package main

import "fmt"

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

	fmt.Println("adding individual letter count to combinations:", n)
	// add the individual letter combinations
	combinations = int64(n)

	for continueFlag {
		fmt.Printf("starting on: %c\n", s[start])
		endOfFirst = checkForRepeater(start, s, n)
		firstLength = endOfFirst - start + 1
		// add the repeat combinations, which is aaaa on the first pass from the example above
		combinations += repeatCombinations(firstLength)
		fmt.Println("endOfFirst", endOfFirst, s[start:endOfFirst+1], "firstLength", firstLength)

		if endOfFirst == n-1 {
			fmt.Println("break due to end of string")
			break
		}
		// if a pattern of aba is detected
		if endOfFirst+2 < n && s[start] == s[endOfFirst+2] {
			fmt.Printf("s[start] == s[endOfFirst+2], %c = %c\n", s[start], s[endOfFirst+2])
			endOfLast = checkForRepeater(endOfFirst+2, s, n)
			lastLength = endOfLast - endOfFirst - 1
			fmt.Println("endOfLast", endOfLast, s[endOfFirst:endOfLast], "lastLength", lastLength)
			if firstLength < lastLength {
				fmt.Println("adding firstLength combinations for patterns:", firstLength)
				fmt.Println("combinations = ", combinations)
				combinations += int64(firstLength)
			} else {
				fmt.Println("adding lastLength combinations for patterns:", lastLength)
				combinations += int64(lastLength)
			}
		} else {
			endOfLast = endOfFirst + 1
		}
		fmt.Println("endOfLast", endOfLast, s[start:endOfLast+1])

		start = endOfFirst + 1
		fmt.Println("start = ", endOfFirst+1)
		fmt.Println("combinations = ", combinations)
		fmt.Println("")
		if start >= n-1 {
			continueFlag = false
		}
	}

	return combinations
}

// aaaabaaacaaaa
// aaaabaaa
func checkForRepeater(start int32, s string, sLength int32) int32 {
	fmt.Println("checkForRepeater called")
	fmt.Println(start, s, sLength)

	for i := int32(start); i < sLength; i++ {
		fmt.Printf("for loop: %c, %d\n", s[i], i)
		if s[start] != s[i] {
			return i - 1
		}
	}
	fmt.Println("out of for loop")

	return sLength - 1
}

func repeatCombinations(repeatLength int32) int64 {
	var sum int64

	for i := int64(1); i <= int64(repeatLength-1); i++ {
		sum += i
	}

	fmt.Println("repeat combinations:", sum)

	return sum
}

func checkPatternsStartingAt(start int32, s string, sLength int32) int64 {
	var firstLength int32
	// var middle int32
	var lastLength int32
	// var change int32

	for i := int32(start); i < sLength; i++ {
		firstLength = checkForRepeater(start, s, sLength)
		startOfLast := start + firstLength + 2
		if s[startOfLast] == s[start] {
			lastLength = checkForRepeater(startOfLast, s, sLength)
		}
		if lastLength >= firstLength {
			return 1
		}

	}

	return 0
}

func main() {}
