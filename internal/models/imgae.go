package models

import (
	"time"
)

type (
	Image struct {
		Id        uint      `db:"id"`
		CreatedAt time.Time `db:"createdat`
		ProductId Product   `db:"productid`
		ImagePath string    `db:"imagepath`
	}
)
