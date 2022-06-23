package models

import "time"

type (
	Category struct {
		Id           uint      `db:"id"`
		CategoryName string    `db:"categoryname"`
		CreatedAt    time.Time `db:"createdat"`
	}
)
