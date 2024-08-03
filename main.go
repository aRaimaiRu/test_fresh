package main

import (
	"fmt"
	"os"
	"strings"
)

type Calculator struct {
	PriceMap       map[string]int
	totalPrice     int
	isMember       bool
	memberDiscount float32
	discountMap    map[string]Discount
}

func (c *Calculator) SetMember(isMember bool) {
	c.isMember = isMember
}

func (c *Calculator) CalculatePrice(order map[string]int) int {
	for item, quantity := range order {
		itemPrice := c.PriceMap[item] * int(quantity)
		if discount, ok := c.discountMap[item]; ok {
			// discount for each quantity = discount.Count

			bundlePrice := c.PriceMap[item] * (quantity / discount.Count) * discount.Count
			itemPrice -= int(float32(bundlePrice) * float32(discount.Percent))
			// itemPrice -= int(float32(itemPrice) * float32(quantity/discount.Count) * discount.Percent)
		}
		c.totalPrice += itemPrice
	}
	// discount if member
	if c.isMember {
		c.totalPrice = int(float32(c.totalPrice) * c.memberDiscount)
	}
	return c.totalPrice
}

// 1.00 = 100
var colorPriceMap = map[string]int{
	"RED":    5000,
	"GREEN":  4000,
	"BLUE":   3000,
	"YELLOW": 5000,
	"PINK":   8000,
	"PURPLE": 9000,
	"ORANGE": 12000,
}

type Discount struct {
	Product string
	Count   int
	Percent float32
}

// map product discount
var discountMap = map[string]Discount{
	"GREEN":  {"GREEN", 2, 0.05},
	"PINK":   {"PINK", 2, 0.05},
	"ORANGE": {"ORANGE", 2, 0.05},
}

func main() {
	order := readArgs()
	Calculator := &Calculator{
		PriceMap:       colorPriceMap,
		isMember:       true,
		totalPrice:     0,
		memberDiscount: 0.9,
		discountMap:    discountMap,
	}
	result := Calculator.CalculatePrice(order)
	fmt.Println(result)
}

func readArgs() map[string]int {
	productCounts := make(map[string]int)
	for _, arg := range os.Args[1:] {
		// trim "--"
		arg = strings.TrimPrefix(arg, "--")
		// Split the argument by '=' to get the product and count
		parts := strings.Split(arg, "=")
		if len(parts) != 2 {
			fmt.Printf("Invalid format for argument: %s\n", arg)
			continue
		}
		product := parts[0]
		var count int
		_, err := fmt.Sscanf(parts[1], "%d", &count)
		if err != nil {
			fmt.Printf("Invalid count for product %s: %s\n", product, parts[1])
			continue
		}
		productCounts[product] = count
	}

	return productCounts
}
