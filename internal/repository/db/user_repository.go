package db

import (
	"context"
	"database/sql"
	"errors"
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
		SELECT up.UserProfileId, u.UserId, u.Email, up.FirstName, up.LastName, 
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
			&userDetail.UserProfileId,
			&userDetail.UserId,
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

// USER VIEW
func (r *userRepository) ViewSchedule(ctx context.Context, userID int) ([]entity.Event, error) {
	query := `
		SELECT e.EventId, e.EventName, e.MinTierRank, e.Schedule
		FROM UserProfiles up
		JOIN Tiers t ON t.TierId = up.MemberTierId
		JOIN Events e ON e.MinTierRank <= t.TierRank
		WHERE up.UserId = ?
		  AND e.Schedule >= NOW()
		ORDER BY e.Schedule
	`
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []entity.Event
	for rows.Next() {
		var e entity.Event
		if err := rows.Scan(&e.EventID, &e.EventName, &e.MinTierRank, &e.Schedule); err != nil {
			return nil, err
		}
		events = append(events, e)
	}
	return events, rows.Err()
}

func (r *userRepository) ViewPendingPayment(ctx context.Context, userID int) ([]entity.Invoice, error) {
	query := `
		SELECT i.InvoiceId, i.Amount, i.DueDate, i.InvoiceStatus
		FROM Invoices i
		JOIN UserProfiles up ON up.UserProfileId = i.UserProfileId
		WHERE up.UserId = ?
		  AND i.InvoiceStatus IN ('pending', 'overdue')
		ORDER BY i.DueDate
	`
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var invoices []entity.Invoice
	for rows.Next() {
		var inv entity.Invoice
		if err := rows.Scan(&inv.InvoiceID, &inv.Amount, &inv.DueDate, &inv.InvoiceStatus); err != nil {
			return nil, err
		}
		invoices = append(invoices, inv)
	}
	return invoices, rows.Err()
}

func (r *userRepository) Login(ctx context.Context, email, password string) (*entity.User, error) {
	query := `
		SELECT UserId, Email, Type FROM Users
			WHERE Email = ? AND Password = ?
	`

	var user entity.User
	err := r.db.QueryRow(query, email, password).Scan(
		&user.UserId,
		&user.Email,
		&user.UserType,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("The email or password is incorrect.")
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) UpdateMember(ctx context.Context, id int, user *entity.UserDetail) error {

	return nil
}
