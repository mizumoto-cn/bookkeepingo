package biz

import (
	"context"
	"errors"

	v1 "github.com/mizumoto-cn/bookkeepingo/api/account/service/v1"

	jwt "github.com/golang-jwt/jwt/v4"
)

var (
	ErrLoginFailed = errors.New("login failed")
)

type AuthUsecase struct {
	auth    *conf.Auth
	accRepo AccountRepo
}

func NewAuthUsecase(authConfig *conf.Auth, accRepo AccountRepo) *AuthUsecase {
	return &AuthUsecase{
		auth:    authConfig,
		accRepo: accRepo,
	}
}

func (au *AuthUsecase) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {
	// get account info
	acc, err := au.accRepo.FetchByAccountMail(ctx, req.Mail)
	if err != nil {
		return nil, v1.ErrorLoginFailed("login failed, mail not found, %s", err.Error())
	}

	// check permission (password/blacklist ...etc)
	err = au.accRepo.VerifyPassword(ctx, acc)
	if err != nil {
		return nil, v1.ErrorLoginFailed("login failed, password not match, %s", err.Error())
	}

	// token-gen
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"account_id":   acc.ID,
		"account_mail": acc.AccountMail,
	})
	signedString, err := claims.SignedString([]byte(au.auth.JwtSecret))
	if err != nil {
		return nil, v1.ErrorLoginFailed("login failed, token gen failed, %s", err.Error())
	}
	return &v1.LoginResponse{
		Token: signedString,
	}, nil
}
