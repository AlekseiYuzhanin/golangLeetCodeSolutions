func lengthOfLongestSubstring(s string) int {
	lastOcc := make(map[byte]int)
	srart := 0
	maxLength := 0
	for i, ch := range []byte(s) {
		if lastI, ok := lastOcc[ch]; ok && lastI >= srart {
			srart = lastI + 1
		}
		if i-srart+1 > maxLength {
			maxLength = i - srart + 1
		}
		lastOcc[ch] = i
	}
	return maxLength
}