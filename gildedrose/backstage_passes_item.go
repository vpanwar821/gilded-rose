package gildedrose

import (
	"github.com/vpanwar821/gilded-rose/common/consts"
	"github.com/vpanwar821/gilded-rose/common/utils"
)

type BackstagePassesItem struct {
	*Item
}

func NewBackstagePassesItem(item *Item) *BackstagePassesItem {
	return &BackstagePassesItem{
		Item: item,
	}
}

func (item *BackstagePassesItem) UpdateQuality() {
	item.SellIn--
	if item.SellIn < 0 {
		item.Quality = consts.MinQuality
	} else if item.SellIn <= 5 {
		item.Quality = utils.ValidateQuality(item.Quality + 3)
	} else if item.SellIn <= 10 {
		item.Quality = utils.ValidateQuality(item.Quality + 2)
	} else {
		item.Quality = utils.ValidateQuality(item.Quality + 1)
	}
}

