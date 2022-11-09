package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/mizumoto-cn/bookkeepingo/api/account/service/v1"
	"github.com/mizumoto-cn/bookkeepingo/app/account/service/internal/biz"
)

type accountService struct {
	v1.UnimplementedAccountServiceServer

	log *log.Helper
	auc *biz.AccountUsecase
}

func NewAccountService(auc *biz.AccountUsecase, logger log.Logger) v1.AccountServiceServer {
	return &accountService{
		log: log.NewHelper(logger),
		auc: auc,
	}
}

func (as *accountService) Register(ctx context.Context, req *v1.RegisterRequest) (*v1.RegisterResponse, error) {
	err := as.auc.Register(ctx, req.GetMail(), req.GetPassword())
	if err != nil {
		return nil, errors.New(500, "register failed", err.Error())
	}
	return &v1.RegisterResponse{}, nil
}
