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

func (r *userRepository) MemberList(ctx context.Context) ([]entity.UserDetail, error) {
	query := `
		SELECT u.Email, up.FirstName, up.LastName, 
			t.TierName, up.Address, up.CreatedAt, up.Status
			FROM UserProfiles up
			JOIN Users u ON u.UserId = up.UserId
			JOIN Tiers t ON t.TierId = up.MemberTierId
			WHERE u.Type = "member"
			ORDER BY up.CreatedAt ASC
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userDetails []entity.UserDetail
	for rows.Next() {
		var userDetail entity.UserDetail
		err := rows.Scan(
			&userDetail.Email,
			&userDetail.FirstName,
			&userDetail.LastName,
			&userDetail.TierName,
			&userDetail.Address,
			&userDetail.CreatedAt,
			&userDetail.Status,
		)
		if err != nil {
			return nil, err
		}
		userDetails = append(userDetails, userDetail)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return userDetails, nil
}
