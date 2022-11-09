package service

import (
	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/mizumoto-cn/bookkeepingo/api/account/service/v1"
	"github.com/mizumoto-cn/bookkeepingo/app/account/service/internal/biz"
)

type accountService struct {
	v1.UnimplementedAccountServer
	log *log.Helper
	auc *biz.AccountUsecase
}

func NewAccountService(auc *biz.AccountUsecase, logger log.Logger) v1.AccountServer {
	return &accountService{
		log: log.NewHelper(logger),
		auc: auc,
	}
}
