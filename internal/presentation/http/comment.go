package http

import (
	"errors"
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
		HandlerHelper
	}

	CommentHandlerContracts interface {
		SendComment(w http.ResponseWriter, r *http.Request)
	}
)

func NewCommentHandler(logger log.Logger, userService user.UserServiceContracts) CommentHandlerContracts {
	return &CommentHandler{
		logger:        logger,
		UserService:   userService,
		HandlerHelper: HandlerHelper{logger: logger},
	}

}

func (c *CommentHandler) SendComment(w http.ResponseWriter, r *http.Request) {
	var comment dto.SendComment
	err := c.readJSON(w, r, &comment)
	if err != nil {
		c.errorJSON(w, errors.New("can not parse values"), http.StatusNotFound)
		return
	}
	comm := dto.CreateCommentReq{
		UserId:    comment.UserId,
		ProductId: comment.ProductId,
		Createdat: time.Now(),
		Text:      comment.Text,
	}
	err = c.UserService.CreateComment(r.Context(), comm)
	if err != nil {
		c.errorJSON(w, errors.New("cant save comment"), http.StatusInternalServerError)
	}
	defer r.Body.Close()
	payload := jsonResponse{
		Error:   false,
		Message: "your comment saved succesfully",
		Data:    comm,
	}
	c.writeJSON(w, http.StatusOK, payload)
}
