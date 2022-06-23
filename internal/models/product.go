package models

import "time"

type (
	Product struct {
		Id            uint      `db:"id"`
		ProductName   string    `db:"productname"`
		ProductNameEn string    `db:"productnameen"`
		Description   string    `db:"description"`
		CreatedAt     time.Time `db:"createdat"`
		Counter       int       `db:"counter"`
		ImagePath     string    `db:"imagepath"`
		Available     bool      `db:"available"`
		Priceid       uint      `db:"priceid"`
		Categoryid    uint      `db:"categoryid"`
	}
)
