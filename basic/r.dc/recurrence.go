package r_dc

// Recurrence 阶乘递归
func Recurrence(n int) int {
	if n == 1 {
		return n
	}

	return n * Recurrence(n-1)
}

// RecurrenceTail 阶乘尾递归
func RecurrenceTail(n, a int) int {
	if n == 1 {
		return a
	}

	return RecurrenceTail(n-1, a*n)
}

// Fib 斐波那契数列
// F(n) = F(n-1) + F(n-2)
// 0 1
// 1 1
// 2 1 + 1 = 2
// 3 2 + 1 = 3
// 4 3 + 2 = 5
// 5 5 + 3 = 8
func Fib(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	return Fib(n-1) + Fib(n-2)
}

// FibTail 斐波那契额数列尾递归
// F(5,1,1)
// F(4,1,1+1) = F(4,1,2)
// F(3,2,1+2) = F(3,2,3)
// F(2,3,2+3) = F(2,3,5)
// F(1,5,3+5) = F(1,5,8)
// F(0,8,5+8) = F(0,8,13)
func FibTail(n, a, b int) int {
	if n == 0 {
		return a
	}

	return FibTail(n-1, b, a+b)
}
