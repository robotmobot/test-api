package handler

import (
	"test-api/controller"
)

type Handler struct {
	ProductController controller.ProductController
}

func NewHandler(pf controller.ProductController) *Handler {
	return &Handler{
		ProductController: pf,
	}
}
