package main

type Discounts []Discount

func (d Discounts) TotalDonation() (res uint) {
	for _, discount := range d {
		res += discount.Donation
	}
	return
}

func (d Discounts) SumPercent() (res float64) {
	for _, discount := range d {
		res += discount.Percent
	}
	return
}

type Discount struct {
	ItemId     string
	UnitValue  uint
	TotalValue uint
	Donation   uint
	Percent    float64
}
