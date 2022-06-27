package http

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
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
		GetAllComents(w http.ResponseWriter, r *http.Request)
	}
)

func NewCommentHandler(logger log.Logger, userService user.UserServiceContracts) CommentHandlerContracts {
	return &CommentHandler{
		logger:        logger,
		UserService:   userService,
		HandlerHelper: HandlerHelper{logger: logger},
	}

}

func (c *CommentHandler) GetAllComents(w http.ResponseWriter, r *http.Request) {
	productid := mux.Vars(r)["productid"]
	offset := r.URL.Query().Get("offset")
	id, err := strconv.Atoi(productid)
	if err != nil {
		c.errorJSON(w, errors.New("failed to handle request"), http.StatusBadRequest)
		return
	}
	offsetint, err := strconv.Atoi(offset)
	if err != nil {
		c.errorJSON(w, errors.New("failed to handle request"), http.StatusBadRequest)
		return
	}
	result, err := c.UserService.GetAllComments(r.Context(), dto.GetAllCommentsReq{
		ProductId: id,
		Offset:    offsetint,
	})
	if err != nil {
		c.errorJSON(w, errors.New("can not found any product"), http.StatusNotFound)
		return
	}
	payload := jsonResponse{
		Error:   false,
		Message: "succesfully",
		Data:    result,
	}
	c.writeJSON(w, http.StatusOK, payload)
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
