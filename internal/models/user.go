package models

import "time"

type (
	User struct {
		Id        uint      `json:"id" db:"id"`
		Username  string    `json:"username" db:"username"`
		Email     string    `json:"email" db:"email"`
		Password  string    `json:"password" db:"password"`
		CreatedAt time.Time `db:"createdat"`
	}

	Comment struct {
		Id        uint      `json:"id" db:"id"`
		Text      string    `json:"commenttext" db:"commenttext"`
		ParentId  uint      `json:"parentid" db:"parentid"`
		UserId    uint      `json:"userid" db:"userid"`
		ProductId uint      `json:"productid" db:"productid"`
		CreatedAt time.Time `db:"createdat"`
	}

	CreateComment struct {
		UserId    int       `json:"userid" db:"userid"`
		ProductId int       `json:"productid" db:"productid"`
		Createdat time.Time `json:"createdat" db:"createdat"`
		Text      string    `json:"text" db:"text"`
	}
)
