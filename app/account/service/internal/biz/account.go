// Usecase `Account` relies on user-data repository UserRepo to operate on user data.
// Also `password` shall be encrypted in a separated encryption service.
// All these dependencies are injected into `Account` usecase through wire.

package biz

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/mizumoto-cn/bookkeepingo/app/account/service/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt/v4"
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
	FetchByUserMail(ctx context.Context, userMail string) (*User, error)

	// Save saves user info and returns the ID of the saved user
	Save(ctx context.Context, user *User) (int64, error)
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

// AccountUsecase.Login returns a token or error if login failed.
func (uc *AccountUsecase) Login(ctx context.Context, userMail, password string) (string, error) {
	if userMail == "" || password == "" {
		return "", fmt.Errorf("invalid user mail or password: %w", ErrInvalidParam)
	}
	// check if user mail exists
	user, err := uc.userRepo.FetchByUserMail(ctx, userMail)
	if err != nil {
		log.Errorf("Login failed,[userMail: %s], err: %v", userMail, err)
		return "", fmt.Errorf("Login failed: %w", err)
	}
	// check password
	err = uc.encryptionService.Compare(ctx, []byte(password), []byte(user.PassWord))
	if err != nil {
		log.Errorf("Login failed,[userMail: %s], err: %v", userMail, err)
		return "", fmt.Errorf("Login failed: %w", err)
	}
	// generate jwt token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(uc.authConfig.GetExpireDuration())),
	})
	token, err := claims.SignedString([]byte(uc.authConfig.GetJwtSecret()))
	if err != nil {
		log.Errorf("Login token creation failed,[userMail: %s], err: %v", userMail, err)
		return "", fmt.Errorf("Login failed: %w", err)
	}
	return token, nil
}
