package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/cidstein/super-brunfo/internal/model"
	"github.com/cidstein/super-brunfo/internal/repository"
)

type SignUpUseCase struct {
	UserRepository repository.UserRepositoryInterface
}

func (s *SignUpUseCase) SignUp(ctx context.Context, db *pgx.Conn, email, password, nickname string) error {
	s.UserRepository = repository.NewUserRepository(db)

	err := s.UserRepository.Save(ctx, model.User{
		ID:       uuid.New().String(),
		Email:    email,
		Password: password,
		Nickname: nickname,
	})
	if err != nil {
		return err
	}

	return nil
}
