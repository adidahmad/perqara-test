package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/adidahmad/perqara-test/app/routes"
	"github.com/adidahmad/perqara-test/config"
	"github.com/adidahmad/perqara-test/docs"
	"github.com/go-playground/validator"

	"github.com/adidahmad/perqara-test/modules/databases/mysql"
	customValidator "github.com/adidahmad/perqara-test/utils/validator"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/adidahmad/perqara-test/docs"
)

func Start() {
	docs.SwaggerInfo.Title = "CRUD API"
	docs.SwaggerInfo.Description = "CRUD API Documentation"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:1125"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	config.Load()

	e := echo.New()

	e.Debug = config.AppConf.DebugMode

	e.Use(middleware.Logger())

	e.Use(middleware.Recover())

	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.Secure())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	e.Validator = &customValidator.BodyRequestValidator{
		Validator: CustomValidator(),
	}

	dbConn, err := mysql.MySQLConnect(config.AppConf.DBConfig)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	route := routes.New(dbConn)
	route.RegisterRoutes(e)

	go func() {
		address := fmt.Sprintf(":%s", config.AppConf.Port)
		if err := e.Start(address); err != nil {
			dbConn.Close()
			log.Fatal(err)
			panic(err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// a timeout of 10 seconds to shutdown the server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	dbConn.Close()
}

func CustomValidator() *validator.Validate {
	customValidator := validator.New()
	return customValidator
}
