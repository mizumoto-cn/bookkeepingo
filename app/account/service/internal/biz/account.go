// Usecase `Account` relies on account-data repository AccountRepo to operate on account data.
// Also `password` shall be encrypted in a separated encryption service.
// All these dependencies are injected into `Account` usecase through wire.

package biz

import (
	"context"

	"github.com/mizumoto-cn/bookkeepingo/app/account/service/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
)

type Account struct {
	ID          int64
	AccountMail string
	PassWord    string
	NickName    string
	Phone       string
}

// Repo interface for Dependency Reverse
// `biz` requires `data` to implement the `repo` interface
type AccountRepo interface {
	// FetchByAccountMail fetch account by account mail address
	FetchByAccountMail(ctx context.Context, accountMail string) (*Account, error)
	// FetchByAccountID fetch account by account id
	FetchByID(ctx context.Context, id int64) (*Account, error)
	// VerifyPassword verify password
	VerifyPassword(ctx context.Context, acc *Account) error
	// Save saves account info and returns the ID of the saved account
	Save(ctx context.Context, account *Account) (int64, error)
}

type AccountUsecase struct {
	// authConfig  *conf.Auth
	accountRepo AccountRepo
	logger      *log.Helper
}

// var (
// 	ErrInvalidParam         = errors.New("invalid param")
// 	ErrAccountNotFound      = errors.New("account not found")
// 	ErrAccountAlreadyExists = errors.New("account already exists")
// 	ErrInvalidPassword      = errors.New("invalid password")
// 	ErrInvalidAccountMail   = errors.New("invalid account mail")
// )

// NewAccountUsecase returns a new account usecase.
// Dependencies are passed in as parameters.
func NewAccountUsecase(authConfig *conf.Auth, accountRepo AccountRepo, logger log.Logger) *AccountUsecase {
	return &AccountUsecase{
		// authConfig:  authConfig,
		accountRepo: accountRepo,
		logger:      log.NewHelper(log.With(logger, "module", "account-service/usecase/account")),
	}
}
