package app

import (
	"DiplomaWork/config"
	"DiplomaWork/internal/app/hanlder"
	"DiplomaWork/internal/app/repository"
	"DiplomaWork/internal/app/service"
	"github.com/joho/godotenv"

	//"DiplomaWork/internal/app/repository"
	//"DiplomaWork/internal/app/service"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
)

func StartNumbleServie(ctx context.Context, errCh chan<- error) {
	cfg := config.NewConfig()
	godotenv.Load()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	godotenv.Load(".env")
	//dbPass := os.Getenv("db_pass")
	//dbIP := os.Getenv("db_ip")
	//urlDb := "postgres://postgres:" + dbPass + "@" + dbIP + ":5432/diploma"
	urlDb := cfg.Database.Primary

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
	e.GET("/categories/:id", handlers.GetCategoryById)
	e.POST("/categories", handlers.AddCategory)
	e.PUT("/categories/:id", handlers.UpdateCategory)
	e.DELETE("/categories/:id", handlers.DeleteCategory)

	e.GET("/menus", handlers.GetAllMenu)
	e.GET("/menus/:id", handlers.GetMenuById)
	e.POST("/menus", handlers.AddMenu)
	e.DELETE("/menus/:id", handlers.DeleteMenuById)
	e.PUT("/menus/:id", handlers.UpdateMenuById)

	e.GET("/partners", handlers.GetPartners)
	e.GET("/partners/:id", handlers.GetPartnerById)
	e.POST("/partners", handlers.AddPartner)
	e.PUT("/partners/:id", handlers.UpdatePartnerById)
	e.DELETE("/partners/:id", handlers.DeletePartnerById)

	e.GET("/admins", handlers.GetAdmins)
	e.GET("/admins/:id", handlers.GetAdminById)
	e.POST("/admins", handlers.AddAdmin)
	e.PUT("/admins/:id", handlers.UpdateAdminById)
	e.DELETE("/admins/:id", handlers.DeleteAdminById)

	e.GET("/techsupports", handlers.GetTechSupports)
	e.GET("/techsupports/:id", handlers.GetTechSupportById)
	e.POST("/techsupports", handlers.AddTechSupport)
	e.PUT("/techsupports/:id", handlers.UpdateTechSupportById)
	e.DELETE("/techsupports/:id", handlers.DeleteTechSupportById)

	e.GET("/customers", handlers.GetCustomers)
	e.GET("/customers/:id", handlers.GetCustomerById)
	e.POST("/customers", handlers.AddCustomer)
	e.PUT("/customers/:id", handlers.UpdateCustomerById)
	e.DELETE("/customers/:id", handlers.DeleteCustomerById)

	e.GET("/deliveryPersonnels", handlers.GetDeliveryPersonnels)
	e.GET("/deliveryPersonnels/:id", handlers.GetDeliveryPersonnelById)
	e.POST("/deliveryPersonnels", handlers.AddDeliveryPersonnel)
	e.PUT("/deliveryPersonnels/:id", handlers.UpdateDeliveryPersonnelById)
	e.DELETE("/deliveryPersonnels/:id", handlers.DeleteDeliveryPersonnelById)

	e.GET("/reviews", handlers.GetReviews)
	e.GET("/reviews/:id", handlers.GetReviewById)
	e.POST("/reviews", handlers.AddReview)
	e.PUT("/reviews/:id", handlers.UpdateReviewById)
	e.DELETE("/reviews/:id", handlers.DeleteReviewById)

	e.GET("/schedules", handlers.GetSchedules)
	e.GET("/schedules/:id", handlers.GetScheduleById)
	e.POST("/schedules", handlers.AddSchedule)
	e.PUT("/schedules/:id", handlers.UpdateScheduleById)
	e.DELETE("/schedules/:id", handlers.DeleteScheduleById)

	e.Logger.Fatal(e.Start(":8080"))

}
