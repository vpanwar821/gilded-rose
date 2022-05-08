package gildedrose

import "github.com/vpanwar821/gilded-rose/common/utils"

type RegularItem struct {
	*Item
}

func NewRegularItem(item *Item) *RegularItem {
	return &RegularItem{
		Item: item,
	}
}

func (item *RegularItem) UpdateQuality() {
	item.SellIn--
	if item.SellIn < 0 {
		item.Quality = utils.ValidateQuality(item.Quality - 2)
	} else {
		item.Quality = utils.ValidateQuality(item.Quality - 1)
	}
}

