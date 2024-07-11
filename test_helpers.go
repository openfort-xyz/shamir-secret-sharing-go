package sss

func generateCombinations(arr [][]byte, k int) [][][]byte {
	var result [][][]byte
	var helper func([][]byte, int, int)
	helper = func(currentCombo [][]byte, start int, k int) {
		if k == 0 {
			comboCopy := make([][]byte, len(currentCombo))
			copy(comboCopy, currentCombo)
			result = append(result, comboCopy)
			return
		}
		for i := start; i <= len(arr)-k; i++ {
			helper(append(currentCombo, arr[i]), i+1, k-1)
		}
	}
	helper([][]byte{}, 0, k)
	return result
}
