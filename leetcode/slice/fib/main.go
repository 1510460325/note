package fib


func fib(n int) int {
	if n < 2 {
		return n
	}
	const mod int = 1e9+7
	p1, p2 := 0, 1
	for i := 2; i <= n; i++ {
		p1, p2 = p2, (p1 + p2) % mod
	}
	return p2
}
