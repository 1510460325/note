package main

func main() {

}

func fib(n int) int {
	if n < 2 {
		return n
	}
	p1, p2 := 0, 1
	for i := 2; i <= n; i++ {
		p1, p2 = p2, p1 + p2
	}
	return p2
}
