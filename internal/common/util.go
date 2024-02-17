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
