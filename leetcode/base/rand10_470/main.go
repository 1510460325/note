package main

import "math/rand"

func main() {

}

func rand7() int {
	return rand.Intn(7) + 1
}

func rand10() int {
	for {
		res := (rand7()-1)*7 + rand7()
		if res > 40 {
			continue
		}
		return res % 10 + 1
	}
}
