package main

func isPalindrome(s string) bool {
	if len(s) < 2 {
		return true
	}
	leftPointer, rightPointer := 0, len(s)-1
	for leftPointer < rightPointer {
		rightIsAlphaNumeric, leftIsAlphaNumeric := isAlphaNumeric(s[rightPointer]), isAlphaNumeric(s[leftPointer])
		if rightIsAlphaNumeric && leftIsAlphaNumeric {
			if getLowerCase(s[rightPointer]) != getLowerCase(s[leftPointer]) {
				return false
			}
			rightPointer--
			leftPointer++
			continue
		}
		if !rightIsAlphaNumeric {
			rightPointer--
		}
		if !leftIsAlphaNumeric {
			leftPointer++
		}
	}
	return true
}

func isAlphaNumeric(code uint8) bool {
	return (code >= 48 && code <= 57) || (code >= 65 && code <= 90) || (code >= 97 && code <= 122)
}

func getLowerCase(code uint8) uint8 {
	if code >= 65 && code <= 90 {
		return code + 32
	}
	return code
}
