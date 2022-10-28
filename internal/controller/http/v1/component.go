package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pahan/internal/entity"
	"pahan/internal/usecase"
	"strconv"
)

type componentRoutes struct {
	t usecase.Component
}

func newComponentRoutes(handler *gin.RouterGroup, t usecase.Component) {
	r := &componentRoutes{t: t}

	handler.GET("/get_components", r.getComponents)
	handler.GET("/get_components_by_vendor_and_type", r.getComponentsByVendorAndType)
	handler.POST("/create_component", r.createComponent)
}

type componentResponse struct {
	Cmp []entity.Component `json:"components"`
}

// GetComponentByVendorIDAndTypeID godoc
// @Summary get component
// @Tags Gets
// @Description Get all components depend on vendorID and typeID
// @Param       vendor-id  query   string  true "id of a vendor"
// @Param       type-id    query   string  true "id of a type"
// @Success     200 {array}  entity.Component
// @Failure     400 {object} errResponse
// @Router      /v1/get_components_by_vendor_and_type [get]
func (f *componentRoutes) getComponentsByVendorAndType(c *gin.Context) {
	vendorIDParam := c.Query("vendor-id")
	typeIDParam := c.Query("type-id")
	vendorID, err := strconv.ParseInt(vendorIDParam, 10, 64)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	typeID, err := strconv.ParseInt(typeIDParam, 10, 64)
	components, err := f.t.GetComponentsByVendorAndType(c.Request.Context(), vendorID, typeID)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, componentResponse{components})
}

// GetComponents godoc
// @Summary get components
// @Tags Gets
// @Description Get all components
// @Success     200 {array}  entity.Component
// @Failure     500 {object} errResponse
// @Router      /v1/get_components [get]
func (f *componentRoutes) getComponents(c *gin.Context) {

	engines, err := f.t.GetComponents(c.Request.Context(), "engine")
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	bumpers, err := f.t.GetComponents(c.Request.Context(), "bumper")
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	transmissions, err := f.t.GetComponents(c.Request.Context(), "transmission")
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	doors, err := f.t.GetComponents(c.Request.Context(), "door")
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, componentsResponseBig{engines, bumpers, transmissions, doors})
}

type componentCreateRequest struct {
	VendorID       int64  `json:"vendor_id"`
	TypeID         int64  `json:"type_id"`
	Name           string `json:"name"`
	AdditionalInfo string `json:"additional_info"`
}

func (c *componentCreateRequest) validate() bool {
	return c.VendorID > 0 &&
		c.TypeID > 0 &&
		len(c.Name) > 0 &&
		len(c.AdditionalInfo) > 0
}

// CreateComponent godoc
// @Summary create component
// @Tags Posts
// @Description Create component based on params
// @Param 		request body componentCreateRequest true "query params"
// @Success     200 {object}  nil
// @Failure     500 {object} errResponse
// @Router      /v1/create_component [post]
func (f *componentRoutes) createComponent(c *gin.Context) {
	var ccr componentCreateRequest
	if err := c.ShouldBindJSON(&ccr); err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	vl := ccr.validate()
	if !vl {
		errorResponse(c, http.StatusInternalServerError, "error in validate component")
		return
	}
	err := f.t.CreateComponent(c.Request.Context(), ccr.VendorID, ccr.TypeID, ccr.Name, ccr.AdditionalInfo)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

type componentsResponseBig struct {
	Engines       []entity.ComponentVendor `json:"engines"`
	Bumpers       []entity.ComponentVendor `json:"bumpers"`
	Transmissions []entity.ComponentVendor `json:"transmissions"`
	Doors         []entity.ComponentVendor `json:"doors"`
}
