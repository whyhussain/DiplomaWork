package app

import (
	"DiplomaWork/internal/app/hanlder"
	"github.com/joho/godotenv"
	"os"

	//"DiplomaWork/internal/app/repository"
	//"DiplomaWork/internal/app/service"
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"net/http"
)

func StartNumbleServie(ctx context.Context, errCh chan<- error) {
	godotenv.Load()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
<<<<<<< Updated upstream
	urlDb := fmt.Sprintf("postgres://%s:%s@localhost:5432/project", os.Getenv("db_user"), os.Getenv("db_pass"))
=======
	godotenv.Load(".env")
	dbPass := os.Getenv("db_pass")
	dbIP := os.Getenv("db_ip")
	urlDb := "postgres://postgres:" + dbPass + "@" + dbIP + ":5432/diploma"
>>>>>>> Stashed changes
	a, Err := pgxpool.Connect(ctx, urlDb)
	if Err != nil {
		fmt.Println(Err)
	}
	defer a.Close()
	e := echo.New()
<<<<<<< Updated upstream
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
=======
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
>>>>>>> Stashed changes

}
