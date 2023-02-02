package app

import (
	"DiplomaWork/internal/app/hanlder"
	"github.com/joho/godotenv"
	"os"

	//"DiplomaWork/internal/app/repository"
	//"DiplomaWork/internal/app/service"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
	"net/http"
)

func StartNumbleServie(ctx context.Context, errCh chan<- error) {
	godotenv.Load()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	urlDb := fmt.Sprintf("postgres://%s:%s@localhost:5432/project", os.Getenv("db_user"), os.Getenv("db_pass"))
	a, Err := pgxpool.Connect(ctx, urlDb)
	defer a.Close()

	if Err != nil {
		fmt.Println(Err)
	}
	e := echo.New()
	//repo := repository.NewDiplomaRepository(a)
	//srv := service.NewDiplomaService(repo)
	//handlers := hanlder.NewDiplomaHandler(srv)
	//e.GET("/restaraunts", handlers.GetAllRestaraunts)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	// ---------- ADD NEW CATEGORY ------
	e.POST("/categories", func(c echo.Context) error {
		// Call the getCategories function
		err := hanlder.AddCategory(c, a)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, "200")
	})
	// ------------- GET ALL CATEGORIES
	e.GET("/categories", func(c echo.Context) error {
		// Call the getCategories function
		categories, err := hanlder.GetCategories(a)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, categories)
	})
	e.GET("/categories/:id", func(c echo.Context) error {
		category, err := hanlder.GetCategory(c, a)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, category)
	})
	e.DELETE("/categories/:id", func(c echo.Context) error {
		id, err := hanlder.DeleteCategory(c, a)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, fmt.Sprintf("Category %s is deleted", id))
	})
	e.PUT("/categories/:id", func(c echo.Context) error {
		category, err := hanlder.UpdateCategory(c, a)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, category)
	})
	e.Logger.Error(e.Start(":1323"))

}
