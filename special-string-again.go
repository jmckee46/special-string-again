package main

import "fmt"

func substrCount(n int32, s string) int64 {
	var combinations int64
	var endOfFirst int32
	var firstLength int32
	var endOfLast int32
	var lastLength int32
	var start int32
	var continueFlag = true

	// so far this calculates the length of one pattern, not the combinations yet!
	// need logic to calculate same letter repeated combinations and
	// need logic to calculate pattern combinations
	// need to start at the right place, (current +1) on next iteration of loop

	combinations = int64(n)

	for continueFlag {
		endOfFirst = checkForRepeater(start, s, n)
		firstLength = endOfFirst - start + 1
		combinations += repeatCombinations(firstLength)
		fmt.Println("endOfFirst", endOfFirst, s[start:endOfFirst+1])

		if endOfFirst == n-1 {
			fmt.Println("break due to end of string")
			break
		}

		if s[start] == s[endOfFirst+2] {
			fmt.Println("s[start] == s[endOfFirst+2]", s[start], s[endOfFirst+2])
			endOfLast = checkForRepeater(endOfFirst+2, s, n)
			lastLength = endOfLast - endOfFirst + 3
			if firstLength < lastLength {
				fmt.Println("adding to combinations for patterns:", firstLength)
				combinations += int64(firstLength)
			} else {
				fmt.Println("adding to combinations for patterns:", lastLength)
				combinations += int64(lastLength)
			}
		} else {
			endOfLast = endOfFirst + 1
		}
		fmt.Println("endOfLast", endOfLast, s[start:endOfLast+1])

		start = endOfFirst + 1
		fmt.Println("start = ", endOfFirst+1)
		fmt.Println("")
		if start >= n-2 {
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
		fmt.Println("for loop:", s[i], i)
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
