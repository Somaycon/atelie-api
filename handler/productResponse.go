package handler

import "github.com/Somaycon/atelie-api/schemas"

type CreateProductResponse struct {
	Message string                  `json:"message"`
	Data    schemas.ProductResponse `json:"data"`
}

type DeleteProductResponse struct {
	Message string                  `json:"message"`
	Data    schemas.ProductResponse `json:"data"`
}
type UpdateProductResponse struct {
	Message string                  `json:"message"`
	Data    schemas.ProductResponse `json:"data"`
}

type GetProductByIdResponse struct {
	Message string                  `json:"message"`
	Data    schemas.ProductResponse `json:"data"`
}

type GetAllProductsResponse struct {
	Message string                    `json:"message"`
	Data    []schemas.ProductResponse `json:"data"`
}