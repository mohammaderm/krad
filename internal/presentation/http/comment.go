package http

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/mohammaderm/krad/log"

	dto "github.com/mohammaderm/krad/internal/dto/user"
	"github.com/mohammaderm/krad/internal/service/user"
)

type (
	CommentHandler struct {
		logger      log.Logger
		UserService user.UserServiceContracts
	}

	CommentHandlerContracts interface {
		SendComment(w http.ResponseWriter, r *http.Request)
	}
)

func NewCommentHandler(logger log.Logger, userService user.UserServiceContracts) CommentHandlerContracts {
	return &CommentHandler{
		logger:      logger,
		UserService: userService,
	}

}

func (c *CommentHandler) SendComment(w http.ResponseWriter, r *http.Request) {
	var comment dto.SendComment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		http.Error(w, "can not parse values.", http.StatusBadRequest)
		return
	}
	err = c.UserService.CreateComment(r.Context(), dto.CreateCommentReq{
		UserId:    comment.UserId,
		ProductId: comment.ProductId,
		Createdat: time.Now(),
		Text:      comment.Text,
	})
	if err != nil {
		http.Error(w, "cant save comment.", http.StatusInternalServerError)
	}
	defer r.Body.Close()
	w.WriteHeader(http.StatusCreated)
	resp := make(map[string]string)
	resp["message"] = "Sucsefully saved comment."
	jsonresp, _ := json.Marshal(resp)
	w.Write(jsonresp)
}
