package gildedrose

import "github.com/vpanwar821/gilded-rose/common/utils"

type AgedBrieItem struct {
	*Item
}

func NewAgedBrieItem(item *Item) *AgedBrieItem {
	return &AgedBrieItem{
		Item: item,
	}
}

func (item *AgedBrieItem) UpdateQuality() {
	item.SellIn--
	item.Quality = utils.ValidateQuality(item.Quality + 1)
}
