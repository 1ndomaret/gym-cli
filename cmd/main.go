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
	uuc := usecase.NewUserUsecase(ur)
	uh := handler.NewUserHandler(uuc, reader)

	rr := dbrepo.NewReportRepository(db)
	ruc := usecase.NewReportUsecase(rr)
	rh := handler.NewReportHandler(ruc, reader)

	ir := dbrepo.NewInvoiceRepository(db)
	iuc := usecase.NewInvoiceUsecase(ir)
	ih := handler.NewInvoiceHandler(iuc, reader)

	pr := dbrepo.NewPaymentRepository(db)
	puc := usecase.NewPaymentUsecase(pr)
	ph := handler.NewPaymentHandler(puc, reader)

	app := handler.NewMenu(uh, rh, ih, ph, reader)
	app.Run()
}
