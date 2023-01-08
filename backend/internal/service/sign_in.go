package service

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/cidstein/super-brunfo/internal/repository"
)

type SignInUseCase struct {
	UserRepository repository.UserRepositoryInterface
}

func (s *SignInUseCase) SignIn(ctx context.Context, db *pgx.Conn, username, password string) error {
	s.UserRepository = repository.NewUserRepository(db)

	user, err := s.UserRepository.FindByEmail(ctx, username)
	if err != nil {
		return err
	}

	if !user.ComparePassword(password) {
		return fmt.Errorf("invalid password")
	}

	return nil
}
