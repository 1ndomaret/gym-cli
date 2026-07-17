package domain

import (
	"context"
	"gym-cli/internal/model/entity"
)

type ReportRepository interface {
	MonthlyIncome(ctx context.Context) ([]entity.IncomeEntry, error)
	MvpMembers(ctx context.Context) ([]entity.MvpMember, error)
	MemberJoins(ctx context.Context) ([]entity.MemberJoin, error)
}

type ReportUsecase interface {
	MonthlyIncome() ([]entity.IncomeEntry, error)
	MvpMembers() ([]entity.MvpMember, error)
	MemberJoins() ([]entity.MemberJoin, error)
}

type ReportHandler interface {
	MonthlyIncome()
	MvpMembers()
	MemberJoins()
}
