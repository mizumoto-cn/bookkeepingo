// Usecase `Account` relies on user-data repository UserRepo to operate on user data.
// Also `password` shall be encrypted in a separated encryption service.
// All these dependencies are injected into `Account` usecase through wire.

package biz

import (
	"bookkeepingo/internal/conf"
	"context"
	"errors"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	ID        int64
	UserMail  string
	PassWord  string
	NickName  string
	AvatarUrl string
}

// Repo interface for Dependency Reverse
// `biz` requires `data` to implement the `repo` interface
type UserRepo interface {
	// FetchByUserMail fetch user by user mail address
	// return nil, ErrUserNotFound if not found
	FetchByUserMail(ctx context.Context, userMail string) (*user, error)

	// Save saves user info and returns the ID of the saved user
	Save(ctx context.Context, user *user) (int64, error)
}

type AccountUsecase struct {
	authConfig        *conf.Auth
	encryptionService EncryptionService // TODO: encryption service
	userRepo          UserRepo
	logger            *log.Helper
}

// NewAccountUsecase returns a new account usecase.
// Dependencies are passed in as parameters.
func NewAccountUsecase(authConfig *conf.Auth, encryptionService EncryptionService, userRepo UserRepo, logger log.Logger) *AccountUsecase {
	return &AccountUsecase{
		authConfig:        authConfig,
		encryptionService: encryptionService,
		userRepo:          userRepo,
		logger:            log.NewHelper(logger),
	}
}

// Registration
func (uc *AccountUsecase) Register(ctx context.Context, userMail, password string) error {
	if userMail == "" || password == "" {
		return fmt.Errorf("invalid user mail or password: %w", ErrInvalidParam)
	}
	// check if user mail exists
	user, err := uc.userRepo.FetchByUserMail(ctx, userMail)
	if err != nil && !errors.Is(err, ErrUserNotFound) {
		log.Errorf("Registration failed,[userMail: %s, pwd: %s], err: %v", userMail, password, err)
		return fmt.Errorf("Registration failed: %w", err)
	}
	if user != nil {
		return fmt.Errorf("user mail already exists: %w", ErrUserAlreadyExists)
	}
	// encrypt password
	encryptedPassword, err := uc.encryptionService.Encrypt(ctx, []byte(password))
	if err != nil {
		log.Errorf("Registration failed,[userMail: %s, pwd: %s], err: %v", userMail, password, err)
		return fmt.Errorf("Registration failed: %w", err)
	}
	_, err = uc.userRepo.Save(ctx, &user{
		UserMail: userMail,
		PassWord: string(encryptedPassword),
	})
	if err != nil {
		log.Errorf("Registration failed,[userMail: %s, pwd: %s], err: %v", userMail, password, err)
		return fmt.Errorf("Registration failed: %w", err)
	}
	return nil
}

// TODO: Login
