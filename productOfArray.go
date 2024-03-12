package main

func productExceptSelf(nums []int) []int {
	numsLen := len(nums)
	if len(nums) < 2 {
		return []int{}
	}
	answerAr := make([]int, numsLen)
	answerAr[numsLen-1] = nums[numsLen-1]
	for i := numsLen - 2; i >= 0; i-- {
		answerAr[i] = answerAr[i+1] * nums[i]
	}
	R := 1
	for i := 0; i < numsLen-1; i++ {
		answerAr[i] = R * answerAr[i+1]
		R = R * nums[i]
	}
	answerAr[numsLen-1] = R
	return answerAr
}

/*
[1,2,3,4,5,6,7]
[0,1,2,3,4,5,6]
i = 3
[1*2*3*5*6*7]
i=2 1*2*3=6
i=4 5*6*7=210
[1*2*3*4*5*6*7-1*2*3*4]
*/
