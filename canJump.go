package lesson7

func canJump(nums []int) bool {

	length := len(nums)
	inf := 0
	for i := 0; i <= length-1 && i <= inf; i++ {

		inf = max(inf, i+nums[i])
	}

	return inf >= length-1
}

func max(x, y int) int {
	if x >= y {

		return x
	}
	return y
}
