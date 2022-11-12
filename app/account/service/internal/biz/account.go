// Usecase `Account` relies on account-data repository AccountRepo to operate on account data.
// Also `password` shall be encrypted in a separated encryption service.
// All these dependencies are injected into `Account` usecase through wire.

package biz

import (
	"context"
	"errors"
	"math/rand"

	v1 "github.com/mizumoto-cn/bookkeepingo/api/account/service/v1"

	"github.com/go-kratos/kratos/v2/log"
)

var (
	// 	ErrInvalidParam         = errors.New("invalid param")
	ErrAccountNotFound = errors.New("account not found")

	// ErrAccountAlreadyExists = errors.New("account already exists")
	// ErrInvalidPassword      = errors.New("invalid password")
	// ErrInvalidAccountMail   = errors.New("invalid account mail")
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

	// CreateAccount saves account info and returns the ID of the saved account
	CreateAccount(ctx context.Context, account *Account) (*Account, error)

	// FetchByAccountMail fetch account by account mail address
	FetchByAccountMail(ctx context.Context, accountMail string) (*Account, error)

	// FetchByAccountID fetch account by account id
	FetchByID(ctx context.Context, id int64) (*Account, error)

	// VerifyPassword verify password
	VerifyPassword(ctx context.Context, acc *Account) (bool, error)
}

type AccountUsecase struct {
	accountRepo AccountRepo
	logger      *log.Helper
}

// NewAccountUsecase returns a new account usecase.
// Dependencies are passed in as parameters.
func NewAccountUsecase(accountRepo AccountRepo, logger log.Logger) *AccountUsecase {
	return &AccountUsecase{
		accountRepo: accountRepo,
		logger:      log.NewHelper(log.With(logger, "module", "account-service/usecase/account")),
	}
}

// FetchByAccountMail fetch account by account mail address
func (au *AccountUsecase) FetchByAccountMail(ctx context.Context, req *v1.FetchAccountByAccountMailRequest) (*v1.FetchAccountByAccountMailResponse, error) {
	acc, err := au.accountRepo.FetchByAccountMail(ctx, req.AccountMail)
	if err != nil {
		// TODO: handle error
		return nil, err
	}
	return &v1.FetchAccountByAccountMailResponse{
		Id:          acc.ID,
		AccountMail: acc.AccountMail,
	}, nil
}

// FetchByID fetch account by account id
func (au *AccountUsecase) FetchByID(ctx context.Context, id int64) (*Account, error) {
	return au.accountRepo.FetchByID(ctx, id)
}

// Create
func (au *AccountUsecase) Create(ctx context.Context, acc *Account) (*Account, error) {
	ret, err := au.accountRepo.CreateAccount(ctx, acc)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// Save
func (au *AccountUsecase) Save(ctx context.Context, req *v1.SaveAccountRequest) (*v1.SaveAccountResponse, error) {
	acc := &Account{
		// TODO: need review
		ID: rand.Int63(),
		// ID:          req.Id,
		AccountMail: req.AccountMail,
		PassWord:    req.Password,
	}
	_, err := au.accountRepo.CreateAccount(ctx, acc)
	if err != nil {
		// TODO: handle error
		return nil, err
	}
	return &v1.SaveAccountResponse{
		Id: acc.ID,
	}, nil
}

// VerifyPassword verify password
func (au *AccountUsecase) VerifyPassword(ctx context.Context, acc *Account) (bool, error) {
	return au.accountRepo.VerifyPassword(ctx, acc)
}
