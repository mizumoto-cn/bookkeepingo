package service

import (
	"context"

	v1 "github.com/mizumoto-cn/bookkeepingo/api/account/service/v1"
	"github.com/mizumoto-cn/bookkeepingo/app/account/service/internal/biz"
)

func (as *AccountService) CreateAccount(ctx context.Context, req *v1.CreateAccountRequest) (*v1.CreateAccountResponse, error) {
	ac, err := as.auc.Create(ctx, &biz.Account{
		AccountMail: req.AccountMail,
		PassWord:    req.Password,
	})
	if err != nil {
		as.log.Errorf("[CreateAccount] failed to create account, err: %v", err)
		return nil, err
	}
	return &v1.CreateAccountResponse{
		Id:          ac.ID,
		AccountMail: ac.AccountMail,
	}, nil
}

func (as *AccountService) FetchByAccountMail(ctx context.Context, req *v1.FetchAccountByAccountMailRequest) (*v1.FetchAccountByAccountMailResponse, error) {
	return as.auc.FetchByAccountMail(ctx, req)
}

func (as *AccountService) FetchByID(ctx context.Context, req *v1.FetchAccountByIDRequest) (*v1.FetchAccountByIDResponse, error) {
	ac, err := as.auc.FetchByID(ctx, req.Id)
	if err != nil {
		as.log.Errorf("[FetchByID] failed to fetch account, err: %v", err)
		return nil, err
	}
	return &v1.FetchAccountByIDResponse{
		Id:          ac.ID,
		AccountMail: ac.AccountMail,
	}, nil
}

func (as *AccountService) VerifyPassword(ctx context.Context, req *v1.VerifyPasswordRequest) (*v1.VerifyPasswordResponse, error) {
	ok, err := as.auc.VerifyPassword(ctx, &biz.Account{AccountMail: req.AccountMail, PassWord: req.Password})
	if err != nil {
		as.log.Errorf("[VerifyPassword] failed, err: %v", err)
		return nil, err
	}

	return &v1.VerifyPasswordResponse{
		Ok: ok,
	}, nil
}

func (as *AccountService) Save(ctx context.Context, req *v1.SaveAccountRequest) (*v1.SaveAccountResponse, error) {
	return as.auc.Save(ctx, req)
}
