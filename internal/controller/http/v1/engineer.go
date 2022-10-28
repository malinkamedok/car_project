package v1

import (
	"net/http"
	"pahan/internal/entity"
	"pahan/internal/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type engineerRoute struct {
	en usecase.Engineer
}

func newEngineerRoutes(handler *gin.RouterGroup, en usecase.Engineer) {
	r := engineerRoute{en: en}

	handler.GET("/get_engineers_by_vendor", r.getEngineerByIdVendor)
}

type engineerResponse struct {
	Engineer []entity.Engineer `json:"engineer"`
}

// GetEngineersByVendorID godoc
// @Summary list of engineers
// @Tags Gets
// @Description Get all engineers with current vendorID
// @Param       vendor-id  query   string  true "id of a vendor"
// @Success     200 {array}  entity.Engineer
// @Failure     400 {object} errResponse
// @Router      /v1/get_engineers_by_vendor [get]
func (en *engineerRoute) getEngineerByIdVendor(c *gin.Context) {

	vendorIDString := c.Query("vendor-id")
	vendorID, _ := strconv.Atoi(vendorIDString)
	listEngineer, err := en.en.GetAllEngineerByIdVendor(c.Request.Context(), int64(vendorID))
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, engineerResponse{listEngineer})
}
