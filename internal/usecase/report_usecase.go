package usecase

import (
	"context"
	"gym-cli/internal/domain"
	"gym-cli/internal/model/entity"
	"time"
)

type ReportUsecase struct {
	repo domain.ReportRepository
}

func NewReportUsecase(repo domain.ReportRepository) domain.ReportUsecase {
	return &ReportUsecase{repo: repo}
}

func (u *ReportUsecase) MonthlyIncome() ([]entity.IncomeEntry, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return u.repo.MonthlyIncome(ctx)
}

func (u *ReportUsecase) MvpMembers() ([]entity.MvpMember, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return u.repo.MvpMembers(ctx)
}

func (u *ReportUsecase) MemberJoins() ([]entity.MemberJoin, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return u.repo.MemberJoins(ctx)
}
