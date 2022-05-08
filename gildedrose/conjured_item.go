package gildedrose

import "github.com/vpanwar821/gilded-rose/common/utils"

type ConjuredItem struct {
	*Item
}

func NewConjuredItem(item *Item) *ConjuredItem {
	return &ConjuredItem{
		Item: item,
	}
}

func (item *ConjuredItem) UpdateQuality() {
	item.SellIn--
	if item.SellIn < 0 {
		item.Quality = utils.ValidateQuality(item.Quality - 4)
	} else {
		item.Quality = utils.ValidateQuality(item.Quality - 2)
	}
}

