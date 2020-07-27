package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_QualityIsNeverNegative(t *testing.T) {
	var items = []*Item{
		{"foo", 5, 0},
	}

	UpdateQuality(items)
	UpdateQuality(items)
	UpdateQuality(items)

	assert.Equal(t, items[0].quality, 0)

}

func Test_QualityNeverMoreThan50(t *testing.T) {
	var items = []*Item{
		{AgedBrie, 5, 5},
	}

	for i := 0; i < 50; i++ {
		UpdateQuality(items)
	}

	assert.Equal(t, items[0].quality, 50)
}

func Test_QualityOfSulfurasIsNeverAlteredQuality(t *testing.T) {
	var items = []*Item{
		{Sulfuras, 5, 80},
	}

	UpdateQuality(items)
	UpdateQuality(items)
	UpdateQuality(items)

	assert.Equal(t, items[0].quality, 80)
}

func Test_SulfurasIsNeverAlteredSellIn(t *testing.T) {
	var items = []*Item{
		{Sulfuras, 5, 80},
	}

	UpdateQuality(items)
	UpdateQuality(items)
	UpdateQuality(items)

	assert.Equal(t, items[0].sellIn, 5)
}

func Test_QualityDecreaseTwiceAfterExpiration(t *testing.T) {
	var items = []*Item{
		{"name", 2, 50},
	}

	UpdateQuality(items)
	assert.Equal(t, items[0].quality, 49)
	UpdateQuality(items)
	assert.Equal(t, items[0].quality, 48)
	UpdateQuality(items)
	assert.Equal(t, items[0].quality, 46)

}

func Test_BackstagePassesQuality(t *testing.T) {

	t.Run("More than 10 days", func(t *testing.T) {
		var items = []*Item{
			{Backstage, 15, 1},
		}
		UpdateQuality(items)
		assert.Equal(t, items[0].quality, 2)
	})

	t.Run("Between 10 and 5", func(t *testing.T) {
		var items = []*Item{
			{Backstage, 10, 1},
		}
		UpdateQuality(items)
		assert.Equal(t, items[0].quality, 3)
	})

	t.Run("Between 5 and 1", func(t *testing.T) {
		var items = []*Item{
			{Backstage, 5, 1},
		}
		UpdateQuality(items)
		assert.Equal(t, items[0].quality, 4)
	})

	t.Run("Expired", func(t *testing.T) {
		var items = []*Item{
			{Backstage, 0, 1},
		}
		UpdateQuality(items)
		assert.Equal(t, items[0].quality, 0)
	})

}

func Test_AgedBrieIncreaseQualityOneTime(t *testing.T) {
	t.Run("Expired", func(t *testing.T) {
		var items = []*Item{
			{AgedBrie, 0, 1},
		}
		UpdateQuality(items)
		assert.Equal(t, 2, items[0].quality)
	})
}
