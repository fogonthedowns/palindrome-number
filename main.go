package main

func isPalindrome(x int) bool {
	// Special cases:
	// 1. numbers that end with 0
	//   If the last digit is zero, the first would need to be zero
	//   so 10 can't be a palendrome. e.g. 10 palendrome is 010
	// 2. Negative numbers
	//   e.g. -121 planedrome is 121-
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	// algorithm, "split" the number and compare
	// do so by geting the last character
	//   e.g. x % 10
	// and appending to revertedNumber
	//   revertedNumber * 10 + revertedNumber
	var revertedNumber int
	for x > revertedNumber {
		// algorithm
		revertedNumber = revertedNumber*10 + x%10
		// chops the last number off the end
		x /= 10
	}

	// even and odd cases considered with this or
	// eliminates middle number
	return x == revertedNumber || x == revertedNumber/10
}
