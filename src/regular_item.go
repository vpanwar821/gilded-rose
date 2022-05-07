package gildedrose

import (
	"github.com/vpanwar821/gilded-rose/common/utils"
)

type regularItem struct {
	Item
}

func (item *regularItem) updateQuality() {
	item.SellIn--
	if item.SellIn < 0 {
		item.Quality = utils.validateQuality(item.Quality - 2)
	} else {
		item.Quality = validateQuality(item.Quality - 1)
	}
}

