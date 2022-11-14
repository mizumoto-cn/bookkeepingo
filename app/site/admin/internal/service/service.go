package service

import (
	"github.com/google/wire"

	v1 "github.com/mizumoto-cn/bookkeepingo/api/site/admin/v1"
	"github.com/mizumoto-cn/bookkeepingo/app/site/admin/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewSiteAdminService)

type SiteAdminService struct {
	v1.UnimplementedSiteAdminServiceServer

	ac  *biz.AccountUsecase
	log *log.Helper
}

func NewSiteAdminService(ac *biz.AccountUsecase, logger log.Logger) *SiteAdminService {
	return &SiteAdminService{
		ac:  ac,
		log: log.NewHelper(log.With(logger, "module", "site-admin/service")),
	}
}
