package main

import (
	"fmt"
	"math"
)

func main() {

	items := []*Item{
		{Id: "1", Quantity: 11, UnitPrice: 1100},
		{Id: "2", Quantity: 5, UnitPrice: 500},
		{Id: "3", Quantity: 4, UnitPrice: 400},
	}
	//items := []*Item{
	//	{Id: "1", Quantity: 3, UnitPrice: 600},
	//	{Id: "2", Quantity: 2, UnitPrice: 200},
	//}

	applyDiscountWithDonation(17, items)
}

func applyDiscountWithDonation(discount uint, items []*Item) {
	var solutions []Discounts
	for _, item := range items {
		for _, discounts := range recursiveDiscount(item, items, Discounts{}, discount) {
			solutions = append(solutions, discounts)
		}
	}

	fmt.Printf("solutions: len()=%d\n", len(solutions))
	for _, solution := range solutions {
		fmt.Printf("%+v\n", solution)
	}

	minDonation := uint(math.MaxUint64)
	var minDonationSolutions []Discounts
	for _, solution := range solutions {
		totalDonation := solution.TotalDonation()
		if totalDonation < minDonation {
			minDonation = totalDonation
			minDonationSolutions = []Discounts{solution}
		} else if totalDonation == minDonation {
			minDonationSolutions = append(minDonationSolutions, solution)
		}
	}

	fmt.Printf("\n\n\n")
	fmt.Printf("minDonationSolutions: len()=%d\n", len(minDonationSolutions))
	for _, solution := range minDonationSolutions {
		fmt.Printf("solution: %v\n", solution.SumPercent())
		for _, v := range solution {
			fmt.Printf("\t%+v\n", v)
		}
	}

	minTotalPercent := math.MaxFloat64
	minTotalPercentIndex := -1
	for i, solution := range minDonationSolutions {
		percent := solution.SumPercent()
		if percent < minTotalPercent {
			minTotalPercent = percent
			minTotalPercentIndex = i
		}
	}

	if minTotalPercentIndex < 0 {
		return
	}

	solution := minDonationSolutions[minTotalPercentIndex]

	fmt.Printf("\n\n\n")
	fmt.Printf("applying solution: %v\n", solution.SumPercent())
	for _, v := range solution {
		fmt.Printf("\t%+v\n", v)
	}

	lookup := make(map[string]*Item, len(items))
	for _, item := range items {
		lookup[item.Id] = item
	}
	for _, discount := range solution {
		item := lookup[discount.ItemId]
		item.UnitDiscount += discount.UnitValue
		item.ItemDiscount += discount.TotalValue
		item.DonatedDiscount += discount.Donation
	}

	fmt.Printf("\n\n\n")
	fmt.Println("result:")
	for i, item := range items {
		fmt.Printf("Item [%d]: %+v\n", i, item)
	}
}

func recursiveDiscount(item *Item, items []*Item, appliedDiscounts Discounts, remaining uint) []Discounts {
	if remaining <= 0 {
		return []Discounts{appliedDiscounts}
	}

	var alreadyApplied uint
	for _, discount := range appliedDiscounts {
		if discount.ItemId == item.Id {
			alreadyApplied += discount.TotalValue
		}
	}

	discountable := item.TotalDiscountedPrice()
	if discountable-alreadyApplied <= 0 {
		return []Discounts{appliedDiscounts}
	}

	onePenny := uint(1) // one cent unit discount
	totalPennies := item.Quantity * onePenny
	var donation uint
	if totalPennies > remaining {
		donation = totalPennies - remaining
		remaining = 0
	} else {
		remaining -= totalPennies
	}
	appliedDiscounts = append(appliedDiscounts, Discount{
		ItemId:     item.Id,
		UnitValue:  onePenny,
		TotalValue: totalPennies,
		Donation:   donation,
		Percent:    float64(totalPennies) / float64(discountable),
	})

	if donation > 0 {
		return []Discounts{appliedDiscounts}
	}

	var res []Discounts
	for _, item := range items {
		solutions := recursiveDiscount(item, items, append(Discounts{}, appliedDiscounts...), remaining)
		res = append(res, solutions...)
	}
	return res
}
