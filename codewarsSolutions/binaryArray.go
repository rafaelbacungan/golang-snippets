package codewarsSolutions

func Binarray(a []int) int {
	prefixSum := make(map[int]int)
	maxLength := 0
	sum := 0

	prefixSum[0] = -1

	for i := 0; i < len(a); i++ {
		if a[i] == 1 {
			sum++
		} else {
			sum--
		}

		// check if the index is found in the map.
		if index, found := prefixSum[sum]; found {
			currentLength := i - index
			if currentLength > maxLength {
				maxLength = currentLength
			}
		} else {
			prefixSum[sum] = i
		}
	}

	return maxLength
}
