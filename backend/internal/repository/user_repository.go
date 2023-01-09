package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/cidstein/super-brunfo/internal/model"
)

type UserRepositoryInterface interface {
	Save(ctx context.Context, user model.User) error
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	SignIn(ctx context.Context, email string) (*model.User, error)
}

type UserRepository struct {
	Db *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) *UserRepository {
	return &UserRepository{Db: db}
}

func (r *UserRepository) Save(ctx context.Context, user model.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		fmt.Printf("Error hashing password: %v", err)
		return err
	}

	_, err = r.Db.Exec(
		ctx,
		`INSERT INTO "user" (id, email, password, nickname) VALUES ($1, $2, $3, $4)`,
		user.ID,
		user.Email,
		string(hashedPassword),
		user.Nickname,
	)

	return err
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User

	err := r.Db.QueryRow(
		ctx,
		`SELECT id, email, password, nickname FROM "user" WHERE email = $1`,
		email,
	).Scan(&user.ID, &user.Email, &user.Password, &user.Nickname)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) SignIn(ctx context.Context, email string) (*model.User, error) {
	var user model.User

	err := r.Db.QueryRow(
		ctx,
		`
			SELECT
				id, email, password, nickname
			FROM
				"user"
			WHERE
				email = $1
		`,
		email,
	).Scan(&user.ID, &user.Email, &user.Password, &user.Nickname)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
