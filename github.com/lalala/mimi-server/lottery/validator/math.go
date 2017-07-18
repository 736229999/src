package validator

import "math"

func Factorial(n int) int {
	if n < 0 {
		panic("n < 0")
	}
	if n <= 1 {
		return 1
	}
	ret := 1
	for i := 2; i <= n; i++ {
		ret *= i
	}
	return ret
}

// 求n个数中选m个数的组合数(m <= n, 且 m >= 0)
func Comb(n, m int) int {
	if n < m {
		return 0
	}
	return Factorial(n) / (Factorial(n-m) * Factorial(m))
}

func IsEqualMoney(f1, f2 float64) bool {
	return math.Abs(f1-f2) < 0.01 // 1分钱精度
}

//求n个数中m个数的排列数
// P(n, m) = n!/(n-m)!
func Perm(n, m int) int {
	if n < m {
		return 0
	}
	return Factorial(n) / Factorial(m)
}

// 求交集数量
func CalInter2Num(list1, list2 []int32) int {
	interNum := 0 // 交集数量
	for _, v1 := range list1 {
		for _, v2 := range list2 {
			if v1 == v2 {
				interNum++
				break
			}
		}
	}
	return interNum
}
