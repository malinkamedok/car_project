package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pahan/internal/usecase"
)

type loginRoutes struct {
	t usecase.Login
}

func newLoginRoutes(handler *gin.RouterGroup, t usecase.Login) {
	r := &loginRoutes{t: t}
	handler.GET("/login", r.login)
}

type loginResponse struct {
	UserID int64
	Name   string
}

func (l *loginRoutes) login(c *gin.Context) {
	name := c.Query("name")
	typeUser := c.Query("type")
	userID, err := l.t.Login(c.Request.Context(), typeUser, name)
	if err != nil {
		errorResponse(c, http.StatusForbidden, err.Error())
		return
	}
	c.JSON(http.StatusOK, loginResponse{
		UserID: userID,
		Name:   name,
	})
}
