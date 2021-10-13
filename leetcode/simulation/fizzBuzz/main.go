package main

import "fmt"

func main() {

}

func fizzBuzz(n int) (ans []string) {
	for i := 1; i <= n; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			ans = append(ans, "FizzBuzz")
		} else if i % 3 == 0 {
			ans = append(ans, "Fizz")
		} else if i % 5 == 0 {
			ans = append(ans, "Buzz")
		} else {
			ans = append(ans, fmt.Sprintf("%v", i))
		}
	}
	return ans
}
