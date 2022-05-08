package utils

import (
	"github.com/vpanwar821/gilded-rose/common/consts"
)

// ValidateQuality validate the quality and return the valid quality
func ValidateQuality(quality int) int {
	if quality > consts.MaxQuality {
		return consts.MaxQuality
	} else if quality < consts.MinQuality {
		return consts.MinQuality
	} else {
		return quality
	}
}