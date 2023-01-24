package main

func polindrome(s string) bool {
	count := -1
	st := []int32{}
	if len(s) <= 1 {
		return false
	}
	for _, word := range s {
		st = append(st, word)
		count++
	}
	for i := 0; i < len(st); i++ {
		if st[i] == st[count] {
			count -= 1
		} else {
			return false
		}
	}
	if count <= len(st)/2 {
		return true
	}
	return false
}
func main() {

}
