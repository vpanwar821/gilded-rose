package gildedrose

import (
	"testing"

	"github.com/vpanwar821/gilded-rose/common/consts"
)

// "Aged Brie" actually increases in Quality the older it gets
func Test_AgedBrie(t *testing.T) {
	sellIn := 25
	quality := 6

	items := []*Item{
		{"Aged Brie", sellIn, quality},
	}

	for i := 0; i < 40; i++ {
		UpdateQuality(items)
		quality++
		if items[0].Quality != quality {
			t.Errorf("Expected %v as quality, got %v and sellIn=%v", quality, items[0].Quality, items[0].SellIn)
		}
	}
}

// "Backstage passes", like aged brie, increases in Quality as its SellIn value approaches;
// Quality increases by 2 when there are 10 days or less and by 3 when there are 5 days or less but
// Quality drops to 0 after the concert
func Test_BackstagePasses(t *testing.T) {
	sellIn := 25
	quality := 8

	items := []*Item{
		{"Backstage passes to a TAFKAL80ETC concert", sellIn, quality},
	}

	for i := 0; i < 40; i++ {
		UpdateQuality(items)

		if items[0].SellIn < 0 {
			quality = consts.MinQuality
		} else if items[0].SellIn <= 5 {
			quality = quality + 3
		} else if items[0].SellIn <= 10 {
			quality = quality + 2
		} else {
			quality++
		}

		if quality > consts.MaxQuality {
			quality = consts.MaxQuality
		}

		if items[0].Quality != quality {
			t.Errorf("Expected %v as quality, got %v and sellIn=%v", quality, items[0].Quality, items[0].SellIn)
		}
	}
}

// "Sulfuras", being a legendary item, never has to be sold or decreases in Quality
func Test_Sulfuras(t *testing.T) {
	sellIn := 25
	quality := 80

	items := []*Item{
		{"Sulfuras, Hand of Ragnaros", sellIn, quality},
	}
	UpdateQuality(items)
	if items[0].Quality != consts.FixedQuality {
		t.Errorf("Expected quality %v, got %v", items[0].Quality, consts.FixedQuality)
	}
	if items[0].SellIn != sellIn {
		t.Errorf("Expected sellIn %v, got %v", items[0].SellIn, sellIn)
	}
}

// "Conjured" items degrade in Quality twice as fast as normal items
func Test_Conjured(t *testing.T) {
	sellIn := 25
	quality := 15

	items := []*Item{
		{"Conjured Mana Cake", sellIn, quality},
	}

	for i := 0; i < 40; i++ {
		UpdateQuality(items)

		if items[0].SellIn < 0 {
			quality = quality - 4
		} else {
			quality = quality - 2
		}

		if quality < consts.MinQuality {
			quality = consts.MinQuality
		}

		if items[0].Quality != quality {
			t.Errorf("Expected %v as quality, got %v and sellIn=%v", quality, items[0].Quality, items[0].SellIn)
		}
	}
}

// Regular items degrade in quality as the day passes
// Once the sell by date has passed, Quality degrades twice as fast
// The Quality of an item is never negative
func Test_RegularItems(t *testing.T) {
	sellIn := 25
	quality := 15

	items := []*Item{
		{"item1", sellIn, quality},
	}

	for i := 0; i < 40; i++ {
		UpdateQuality(items)

		if items[0].SellIn < 0 {
			quality = quality - 2
		} else {
			quality = quality - 1
		}

		if quality < consts.MinQuality {
			quality = consts.MinQuality
		}

		if items[0].Quality != quality {
			t.Errorf("Expected %v as quality, got %v and sellIn=%v", quality, items[0].Quality, items[0].SellIn)
		}
	}
}