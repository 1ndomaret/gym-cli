package db

import (
	"context"
	"database/sql"
	"gym-cli/internal/domain"
	"gym-cli/internal/model/entity"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(ctx context.Context, user *entity.User, userProfile *entity.UserProfile) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	userQuery := `
    INSERT INTO Users (Email, Password, Type) 
    	VALUES (?, ?, 'member')
	`
	result, err := tx.ExecContext(ctx, userQuery, user.Email, user.Password)
	if err != nil {
		return err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	newUserID := int(lastInsertID)

	profileQuery := `
		INSERT INTO UserProfiles (UserId, MemberTierId, FirstName, LastName, Address)
			VALUES (?, ?, ?, ?, ?);
	`
	_, err = tx.ExecContext(
		ctx,
		profileQuery,
		newUserID,
		userProfile.MemberTierId,
		userProfile.FirstName,
		userProfile.LastName,
		userProfile.Address,
	)

	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	user.UserId = newUserID

	return nil
}
