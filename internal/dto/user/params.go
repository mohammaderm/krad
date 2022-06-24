package dto

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
)
