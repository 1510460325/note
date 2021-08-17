package main

func main() {

}

func checkRecord(s string) bool {
	absentTime, sequencesLateTime := 0, 0
	for i := range s {
		if s[i] == 'A' {
			absentTime++
			if absentTime == 2 {
				return false
			}
		} else if s[i] == 'L' {
			sequencesLateTime++
			if sequencesLateTime == 3 {
				return false
			}
		} else {
			sequencesLateTime = 0
		}
	}
	return true
}
