package main

import (
	"context"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"

	"github.com/mehkey/go-pastebin-web-service/internal/datasource"
	"github.com/mehkey/go-pastebin-web-service/internal/handler"
	"github.com/mehkey/go-pastebin-web-service/pkg/database"
	"github.com/mehkey/go-pastebin-web-service/pkg/middleware"

	_ "github.com/joho/godotenv/autoload"
)

type server struct{}

func Chain(h echo.HandlerFunc, middleware ...func(echo.HandlerFunc) echo.HandlerFunc) echo.HandlerFunc {
	for _, m := range middleware {
		h = m(h)
	}
	return h
}

func main() {
	//get the DATABASE_CONNECTION_URL
	connStr := os.Getenv("DATABASE_CONNECTION_URL")

	//create the pool
	pool, err := database.PGPool(context.Background(), connStr)
	if err != nil {
		log.Fatalln(err)
	}
	defer pool.Close()

	p := datasource.NewPostgres(pool)

	h := handler.NewHandler(p)

	e := echo.New()

	specialLogger := echoMiddleware.LoggerWithConfig(echoMiddleware.LoggerConfig{
		Format: "time=${time_rfc3339} method=${method}, uri=${uri}, status=${status}, latency=${latency_human}, \n",
	})
	e.Use(middleware.Logger, specialLogger)

	auth := e.Group("/auth")
	auth.Use(middleware.JWT)
	auth.GET("/test", handler.Authenticated)
	api := e.Group("/api/v1")

	api.GET("/users", h.GetAllUsers)
	api.GET("/pastebins", h.GetAllPastebins)

	api.GET("/users/:id", h.GetUserByID)
	//api.GET("/pastebin/:id", h.GetPastebinByID)

	api.GET("/pastebins/user/:userID", h.GetPastebinsForUser)

	api.POST("/users", h.CreateNewUser)
	//api.POST("/users/:id/pastebins", h.AddUserPastebin)

	port := "7999"

	e.Logger.Fatal(e.Start(":" + port))

}
