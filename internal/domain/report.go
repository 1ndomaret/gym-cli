package domain

import (
	"context"
	"gym-cli/internal/model/entity"
)

type ReportRepository interface {
	MonthlyIncome(ctx context.Context) ([]entity.MonthlyIncome, error)
}

type ReportUsecase interface {
	MonthlyIncome() ([]entity.MonthlyIncome, error)
}

type ReportHandler interface {
	MonthlyIncome()
}
