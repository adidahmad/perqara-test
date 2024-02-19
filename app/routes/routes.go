package routes

import (
	"database/sql"

	userCtrl "github.com/adidahmad/perqara-test/app/controllers/users"
	userRepo "github.com/adidahmad/perqara-test/core/users/repository"
	userSrvc "github.com/adidahmad/perqara-test/core/users/service"
	"github.com/adidahmad/perqara-test/modules/databases/gorm"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type RouteDependency struct {
	UsersController userCtrl.Controller
}

const (
	APIV1 = "/api/v1"
)

func New(dbConn *sql.DB) *RouteDependency {
	gorm.GormInit(dbConn)

	userRepository := userRepo.NewUsersRepository(gorm.DB)
	userService := userSrvc.NewUsersService(userRepository)

	return &RouteDependency{
		UsersController: *userCtrl.NewController(userService),
	}
}

func (rd RouteDependency) RegisterRoutes(e *echo.Echo) {
	route := e.Group(APIV1)

	user := route.Group("/users")
	user.POST("", rd.UsersController.Create)
	user.GET("", rd.UsersController.GetList)
	user.GET("/:id", rd.UsersController.GetById)
	user.PUT("/:id", rd.UsersController.Update)
	user.DELETE("/:id", rd.UsersController.DeleteById)

	route.Any("/swagger/*any", echoSwagger.WrapHandler)
}
