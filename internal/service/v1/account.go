package v1

import (
	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/mizumoto-cn/bookkeepingo/api/account/v1"
	"github.com/mizumoto-cn/bookkeepingo/internal/biz"
)

type accountService struct {
	v1.UnimplementedAccountServer
	log *log.Helper
	auc *biz.AccountUseCase
}

func NewAccountService(auc *biz.AccountUseCase, logger log.Logger) v1.AccountServer {
	return &accountService{
		log: log.NewHelper(logger),
		auc: auc,
	}
}
