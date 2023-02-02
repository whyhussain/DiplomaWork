package app

import (
	"DiplomaWork/internal/app/hanlder"
	"DiplomaWork/internal/app/repository"
	"DiplomaWork/internal/app/service"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
	"os"
)

func StartNumbleServie(ctx context.Context, errCh chan<- error) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	user, password := os.Getenv("db_user"), os.Getenv("db_pass")
	urlDb := fmt.Sprintf("postgres://%s:%s@localhost:5432/project", user, password)
	a, Err := pgxpool.Connect(ctx, urlDb)
	defer a.Close()

	if Err != nil {
		fmt.Println(Err)
	}
	e := echo.New()
	repo := repository.NewDiplomaRepository(a)
	srv := service.NewDiplomaService(repo)
	handlers := hanlder.NewDiplomaHandler(srv)
	e.GET("/restaraunts", handlers.GetAllRestaraunts)

	e.Logger.Fatal(e.Start(":1323"))

}
