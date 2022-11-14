package service

import (
	"context"

	v1 "github.com/mizumoto-cn/bookkeepingo/api/site/admin/v1"
)

func (s *SiteAdminService) Login(ctx context.Context, req *v1.LoginRequire) (*v1.LoginResponse, error) {
	// TODO: implement the business logic of Login
	return nil, nil
}

func (s *SiteAdminService) Logout(ctx context.Context, req *v1.LogoutRequire) (*v1.LogoutResponse, error) {
	return nil, nil
}
