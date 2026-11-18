package pile

func EstTriee(piles Piles) bool {
	if len(piles.B) != 0 {
		return false
	}
	for i := 0; i < len(piles.A)-1; i++ {
		if piles.A[i] > piles.A[i+1] {
			return false
		}
	}
	return true
}
