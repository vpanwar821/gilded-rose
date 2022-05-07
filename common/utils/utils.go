package utils

// validateQuality validate the quality and return the valid quality
func validateQuality(quality int) int {
	if quality > maxQuality {
		return maxQuality
	} else if quality < minQuality {
		return minQuality
	} else {
		return quality
	}
}