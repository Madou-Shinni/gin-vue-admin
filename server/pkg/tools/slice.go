package tools

func ContainsUint(slice []uint, item uint) bool {
	for _, i := range slice {
		if i == item {
			return true
		}
	}
	return false
}
