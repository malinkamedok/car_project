package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pahan/internal/entity"
	"pahan/internal/usecase"
)

type typeRouter struct {
	t usecase.Type
}

func newTypeRouter(handler *gin.RouterGroup, t usecase.Type) {
	r := &typeRouter{t: t}

	handler.GET("/get_types", r.getTypes)
}

type typeResponse struct {
	Types []entity.Type `json:"types"`
}

// GetTypes godoc
// @Summary list of types
// @Tags Gets
// @Description Get all types
// @Success     200 {array}  entity.Type
// @Failure     400 {object} errResponse
// @Router      /v1/get_types [get]
func (t *typeRouter) getTypes(c *gin.Context) {
	types, err := t.t.GetTypes(c.Request.Context())
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, typeResponse{types})
}
