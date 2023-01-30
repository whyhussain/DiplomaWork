package app

import (
	"DiplomaWork/internal/app/hanlder"
	"DiplomaWork/internal/app/repository"
	"DiplomaWork/internal/app/service"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
)

func StartNumbleServie(ctx context.Context, errCh chan<- error) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	urlDb := "postgres://postgres:JHsdayd78231@localhost:5432/parse"
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
