package dto

import "time"

type (
	UserLogin struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	SendComment struct {
		UserId    int    `json:"userid"`
		ProductId int    `json:"productid"`
		Text      string `json:"text"`
	}
	GetAllComment struct {
		Id        uint      `json:"id" db:"id"`
		Text      string    `json:"commenttext" db:"commenttext"`
		ParentId  uint      `json:"parentid" db:"parentid"`
		UserId    uint      `json:"userid" db:"userid"`
		UserName  string    `json:"username" db:"username"`
		ProductId uint      `json:"productid" db:"productid"`
		CreatedAt time.Time `json:"createdat" db:"createdat"`
	}
)
