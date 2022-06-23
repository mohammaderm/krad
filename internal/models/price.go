package models

import "time"

type (
	Price struct {
		Id              uint      `db:"id"`
		CreatedAt       time.Time `db:"createdat"`
		DiscountStatus  bool      `db:"discountstatus"`
		DiscountPercent int       `db:"discountpercent"`
		Price           int       `db:"price"`
		PriceDis        int       `db:"pricedis"`
	}
)
