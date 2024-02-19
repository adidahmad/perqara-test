package users

import (
	"net/http"

	"github.com/adidahmad/perqara-test/core/users/domain"
	"github.com/adidahmad/perqara-test/core/users/port"

	"github.com/labstack/echo/v4"
)

type errResponse struct {
	Message string `json:"message"`
}

type Controller struct {
	UsersService port.IUsersService
}

func NewController(usersService port.IUsersService) *Controller {
	return &Controller{
		UsersService: usersService,
	}
}

// Create User godoc
//
// @Summary 		Create a new user
// @Description 	Create a new user with the provided email and password
// @Tags			Users
// @Accept 			json
// @Produce 		json
// @Param 			request body CreateRequest true "User data"
// @Success 		201 {object} domain.Users
// @Failure 		400 {object} errResponse
// @Router 			/users [post]
func (ctrl *Controller) Create(c echo.Context) error {
	bodyReq := CreateRequest{}
	if err := c.Bind(&bodyReq); err != nil {
		return c.JSON(http.StatusBadRequest, errResponse{
			Message: err.Error(),
		})
	}

	if err := c.Validate(bodyReq); err != nil {
		return c.JSON(http.StatusBadRequest, errResponse{
			Message: err.Error(),
		})
	}

	data := domain.CreateUserRequest{
		Email:    bodyReq.Email,
		Password: bodyReq.Password,
	}
	res, err := ctrl.UsersService.Create(data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, res)
}

// Get List User godoc
//
// @Summary 		Get List user
// @Description 	Get list user
// @Tags			Users
// @Accept 			json
// @Produce 		json
// @Success 		200 {array} domain.Users
// @Failure 		400 {object} errResponse
// @Router 			/users [get]
func (ctrl *Controller) GetList(c echo.Context) error {
	res, err := ctrl.UsersService.GetList()
	if err != nil {
		return c.JSON(http.StatusBadRequest, errResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": res,
	})
}

// Get User By ID godoc
//
// @Summary 		Get User By ID
// @Description 	Get user by id
// @Tags			Users
// @Accept 			json
// @Produce 		json
// @Param			id	path	int	true	"user id"
// @Success 		200 {object} domain.Users
// @Failure 		400 {object} errResponse
// @Router 			/users/:id [get]
func (ctrl *Controller) GetById(c echo.Context) error {
	id := c.Param("id")
	res, err := ctrl.UsersService.GetById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

// Update User By ID godoc
//
// @Summary 		Update User By ID
// @Description 	Update user by id with the provided email and password
// @Tags			Users
// @Accept 			json
// @Produce 		json
// @Param 			request body UpdateRequest true "User data"
// @Param			id	path	int	true	"user id"
// @Success 		201 {object} domain.Users
// @Failure 		400 {object} errResponse
// @Router 			/users/:id [put]
func (ctrl *Controller) Update(c echo.Context) error {
	id := c.Param("id")

	bodyReq := UpdateRequest{}
	if err := c.Bind(&bodyReq); err != nil {
		return c.JSON(http.StatusBadRequest, errResponse{
			Message: err.Error(),
		})
	}

	if err := c.Validate(bodyReq); err != nil {
		return c.JSON(http.StatusBadRequest, errResponse{
			Message: err.Error(),
		})
	}

	data := domain.UpdateUserRequest{
		Email:    bodyReq.Email,
		Password: bodyReq.Password,
		IsActive: bodyReq.IsActive,
	}
	res, err := ctrl.UsersService.Update(id, data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

// Delete User By ID godoc
//
// @Summary 		Delete User By ID
// @Description 	Delete user by id
// @Tags			Users
// @Accept 			json
// @Produce 		json
// @Param			id	path	int	true	"user id"
// @Success 		200 {object} domain.Users
// @Failure 		400 {object} errResponse
// @Router 			/users/:id [delete]
func (ctrl *Controller) DeleteById(c echo.Context) error {
	id := c.Param("id")
	err := ctrl.UsersService.DeleteById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "deleted successfully",
	})
}
