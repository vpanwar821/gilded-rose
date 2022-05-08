package gildedrose

import (
	"github.com/vpanwar821/gilded-rose/common/consts"
)

type Item struct {
	Name            string
	SellIn, Quality int
}

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {
		switch items[i].Name {
		case consts.AgedBrie:
			NewAgedBrieItem(items[i]).UpdateQuality()
		case consts.BackstagePasses:
			NewBackstagePassesItem(items[i]).UpdateQuality()
		case consts.Sulfuras:
			NewSulfurasItem(items[i]).UpdateQuality()
		case consts.Conjured:
			NewConjuredItem(items[i]).UpdateQuality()
		default:
			NewRegularItem(items[i]).UpdateQuality()
		}
	}
}
