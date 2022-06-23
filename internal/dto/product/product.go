package dto

import "github.com/mohammaderm/krad/internal/models"

type (
	ProductRes struct {
		Products []GetProducts `json:"products"`
	}
	FindProductRes struct {
		Product *models.Product `json:"product"`
	}
	FindProductReq struct {
		Id int `json:"id"`
	}
	FindByCategoryIdReq struct {
		Offset int      `json:"offset"`
		Id     int      `json:"id"`
		Filter []string `json:"filter"`
		Order  string   `json:"order"`
	}
	FindByCategoryIdRes struct {
		Products []GetProductsByCategory `json:"products"`
	}
)
