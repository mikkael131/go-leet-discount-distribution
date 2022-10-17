package main

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

type args struct {
	discount uint
	items    []*Item
}

var (
	tests = map[string]struct {
		args args
		want []*Item
	}{
		"find the no donation distribution": {
			args: args{
				discount: 17,
				items: []*Item{
					{Id: "1", Quantity: 11, UnitPrice: 1100},
					{Id: "2", Quantity: 5, UnitPrice: 500},
					{Id: "3", Quantity: 4, UnitPrice: 400},
				},
			},
			want: []*Item{
				{Id: "1", Quantity: 11, UnitPrice: 1100, UnitDiscount: 0, ItemDiscount: 0, DonatedDiscount: 0},
				{Id: "2", Quantity: 5, UnitPrice: 500, UnitDiscount: 1, ItemDiscount: 5, DonatedDiscount: 0},
				{Id: "3", Quantity: 4, UnitPrice: 400, UnitDiscount: 3, ItemDiscount: 12, DonatedDiscount: 0},
			},
		},
		"prioritize lowest overall discount percentage": {
			args: args{
				discount: 10,
				items: []*Item{
					{Id: "1", Quantity: 3, UnitPrice: 600},
					{Id: "2", Quantity: 2, UnitPrice: 200},
				},
			},
			want: []*Item{
				{Id: "1", Quantity: 3, UnitPrice: 600, UnitDiscount: 2, ItemDiscount: 6, DonatedDiscount: 0},
				{Id: "2", Quantity: 2, UnitPrice: 200, UnitDiscount: 2, ItemDiscount: 4, DonatedDiscount: 0},
			},
		},
		"find allocation with the least donation": {
			args: args{
				discount: 11,
				items: []*Item{
					{Id: "1", Quantity: 5, UnitPrice: 500},
					{Id: "2", Quantity: 4, UnitPrice: 400},
				},
			},
			want: []*Item{
				{Id: "1", Quantity: 5, UnitPrice: 500, UnitDiscount: 0, ItemDiscount: 0, DonatedDiscount: 0},
				{Id: "2", Quantity: 4, UnitPrice: 400, UnitDiscount: 3, ItemDiscount: 12, DonatedDiscount: 1},
			},
		},
		"exceeding discount": {
			args: args{
				discount: 1000,
				items: []*Item{
					{Id: "1", Quantity: 5, UnitPrice: 5},
					{Id: "2", Quantity: 4, UnitPrice: 4},
				},
			},
			want: []*Item{
				{Id: "1", Quantity: 5, UnitPrice: 5, UnitDiscount: 5, ItemDiscount: 25, DonatedDiscount: 0},
				{Id: "2", Quantity: 4, UnitPrice: 4, UnitDiscount: 4, ItemDiscount: 16, DonatedDiscount: 0},
			},
		},
	}
)

func Test_applyDiscountWithDonation(t *testing.T) {
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			applyDiscountWithDonation(tt.args.discount, tt.args.items)

			if diff := cmp.Diff(tt.want, tt.args.items); diff != "" {
				t.Errorf("diff(): %+v", diff)
			}
		})
	}
}

func Benchmark_applyDiscountWithDonation(b *testing.B) {
	for name, tt := range tests {
		b.Run(name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				applyDiscountWithDonation(tt.args.discount, tt.args.items)
			}
		})
	}
}
