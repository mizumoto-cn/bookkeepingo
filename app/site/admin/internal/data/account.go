package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	accountV1 "github.com/mizumoto-cn/bookkeepingo/api/account/service/v1"
	"github.com/mizumoto-cn/bookkeepingo/app/site/admin/internal/biz"
)

var _ biz.AccountRepo = (*accountRepo)(nil)

type accountRepo struct {
	log  *log.Helper
	data *Data
}

func NewAccountRepo(data *Data, logger log.Logger) biz.AccountRepo {
	return &accountRepo{
		log:  log.NewHelper(log.With(logger, "module", "site-admin/data/account")),
		data: data,
	}
}

func (r *accountRepo) FetchAccountByID(ctx context.Context, id int64) (*biz.Account, error) {
	acc, err := r.data.client.FetchAccountByID(ctx, &accountV1.FetchAccountByIDRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return &biz.Account{
		Id: acc.Id,
	}, nil
}

func (r *accountRepo) ListAccounts(ctx context.Context, pageNum, pageSize int64) ([]*biz.Account, error) {
	resp, err := r.data.client.ListAccount(ctx, &accountV1.ListAccountRequest{})
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Account, 0)
	for _, acc := range resp.Accounts {
		rv = append(rv, &biz.Account{
			Id: acc.Id,
		})
	}
	return rv, nil
}
