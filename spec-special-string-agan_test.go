package main

import "testing"

func TestSpecialStringAgain1(t *testing.T) {
	n := int32(8)
	s := "mnonopoo"

	possibleStrings := substrCount(n, s)

	if possibleStrings != 12 {
		t.Errorf("got %d instead of 12", possibleStrings)
	}
}

func TestSpecialStringAgain2(t *testing.T) {
	n := int32(5)
	s := "asasd"

	possibleStrings := substrCount(n, s)

	if possibleStrings != 7 {
		t.Errorf("got %d instead of 7", possibleStrings)
	}
}

func TestSpecialStringAgain3(t *testing.T) {
	n := int32(7)
	s := "abcbaba"

	possibleStrings := substrCount(n, s)

	if possibleStrings != 10 {

		t.Errorf("got %d instead of 10", possibleStrings)
	}
}

func TestSpecialStringAgain4(t *testing.T) {
	n := int32(4)
	s := "aaaa"

	possibleStrings := substrCount(n, s)

	if possibleStrings != 10 {
		t.Errorf("got %d instead of 10", possibleStrings)
	}
}

func TestSpecialStringAgain5(t *testing.T) {
	n := int32(11)
	s := "aaaabaaacaa"

	possibleStrings := substrCount(n, s)

	if possibleStrings != 26 {
		t.Errorf("got %d instead of 26", possibleStrings)
	}
}
