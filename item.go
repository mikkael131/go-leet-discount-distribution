package main

type Item struct {
	Id              string
	Quantity        uint
	UnitPrice       uint // primitive for now, 1 = 1 cent
	UnitDiscount    uint // primitive for now, 1 = 1 cent
	ItemDiscount    uint // primitive for now, 1 = 1 cent
	DonatedDiscount uint // primitive for now, 1 = 1 cent
}

func (i Item) TotalGrossPrice() uint {
	return i.UnitPrice * i.Quantity
}

func (i Item) TotalDiscountedPrice() uint {
	return (i.UnitPrice - i.UnitDiscount) * i.Quantity
}
