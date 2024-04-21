package common

// StringInSlice returns true if passed slice contains target.
func StringInSlice(target string, slice []string) bool {
	for _, elem := range slice {
		if elem == target {
			return true
		}
	}
	return false
}

func BytesToMegabytes(bytes uint64) float64 {
	return float64(bytes) / 1024 / 1024
}
