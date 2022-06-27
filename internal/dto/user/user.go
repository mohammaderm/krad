package dto

import (
	"time"

	"github.com/mohammaderm/krad/internal/models"
)

type (
	CreateUserReq struct {
		UserName string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	GetByEmailReq struct {
		Email string `json:"email"`
	}
	GetByEmailRes struct {
		User *models.User `json:"user"`
	}
	GetByUsernameReq struct {
		UserName string `json:"username"`
	}
	GetByUsernameRes struct {
		User *models.User `json:"user"`
	}
	CreateCommentReq struct {
		UserId    int       `json:"userid" db:"userid"`
		ProductId int       `json:"productid" db:"productid"`
		Createdat time.Time `json:"createdat" db:"createdat"`
		Text      string    `json:"text" db:"text"`
	}
	GetAllCommentsRes struct {
		Commnets *[]GetAllComment `json:"comments"`
	}
	GetAllCommentsReq struct {
		ProductId int `json:"productid" db:"productid"`
		Offset    int `json:"offset"`
	}
)
