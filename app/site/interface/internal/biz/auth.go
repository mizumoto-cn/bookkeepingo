package biz

import (
	"context"
	"errors"

	v1 "github.com/mizumoto-cn/bookkeepingo/api/site/interface/v1"
	"github.com/mizumoto-cn/bookkeepingo/app/site/interface/internal/conf"

	"github.com/golang-jwt/jwt"
)

var (
	ErrLoginFailed = errors.New("login failed")
)

type AuthUsecase struct {
	key         string
	accountRepo AccountRepo
}

func NewAuthUsecase(conf *conf.Auth, accountRepo AccountRepo) *AuthUsecase {
	return &AuthUsecase{
		key:         conf.ApiKey,
		accountRepo: accountRepo,
	}
}

func (receiver *AuthUsecase) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {

	// get account
	account, err := receiver.accountRepo.FetchByMail(ctx, req.AccountMail)
	if err != nil {
		return nil, v1.ErrorLoginFailed("account not found: %s", err.Error())
	}
	// check permission(password blacklist etc...)
	err = receiver.accountRepo.VerifyPassword(ctx, account, req.Password)
	if err != nil {
		return nil, v1.ErrorLoginFailed("password not match")
	}
	// generate token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"account_id": account.Id,
	})
	signedString, err := claims.SignedString([]byte(receiver.key))
	if err != nil {
		return nil, v1.ErrorLoginFailed("generate token failed: %s", err.Error())
	}
	return &v1.LoginResponse{
		Token: signedString,
	}, nil
}

func (receiver *AuthUsecase) Register(ctx context.Context, req *v1.RegisterRequest) (*v1.RegisterResponse, error) {

	// check account mail
	_, err := receiver.accountRepo.FetchByMail(ctx, req.AccountMail)
	if !errors.Is(err, ErrAccountNotFound) {
		return nil, v1.ErrorRegisterFailed("account mail already exists")
	}
	// create account
	account, err := NewAccount(req.AccountMail, req.Password)
	if err != nil {
		return nil, v1.ErrorRegisterFailed("create account failed: %s", err.Error())
	}
	// save account
	err = receiver.accountRepo.Save(ctx, &account)
	if err != nil {
		return nil, v1.ErrorRegisterFailed("save account failed: %s", err.Error())
	}
	return &v1.RegisterResponse{
		Id: account.Id,
	}, nil
}
