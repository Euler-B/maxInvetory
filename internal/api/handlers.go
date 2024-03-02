package api

import (
	"errors"
	"github.com/euler-b/maxInventoryProject/internal/models"
	"log"
	"net/http"

	"github.com/euler-b/maxInventoryProject/encryption"
	"github.com/euler-b/maxInventoryProject/internal/api/dtos"
	"github.com/euler-b/maxInventoryProject/internal/service"
	"github.com/labstack/echo/v4"
)

type responseMessage struct {
	Message string `json:"message"`
}

func (a *API) RegisterUser(c echo.Context) error {
	ctx := c.Request().Context()
	params := &dtos.RegisterUser{}

	err := c.Bind(&params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid Request"})
	}

	err = a.dataValidator.Struct(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	err = a.serv.RegisterUser(ctx, params.Email, params.Name, params.Password)
	if err != nil {
		if err == service.ErrUserAlreadyExists {
			return c.JSON(http.StatusConflict, responseMessage{Message: "User already exists"})
		}
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Internal server error"})
	}
	return c.JSON(http.StatusCreated, nil)
}

func (a *API) LoginUser(c echo.Context) error {
	ctx := c.Request().Context()
	params := dtos.LoginUser{}

	err := c.Bind(&params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid request"})
	}

	err = a.dataValidator.Struct(params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	u, err := a.serv.LoginUser(ctx, params.Email, params.Password)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Internal server error"})
	}

	token, err := encryption.SignedLoginToken(u)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Internal server error"})
	}

	cookie := &http.Cookie{
		Name:     "Authorization",
		Value:    token,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
		HttpOnly: true,
		Path: "/",
	}

	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, map[string]string{"success": "true"})
}

func (a *API) AddProducts(c echo.Context) error {
	// obtenemos la autenticacion de la cookie
	cookie, err := c.Cookie("Authorization")
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusUnauthorized, responseMessage{Message: "Unauthorized"})
	}

	// parseamos el JWT obtenido
	claims, err := encryption.ParseLoginJWT(cookie.Value)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusUnauthorized, responseMessage{Message: "Unauthorized"})
	}
	email := claims["email"].(string)

	ctx := c.Request().Context()

	// obtenemos la informacion del request
	params := dtos.AddProduct{}
	err = c.Bind(&params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid request"})
	}

	err = a.dataValidator.Struct(params)
	if err != nil {
		log.Println(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	p := models.Product{
		Name:        params.Name,
		Description: params.Description,
		Price:       params.Price,
	}

	err = a.serv.AddProduct(ctx, p, email)
	if err != nil {
		log.Println(err)

		if errors.Is(err, service.ErrInvalidPermission) {
			return c.JSON(http.StatusForbidden, responseMessage{Message: "Invalid Permissions"})
		}
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Internal Server Error"})
	}
	return c.JSON(http.StatusOK, nil)
}
