package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pahan/internal/entity"
	"pahan/internal/usecase"
)

type designRoutes struct {
	t usecase.Model
}

func newModelRoutes(handler *gin.RouterGroup, t usecase.Model) {
	r := &designRoutes{t: t}

	handler.GET("/get_models", r.getAllModels)
	handler.POST("/create_model", r.doCreateModel)
}

type modelResponse struct {
	Models []entity.ModelBig `json:"models"`
}

// GetAllModels godoc
// @Summary list of models
// @Tags Gets
// @Description Get all models
// @Success     200 {array}  entity.Model
// @Failure     400 {object} errResponse
// @Router      /v1/get_models [get]
func (r *designRoutes) getAllModels(c *gin.Context) {
	listModels, err := r.t.GetAllModels(c.Request.Context())
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, modelResponse{listModels})

}

// FIXME
// структура для тела запроса
type doDesignRequest struct {
	VendorID     int64  `json:"vendor_id"`
	Significance int64  `json:"significance"`
	Price        int64  `json:"price"`
	Name         string `json:"name"`
	EngineerID   int64  `json:"engineer_id"`
	FactoryID    int64  `json:"factory_id"`
	Wheeldrive   string `json:"wheeldrive"`
}

// CreateModel godoc
// @Summary create model
// @Tags Posts
// @Description Create model based on params
// @Param 		request body doDesignRequest true "query params"
// @Success     200 {object}  nil
// @Failure     500 {object} errResponse
// @Router      /v1/create_model [post]
func (r *designRoutes) doCreateModel(c *gin.Context) {
	var req doDesignRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := r.t.NewModel(c.Request.Context(),
		entity.Model{
			VendorID:     req.VendorID,
			Significance: req.Significance,
			Price:        req.Price,
			Name:         req.Name,
			EngineerID:   req.EngineerID,
			FactoryID:    req.FactoryID,
			WheelDrive:   req.Wheeldrive,
		})
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}
