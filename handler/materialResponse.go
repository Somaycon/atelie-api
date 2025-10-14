package handler

import "github.com/Somaycon/atelie-api/schemas"

type CreateMaterialResponse struct {
	Message string                   `json:"message"`
	Data    schemas.MaterialResponse `json:"data"`
}

type DeleteMaterialResponse struct {
	Message string                   `json:"message"`
	Data    schemas.MaterialResponse `json:"data"`
}
type UpdateMaterialResponse struct {
	Message string                   `json:"message"`
	Data    schemas.MaterialResponse `json:"data"`
}

type GetMaterialByIdResponse struct {
	Message string                   `json:"message"`
	Data    schemas.MaterialResponse `json:"data"`
}

type GetAllMaterialsResponse struct {
	Message string                     `json:"message"`
	Data    []schemas.MaterialResponse `json:"data"`
}