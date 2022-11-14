package biz

import (
	accountV1 "github.com/mizumoto-cn/bookkeepingo/api/account/service/v1"

	log "github.com/go-kratos/kratos/v2/log"
)

type Account struct {
	Id int64
}

type AccountRepo interface {
}

type AccountUsecase struct {
	repo   AccountRepo
	client accountV1.AccountServiceClient
	log    *log.Helper
}

func NewAccountUsecase(repo AccountRepo, client accountV1.AccountServiceClient, logger log.Logger) *AccountUsecase {
	log := log.NewHelper(log.With(logger, "module", "site-admin/usecase/account"))
	return &AccountUsecase{
		repo:   repo,
		client: client,
		log:    log,
	}
}
