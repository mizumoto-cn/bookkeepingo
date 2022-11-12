package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "github.com/mizumoto-cn/bookkeepingo/api/account/service/v1"
	"github.com/mizumoto-cn/bookkeepingo/app/account/service/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewAccountService)

type AccountService struct {
	v1.UnimplementedAccountServiceServer

	auc *biz.AccountUsecase
	log *log.Helper
}

func NewAccountService(auc *biz.AccountUsecase, logger log.Logger) *AccountService {
	return &AccountService{
		auc: auc,
		log: log.NewHelper(logger),
	}
}
