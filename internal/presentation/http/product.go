package http

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	dto "github.com/mohammaderm/krad/internal/dto/product"
	"github.com/mohammaderm/krad/internal/service/product"
	"github.com/mohammaderm/krad/log"
)

type (
	ProductHandler struct {
		logger         log.Logger
		productService product.ProductServiceContract
		HandlerHelper
	}
	ProductHandlerContract interface {
		GLTProduct(w http.ResponseWriter, r *http.Request)
		GetByID(w http.ResponseWriter, r *http.Request)
		GetByCategoryId(w http.ResponseWriter, r *http.Request)
	}
)

func NewProductHanlder(logger log.Logger, productservice product.ProductServiceContract) ProductHandlerContract {
	return &ProductHandler{
		logger:         logger,
		productService: productservice,
		HandlerHelper:  HandlerHelper{logger: logger},
	}
}

func (h *ProductHandler) GLTProduct(w http.ResponseWriter, r *http.Request) {
	result, err := h.productService.GLTProduct(r.Context())
	if err != nil {
		h.errorJSON(w, errors.New("failed to handle request"), http.StatusNotFound)
		h.logger.Error(&log.Field{
			Package:  "http",
			Function: "product/GLTProduct",
			Params:   "_",
			Message:  err.Error(),
		})
		return
	}
	payload := jsonResponse{
		Error:   false,
		Message: "succesfully",
		Data:    result,
	}
	h.writeJSON(w, http.StatusOK, payload)
}

func (h *ProductHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	idint, err := strconv.Atoi(id)
	if err != nil {
		h.errorJSON(w, errors.New("failed to handle request"), http.StatusBadRequest)
		return
	}
	result, err := h.productService.GetByID(r.Context(), dto.FindProductReq{Id: idint})
	if err != nil {
		h.errorJSON(w, errors.New("can not found product"), http.StatusNotFound)
		return
	}
	payload := jsonResponse{
		Error:   false,
		Message: "succesfully",
		Data:    result,
	}
	h.writeJSON(w, http.StatusOK, payload)
}

func (h *ProductHandler) GetByCategoryId(w http.ResponseWriter, r *http.Request) {
	categoryid := mux.Vars(r)["categoryid"]
	order := r.URL.Query().Get("order")
	offset := r.URL.Query().Get("offset")
	filters := r.URL.Query()["filters"]
	categoryidint, err := strconv.Atoi(categoryid)
	if err != nil {
		h.errorJSON(w, errors.New("failed to handle request"), http.StatusBadRequest)
		return
	}
	offsetint, err := strconv.Atoi(offset)
	if err != nil {
		h.errorJSON(w, errors.New("failed to handle request"), http.StatusBadRequest)
		return
	}
	result, err := h.productService.GetByCategoryId(r.Context(), dto.FindByCategoryIdReq{
		Offset: offsetint,
		Id:     categoryidint,
		Filter: filters,
		Order:  "p2." + order,
	})
	if err != nil {
		h.errorJSON(w, errors.New("can not found product"), http.StatusNotFound)
		return
	}
	payload := jsonResponse{
		Error:   false,
		Message: "succesfully",
		Data:    result,
	}
	h.writeJSON(w, http.StatusOK, payload)
}
