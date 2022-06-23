package dto

import "github.com/mohammaderm/krad/internal/models"

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
)
