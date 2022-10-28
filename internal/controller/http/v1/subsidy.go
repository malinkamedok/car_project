package v1

import (
	"net/http"
	"pahan/internal/entity"
	"pahan/internal/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type subsidyRoutes struct {
	t usecase.Subsidy
}

func newSubsidyRoutes(handler *gin.RouterGroup, t usecase.Subsidy) {
	r := subsidyRoutes{t: t}

	handler.GET("/get_subsidies", r.getSubsidies)
	handler.GET("/get_subsidies_by_country", r.getSubsidiesByCountry)
	handler.POST("/create_subsidy", r.createSubsidy)
	handler.POST("/accept_subsidy", r.acceptSubsidy)
}

type subsidyResponse struct {
	Subsidy []entity.Subsidy `json:"subsidy"`
}

type subsidyResponseCountry struct {
	Subsidy []entity.SubsidyCountry `json:"subsidy"`
}

// GetSubsidies godoc
// @Summary list of subsidies
// @Tags Gets
// @Description Get all subsidies
// @Success     200 {array}  entity.Subsidy
// @Failure     500 {object} errResponse
// @Router      /v1/get_subsidies [get]
func (s *subsidyRoutes) getSubsidies(c *gin.Context) {
	listSubsidies, err := s.t.GetAllSubsidies(c.Request.Context())
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, subsidyResponseCountry{listSubsidies})
}

func (s *subsidyRoutes) getSubsidiesByCountry(c *gin.Context) {
	vendorIDParam := c.Query("country-id")
	vendorID, err := strconv.ParseInt(vendorIDParam, 10, 64)
	listSubsidies, err := s.t.GetSubsidyByCountry(c.Request.Context(), vendorID)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, subsidyResponse{listSubsidies})
}

type createSubsidyRequest struct {
	CountryIDBy    int64   `json:"country-id-by"`
	RequirePriceBy float64 `json:"require-price-by"`
	RequiredWdBy   string  `json:"required-wd-by"`
}

func (c *createSubsidyRequest) validate() bool {
	return len(c.RequiredWdBy) == 3 &&
		c.CountryIDBy > 0 &&
		c.RequirePriceBy > 0
}

// CreateSubsidy godoc
// @Summary create subsidy
// @Tags Posts
// @Description Create and link subsidy with dependent values
// @Param 		request body createSubsidyRequest true "query params"
// @Success     200 {object} nil
// @Failure     400 {object} errResponse
// @Router      /v1/create_subsidy [post]
func (s *subsidyRoutes) createSubsidy(c *gin.Context) {
	var sbs createSubsidyRequest
	if err := c.ShouldBindJSON(&sbs); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	isValid := sbs.validate()
	if !isValid {
		errorResponse(c, http.StatusBadRequest, "error in validating subsidy request")
		return
	}
	err := s.t.CreateSubsidy(c.Request.Context(), sbs.CountryIDBy, sbs.RequirePriceBy, sbs.RequiredWdBy)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

type createAcceptSubsidyRequest struct {
	SubsidyID               int64  `json:"subsidyId"`
	VendorID                int64  `json:"vendor-id"`
	Name                    string `json:"name"`
	Significance            int64  `json:"significance"`
	EngineerID              string `json:"engineer-id"`
	FactoryID               string `json:"factory-id"`
	ComponentEngineID       string `json:"component-engine-id"`
	ComponentDoorID         string `json:"component-door-id"`
	ComponentBumperID       string `json:"component-bumper-id"`
	ComponentTransmissionID string `json:"component-transmission-id"`
}

// AcceptSubsidy godoc
// @Summary accept subsidy
// @Tags Posts
// @Description Accept and link subsidy
// @Param 		request body createAcceptSubsidyRequest true "query params"
// @Success     200 {object} nil
// @Failure     400 {object} errResponse
// @Router      /v1/accept_subsidy [post]
func (s *subsidyRoutes) acceptSubsidy(c *gin.Context) {
	var cars createAcceptSubsidyRequest
	if err := c.ShouldBindJSON(&cars); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	EngineerIdInt, _ := strconv.ParseInt(cars.EngineerID, 10, 64)
	FactoryIdInt, _ := strconv.ParseInt(cars.FactoryID, 10, 64)
	EngineInt, _ := strconv.ParseInt(cars.ComponentEngineID, 10, 64)
	DoorInt, _ := strconv.ParseInt(cars.ComponentDoorID, 10, 64)
	BumperInt, _ := strconv.ParseInt(cars.ComponentBumperID, 10, 64)
	TransInt, _ := strconv.ParseInt(cars.ComponentTransmissionID, 10, 64)

	err := s.t.AcceptSubsidyUs(
		c.Request.Context(),
		cars.SubsidyID,
		entity.Model{
			VendorID:     cars.VendorID,
			Name:         cars.Name,
			Significance: cars.Significance,
			EngineerID:   EngineerIdInt,
			FactoryID:    FactoryIdInt,
		},
		EngineInt,
		DoorInt,
		BumperInt,
		TransInt,
	)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}
