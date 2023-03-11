package app

import (
	"DiplomaWork/internal/app/hanlder"
	"DiplomaWork/internal/app/repository"
	"DiplomaWork/internal/app/service"
	"os"

	"github.com/joho/godotenv"

	//"DiplomaWork/internal/app/repository"
	//"DiplomaWork/internal/app/service"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
)

func StartNumbleServie(ctx context.Context, errCh chan<- error) {
	godotenv.Load()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	godotenv.Load(".env")
	dbPass := os.Getenv("db_pass")
	dbIP := os.Getenv("db_ip")
	urlDb := "postgres://postgres:" + dbPass + "@" + dbIP + ":5432/diploma"
	a, Err := pgxpool.Connect(ctx, urlDb)
	if Err != nil {
		fmt.Println(Err)
	}
	defer a.Close()
	e := echo.New()
	repo := repository.NewDiplomaRepository(a)
	srv := service.NewDiplomaService(repo)
	handlers := hanlder.NewDiplomaHandler(srv)
	e.GET("/restaurants", handlers.GetAllRestaurants)
	e.GET("/restaurants/:id", handlers.GetRestaurantById)
	e.PUT("/restaurants/:id", handlers.UpdateRestaurantById)
	e.DELETE("/restaurants/:id", handlers.DeleteRestaurantById)
	e.POST("/restaurants", handlers.AddRestaurant)
	e.GET("/categories", handlers.GetCategories)
	e.POST("/categories", handlers.AddCategory)
	e.GET("/menus", handlers.GetAllMenu)
	e.GET("/menus/:id", handlers.GetMenuById)
	e.POST("/menus", handlers.AddMenu)
	e.DELETE("/menus/:id", handlers.DeleteMenuById)
	e.PUT("/menus/:id", handlers.UpdateMenuById)

	e.Logger.Fatal(e.Start(":8080"))

}
