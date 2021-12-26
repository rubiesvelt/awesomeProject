package main

/**
279. 完全平方数
*/
func numSquares(n int) int {
	f := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		f[i] = 0x3f3f3f3f
	}
	for i := 1; i*i <= n; i++ {
		u := i * i
		for j := u; j <= n; j++ {
			f[j] = min(f[j], f[j-u]+1)
		}
	}
	return f[n]
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

/**
322. 零钱兑换
amount 最少 由coins中几个硬币构成
典型的完全背包问题
*/
func coinChange(coins []int, amount int) int {
	f := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		f[i] = 0x3f3f3f3f
	}
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			f[i] = min(f[i], f[i-coin]+1)
		}
	}
	if f[amount] == 0x3f3f3f3f {
		return -1
	}
	return f[amount]
}

/**
518. 零钱兑换 II
amount 由 coins 中硬币构成的 不同方案数

5=5
5=2+2+1
5=2+1+1+1
5=1+1+1+1+1
*/
func change(amount int, coins []int) int {
	f := make([]int, amount+1)
	f[0] = 1
	for _, coin := range coins { // 必须 _, coin 不能直接 coin
		for i := coin; i <= amount; i++ {
			// 3 = 1 + 2
			f[i] += f[i-coin]
		}
	}
	return f[amount]
}

/*
 * 343. 整数拆分
 */
func integerBreak(n int) int {
	f := make([]int, n+1) // 创建切片
	for i := 1; i <= n; i++ {
		for j := 1; j < i; j++ {
			f[i] = max(f[i], max((i-j)*f[j], (i-j)*j))
		}
	}
	return f[n]
}
