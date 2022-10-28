package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pahan/internal/entity"
	"pahan/internal/usecase"
	"strconv"
)

type factoryRoutes struct {
	t usecase.Factory
}

func newFactoryRoutes(handler *gin.RouterGroup, t usecase.Factory) {
	f := &factoryRoutes{t: t}

	handler.GET("/get_factories_by_vendor", f.getFactoriesByVendor)
}

type factoryResponse struct {
	Fct []entity.Factory `json:"factories"`
}

// GetFactoriesByVendor godoc
// @Summary list of factories
// @Tags Gets
// @Description Get all factories with current vendorID
// @Param       vendor-id  query   string  true "id of a vendor"
// @Success     200 {array}  entity.Factory
// @Failure     400 {object} errResponse
// @Router      /v1/get_factories_by_vendor [get]
func (f *factoryRoutes) getFactoriesByVendor(c *gin.Context) {
	vendorIDParam := c.Query("vendor-id")
	vendorID, err := strconv.ParseInt(vendorIDParam, 10, 64)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	factories, err := f.t.GetFactoriesByVendor(c.Request.Context(), vendorID)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, factoryResponse{factories})
}
