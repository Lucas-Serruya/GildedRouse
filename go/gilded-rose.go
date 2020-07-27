package main

type Item struct {
	name            string
	sellIn, quality int
}

const Backstage = "Backstage passes to a TAFKAL80ETC concert"
const AgedBrie = "Aged Brie"
const Sulfuras = "Sulfuras, Hand of Ragnaros"

// objetivos:
// simplificar las condiciones
// extraer condiciones repetidas
// separar logica en 3 if por tipo de item
// hacer un strategy con cada tipo

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {
		itemzz := ItemFactory(items[i])
		itemzz.UpdateSellIn()
		//items[i].updateQuality()
	}
}

type Itemzz interface {
	UpdateSellIn()
}

type AgedBrieItem struct {
	Item *Item
}

func (a AgedBrieItem) UpdateSellIn() {
	a.Item.updateQuality()
}

type SulfurasItem struct {
	Item *Item
}

func (s SulfurasItem) UpdateSellIn() {
	panic("implement me")
}

type BackstageItem struct {
	Item *Item
}

func (b BackstageItem) UpdateSellIn() {
	panic("implement me")
}

type DefaultItem struct {
	Item *Item
}

func (d DefaultItem) UpdateSellIn() {
	d.Item.updateQuality()
}

var _ Itemzz = (*AgedBrieItem)(nil)
var _ Itemzz = (*SulfurasItem)(nil)
var _ Itemzz = (*BackstageItem)(nil)
var _ Itemzz = (*DefaultItem)(nil)

func ItemFactory(item *Item) Itemzz {
	switch item.name {
	case AgedBrie:
		return &AgedBrieItem{Item: item}
	/*case Sulfuras:
		return &SulfurasItem{Item: item}
	case Backstage:
		return &BackstageItem{Item: item}*/
	default:
		return &DefaultItem{Item: item}
	}
}

func (item *Item) updateQuality() {
	item.decreaseSellIn()

	if item.agesOverTime() {
		item.increaseQuality()
	} else {
		item.decreaseQuality()
	}

	if item.hasExpired() {
		if item.name == Backstage {
			item.resetQuality()
		} else if item.name != AgedBrie {
			item.decreaseQuality()
		}
	}
}

func (item *Item) agesOverTime() bool {
	return item.name == AgedBrie || item.name == Backstage
}

func (item *Item) canDecreaseQuality() bool {
	return item.quality > 0 && item.name != Sulfuras
}

func (item *Item) canIncreaseQuality() bool {
	return item.quality < 50 && item.name != Sulfuras
}

func (item *Item) hasExpired() bool {
	return item.sellIn < 0
}

func (item *Item) decreaseSellIn() {
	if item.name == Sulfuras {
		return
	}
	item.sellIn--
}

func (item *Item) resetQuality() {
	item.quality = 0
}

func (item *Item) increaseQuality() {
	if !item.canIncreaseQuality() {
		return
	}
	if item.name == Backstage {
		if item.sellIn < 11 {
			item.quality++
		}
		if item.sellIn < 6 {
			item.quality++
		}
	}
	item.quality++
}

func (item *Item) decreaseQuality() {
	if !item.canDecreaseQuality() {
		return
	}
	item.quality--
}
