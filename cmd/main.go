package main

import (
	"bufio"
	"gym-cli/internal/config"
	"gym-cli/internal/handler"
	dbrepo "gym-cli/internal/repository/db"
	"gym-cli/internal/usecase"
	"log"
	"os"
)

func main() {
	db, err := config.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	reader := bufio.NewReader(os.Stdin)

	ur := dbrepo.NewUserRepository(db)
	uc := usecase.NewUserUsecase(ur)
	uh := handler.NewUserHandler(uc, reader)

	rr := dbrepo.NewReportRepository(db)
	ruc := usecase.NewReportUsecase(rr)
	rh := handler.NewReportHandler(ruc, reader)

	app := handler.NewMenu(uh, rh, reader)
	app.Run()
}
