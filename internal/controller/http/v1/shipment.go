package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pahan/internal/entity"
	"pahan/internal/usecase"
	"time"
)

type shipmentRoutes struct {
	t usecase.Shipment
}

func newShipmentRoutes(handler *gin.RouterGroup, t usecase.Shipment) {
	r := &shipmentRoutes{t: t}

	handler.GET("/get_shipments", r.getShipments)
	handler.POST("/create_shipment", r.doNewShipment)
}

type createShipmentRequest struct {
	OrderID   int64  `json:"order_id"`
	CountryID int64  `json:"country_id"`
	Date      string `json:"date"`
}

// CreateShipment godoc
// @Summary create new shipment
// @Tags Posts
// @Description Create and link new shipment
// @Param 		request body createShipmentRequest true "query params"
// @Success     200 {object} nil
// @Failure     400 {object} errResponse
// @Failure     500 {object} errResponse
// @Router      /v1/create_shipment [post]
func (s *shipmentRoutes) doNewShipment(c *gin.Context) {
	var request createShipmentRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	t, err := time.Parse(time.RFC3339, request.Date)

	if err != nil {
		errorResponse(c, http.StatusBadRequest, "error in parsing time")
		return
	}

	err = s.t.CreateShipment(c.Request.Context(),
		entity.Shipment{
			OrderID:     request.OrderID,
			CountryToID: request.CountryID,
			Date:        t,
		})
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

type shipmentsResponse struct {
	Shipments []entity.Shipment `json:"shipments"`
}

// GetShipments godoc
// @Summary get shipments
// @Tags Gets
// @Description Get all shipments
// @Success     200 {array} entity.Shipment
// @Failure     500 {object} errResponse
// @Router      /v1/get_shipments [get]
func (s *shipmentRoutes) getShipments(c *gin.Context) {
	shipments, err := s.t.GetShipments(c.Request.Context())
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, shipmentsResponse{shipments})
}
