package dto

import "time"

type (
	GetProducts struct {
		Id              uint      `db:"id"`
		ProductName     string    `db:"productname"`
		ProductNameEn   string    `db:"productnameen"`
		ImagePath       string    `db:"imagepath"`
		Available       bool      `db:"available"`
		DiscountStatus  bool      `db:"discountstatus"`
		DiscountPercent int       `db:"discountpercent"`
		Price           int       `db:"price"`
		PriceDis        int       `db:"pricedis"`
		CategoryName    string    `db:"categoryname"`
		Categoryid      uint      `db:"categoryid"`
		CreatedAt       time.Time `db:"createdat"`
	}
	GetProductsByCategory struct {
		Id              uint      `db:"id"`
		ProductName     string    `db:"productname"`
		ProductNameEn   string    `db:"productnameen"`
		ImagePath       string    `db:"imagepath"`
		Available       bool      `db:"available"`
		DiscountStatus  bool      `db:"discountstatus"`
		DiscountPercent int       `db:"discountpercent"`
		Price           int       `db:"price"`
		PriceDis        int       `db:"pricedis"`
		CreatedAt       time.Time `db:"createdat"`
	}
)
