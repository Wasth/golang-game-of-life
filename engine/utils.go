package engine

func IsValidState(state uint8) bool {
	for _, v := range STATES {
		if v == state {
			return true
		}
	}
	return false
}
