package service

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/cidstein/super-brunfo/internal/repository"
)

type SignInUseCase struct {
	UserRepository repository.UserRepositoryInterface
}

func (s *SignInUseCase) SignIn(ctx context.Context, db *pgx.Conn, email, password string) error {
	s.UserRepository = repository.NewUserRepository(db)

	user, err := s.UserRepository.SignIn(ctx, email)
	if err != nil {
		return err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return fmt.Errorf("invalid password")
	}

	return nil
}
