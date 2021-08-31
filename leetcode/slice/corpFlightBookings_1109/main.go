package main

func main() {

}

func corpFlightBookings(bookings [][]int, n int) []int {
	ans := make([]int, n)
	for _, book := range bookings {
		for start := book[0]; start <= book[1]; start++ {
			ans[start-1] += book[2]
		}
	}
	return ans
}