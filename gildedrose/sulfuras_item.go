package gildedrose

import (
	"github.com/vpanwar821/gilded-rose/common/consts"
)

type SulfurasItem struct {
	*Item
}

func NewSulfurasItem(item *Item) *SulfurasItem {
	return &SulfurasItem{
		Item: item,
	}
}

func (item *SulfurasItem) UpdateQuality() {
	item.Quality = consts.FixedQuality
}
