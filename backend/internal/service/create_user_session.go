package service

import (
	"fmt"
	"time"

	"github.com/cidstein/super-brunfo/internal/model"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
)

func CreateUserSession(redisClient *redis.Client, user model.User) error {
	sessionID := uuid.New().String()

	sessionUser := model.UserSession{
		ID:            sessionID,
		UserID:        user.ID,
		Authenticated: true,
	}

	err := redisClient.Set(sessionID, sessionUser, time.Hour*2)
	if err != nil {
		fmt.Printf("Error setting session: %v", err)
		fErr := fmt.Errorf("error setting session: %v", err)
		return fErr
	}

	return nil
}
