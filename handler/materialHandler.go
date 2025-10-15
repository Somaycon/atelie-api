package handler

import (
	"fmt"
	"net/http"

	"github.com/Somaycon/atelie-api/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary List materials
// @Description List all materials
// @Tags Material
// @Accept json
// @Produce json
// @Success 200 {object} GetAllMaterialsResponse
// @Failure 500 {object} ErrorResponse
// @Router /materials [get]
func GetAllMaterialsHandler(ctx *gin.Context) {
	materials := []schemas.Materials{}

	if err := db.Find(&materials).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing materials")
		return
	}
	sendSucess(ctx, "get-all-materials", materials)
}

// @BasePath /api/v1

// @Summary Show material
// @Description Show a material
// @Tags Material
// @Accept json
// @Produce json
// @Param id query string true "Material identification"
// @Success 200 {object} GetMaterialByIdResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /material/{id} [get]
func GetMaterialByIdHandler(ctx *gin.Context){
	id := ctx.Param("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, "id is required")
		return
	}
	material := schemas.Materials{}

	if err := db.First(&material, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("material with id: %s not found", id))
		return 
	}
	sendSucess(ctx, "get-material-by-id", material)
}

// @BasePath /api/v1

// @Summary Create material
// @Description Create a new material
// @Tags Material
// @Accept json
// @Produce json
// @Param request body CreateMaterialRequest true "Request body"
// @Success 200 {object} CreateMaterialResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /material [post]
func CreateMaterialHandler(ctx *gin.Context) {
	request := CreateMaterialRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	material := schemas.Materials{
		Name:         request.Name,
		CurrentStock:     request.CurrentStock,
		PricePerUnit: request.PricePerUnit,
	}

	if err := db.Create(&material).Error; err != nil{
		logger.Errorf("error creating material: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating material on database")
		return
	}
	sendSucess(ctx, "create-material", material)
}



// @BasePath /api/v1

// @Summary Update material
// @Description Update a material
// @Tags Material
// @Accept json
// @Produce json
// @Param id query string true "Material Identification"
// @Param opening body UpdateMaterialRequest true "Material data to Update"
// @Success 200 {object} UpdateMaterialResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /material/{id} [put]
func UpdateMaterialHandler(ctx *gin.Context){
	id := ctx.Param("id")
	request := UpdateMaterialRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if id == ""{
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
	}

	material := schemas.Materials{}

	if err := db.First(&material, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("material with id: %s not found", id))
		return
	}

	if request.Name != "" {
		material.Name = request.Name
	}

	if request.CurrentStock >= 0 {
		material.CurrentStock = request.CurrentStock
	}

	if request.PricePerUnit >= 0 {
		material.PricePerUnit = request.PricePerUnit
	}

	if err := db.Save(&material).Error; err != nil {
		logger.Errorf("error updating material: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating material on database")
		return
	}
	sendSucess(ctx, "update-material", material)
}

// @BasePath /api/v1

// @Summary Delete material
// @Description Delete a material
// @Tags Material
// @Accept json
// @Produce json
// @Param id query string true "Material identification"
// @Success 200 {object} DeleteMaterialResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /material/{id} [delete]
func DeleteMaterialHandler(ctx *gin.Context) {
	id:= ctx.Param("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	material := schemas.Materials{}

	if err := db.First(&material, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("material with id: %s not found", id))
		return
	}

	if err := db.Delete(&material).Error; err != nil {
		logger.Errorf("error deleting material: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error deleting material from database")
		return
	}
	sendSucess(ctx, "delete-material", material)
}
