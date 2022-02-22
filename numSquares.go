package lesson7

//numSquares: 给你一个整数 n ，返回和为n的完全平方数的最少数量
func numSquares(n int) int {

	opt := make([]int, n+1)
	for i := 1; i <= n; i++ {

		opt[i] = i
		for j := 1; j*j <= i; j++ {

			opt[i] = min(opt[i], opt[i-j*j]+1)
		}
	}

	return opt[n]
}

func min(x, y int) int {

	if x <= y {

		return x
	}

	return y
}
