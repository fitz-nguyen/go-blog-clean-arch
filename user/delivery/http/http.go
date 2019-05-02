package http

import (
	"blog/models"
	"blog/user"
	// "context"
	"net/http"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

// UserHandler  represent the httphandler for article
type UserHandler struct {
	AUsecase user.Usecase
}


// NewUserHTTPHandler  represent the httphandler for article
func NewUserHTTPHandler(e *echo.Echo, us user.Usecase) {
	Handler := &UserHandler{
		AUsecase: us,
	}
	e.POST("/signup", Handler.Signup)
	// e.POST("/login", h.Login)
}


// Signup  represent the httphandler for article
func (h *UserHandler) Signup(c echo.Context) (err error) {
	// Bind
	u := &models.User{ ID: bson.NewObjectId() }
	// if err = c.Bind(u); err != nil {
	// 	return
	// }

	// ctx := c.Request().Context()
	// if ctx == nil {
	// 	ctx = context.Background()
	// }

	// // Validate
	// if u.Email == "" || u.Password == "" {
	// 	return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid email or password"}
	// }

	// Save user
	if err = h.AUsecase.Save(u); err != nil {
		return c.JSON(http.StatusOK, err)
	}

	return c.JSON(http.StatusOK, u)
}
