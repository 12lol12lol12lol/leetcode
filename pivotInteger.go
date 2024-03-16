package main

func pivotInteger(n int) int {
	leftPointer, rightPointer := 1, n
	sumLeft, sumRight := leftPointer, rightPointer
	for leftPointer < rightPointer {
		if sumLeft < sumRight {
			leftPointer++
			sumLeft += leftPointer
		} else {
			rightPointer--
			sumRight += rightPointer
		}
	}
	if sumLeft == sumRight {
		return leftPointer
	}
	return -1
}
