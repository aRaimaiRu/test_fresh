package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getCalculator() Calculator {
	return Calculator{
		PriceMap:       colorPriceMap,
		totalPrice:     0,
		isMember:       false,
		memberDiscount: 0.9,
		discountMap:    discountMap,
	}

}
func TestCalculatePrceWithoutMember(t *testing.T) {
	calculator := getCalculator()

	order := map[string]int{
		"RED":    4,
		"GREEN":  4,
		"BLUE":   4,
		"YELLOW": 4,
		"PINK":   4,
		"PURPLE": 4,
		"ORANGE": 4,
	}
	var expectedPrice int = 4*colorPriceMap["RED"] + (4 * int(float32(colorPriceMap["GREEN"])*float32(0.95))) + 4*colorPriceMap["BLUE"] + 4*colorPriceMap["YELLOW"] + (4 * int(float32(colorPriceMap["PINK"])*float32(0.95))) + 4*colorPriceMap["PURPLE"] + (4 * int(float32(colorPriceMap["ORANGE"])*float32(0.95)))
	actualPrice := calculator.CalculatePrice(order)
	assert.Equal(t, expectedPrice, actualPrice)
}

func TestCalculatePrceWithMember(t *testing.T) {
	calculator := getCalculator()
	calculator.SetMember(true)

	order := map[string]int{
		"RED":    4,
		"GREEN":  4,
		"BLUE":   4,
		"YELLOW": 4,
		"PINK":   4,
		"PURPLE": 4,
		"ORANGE": 4,
	}
	var expectedPrice int = (4*colorPriceMap["RED"] + (4 * int(float32(colorPriceMap["GREEN"])*float32(0.95))) + 4*colorPriceMap["BLUE"] + 4*colorPriceMap["YELLOW"] + (4 * int(float32(colorPriceMap["PINK"])*float32(0.95))) + 4*colorPriceMap["PURPLE"] + (4 * int(float32(colorPriceMap["ORANGE"])*float32(0.95))))
	expectedPrice = int(float32(expectedPrice) * 0.9)
	actualPrice := calculator.CalculatePrice(order)

	assert.Equal(t, expectedPrice, actualPrice)
}

func TestCalculateDiscountBundleWithExcessiveCount(t *testing.T) {
	calculator := getCalculator()
	calculator.SetMember(true)

	order := map[string]int{
		"ORANGE": 3,
	}

	var expectedPrice int = int(float32(colorPriceMap["ORANGE"]) * 2 * 0.95 * 0.9)
	expectedPrice += int(float32(colorPriceMap["ORANGE"]) * 1 * 0.9)

	actualPrice := calculator.CalculatePrice(order)

	assert.Equal(t, expectedPrice, actualPrice)
}

func TestCalculateDiscountBundle(t *testing.T) {
	calculator := getCalculator()
	calculator.SetMember(true)

	order := map[string]int{
		"ORANGE": 2,
	}

	var expectedPrice int = int(float32(colorPriceMap["ORANGE"]) * 2 * 0.95 * 0.9)

	actualPrice := calculator.CalculatePrice(order)

	assert.Equal(t, expectedPrice, actualPrice)
}

func TestCalculatePriceWithExcessBundle(t *testing.T) {
	calculator := getCalculator()
	calculator.SetMember(true)

	order := map[string]int{
		"RED":    4,
		"GREEN":  5,
		"BLUE":   4,
		"YELLOW": 4,
		"PINK":   5,
		"PURPLE": 4,
		"ORANGE": 5,
	}
	var expectedPrice int = (4*colorPriceMap["RED"] + (4 * int(float32(colorPriceMap["GREEN"])*float32(0.95))) + 4*colorPriceMap["BLUE"] + 4*colorPriceMap["YELLOW"] + (4 * int(float32(colorPriceMap["PINK"])*float32(0.95))) + 4*colorPriceMap["PURPLE"] + (4 * int(float32(colorPriceMap["ORANGE"])*float32(0.95))))
	// sum order count that excess from bundle
	expectedPrice += colorPriceMap["GREEN"] + colorPriceMap["PINK"] + colorPriceMap["ORANGE"]
	expectedPrice = int(float32(expectedPrice) * 0.9)
	actualPrice := calculator.CalculatePrice(order)

	assert.Equal(t, expectedPrice, actualPrice)
}
