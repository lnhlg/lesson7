package lesson7

//canJump: 跳跃游戏
//给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。数组中的每个元素代表你在该位置可以跳跃的最大长度。判断你是否能够到达最后一个下标。
/*parameter
nums: 非负整数数组
return: 是否能够到达最后一个下标
*/
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
