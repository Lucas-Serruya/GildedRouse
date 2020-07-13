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

		item := items[i]

		if item.name != AgedBrie && item.name != Backstage {
			if item.quality > 0 {
				if item.name != Sulfuras {
					item.decreaseQuality()
				}
			}
		} else {
			if item.quality < 50 {
				item.increaseQuality()
				if item.name == Backstage {
					if item.sellIn < 11 {
						item.increaseQuality()
					}
					if item.sellIn < 6 {
						item.increaseQuality()
					}
				}
			}
		}

		if item.name != Sulfuras {
			item.decreaseSellIn()
		}

		if item.sellIn < 0 {
			if item.name != AgedBrie {
				if item.name != Backstage {
					if item.quality > 0 && item.name != Sulfuras {
						item.decreaseQuality()
					}
				} else {
					item.expire()
				}
			} else {
				if item.quality < 50 {
					item.increaseQuality()
				}
			}
		}
	}

}

func (item *Item) decreaseSellIn() {
	item.sellIn = item.sellIn - 1
}

func (item *Item) expire() {
	item.quality = 0
}

func (item *Item) increaseQuality() {
	item.quality = item.quality + 1
}

func (item *Item) decreaseQuality() {
	item.quality = item.quality - 1
}
