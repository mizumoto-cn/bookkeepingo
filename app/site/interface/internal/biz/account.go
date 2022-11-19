package biz

import (
	"context"
	"errors"

	"github.com/go-kratos/kratos/v2/log"
)

var (
	ErrPasswordInvalid     = errors.New("password invalid")
	ErrAccountNotFound     = errors.New("account not found")
	ErrAccountMailNotValid = errors.New("account mail not valid")
)

type Account struct {
	Id          int64
	AccountMail string
	Password    string
}

func NewAccount(mail string, password string) (Account, error) {
	if len(password) < 6 {
		return Account{}, ErrPasswordInvalid
	}
	if len(mail) < 1 {
		return Account{}, ErrAccountMailNotValid
	}
	return Account{
		AccountMail: mail,
		Password:    password,
	}, nil
}

type AccountRepo interface {
	FetchById(ctx context.Context, id int64) (*Account, error)
	FetchByMail(ctx context.Context, mail string) (*Account, error)
	Save(ctx context.Context, account *Account) error

	VerifyPassword(ctx context.Context, account *Account, password string) error
}

type AccountUsecase struct {
	repo        AccountRepo
	log         *log.Helper
	authUsecase *AuthUsecase
}

func NewAccountUsecase(repo AccountRepo, logger log.Logger, authUsecase *AuthUsecase) *AccountUsecase {
	log := log.NewHelper(log.With(logger, "module", "site-interface/usecase/account"))
	return &AccountUsecase{
		repo:        repo,
		log:         log,
		authUsecase: authUsecase,
	}
}

func (auc *AccountUsecase) Logout(ctx context.Context, a *Account) error {
	return nil
	// TODO: implement
}
