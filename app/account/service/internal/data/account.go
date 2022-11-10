package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"

	"github.com/mizumoto-cn/bookkeepingo/app/account/service/internal/biz"
)

// type AccountRepo implements the `AccountRepo` interface defined in `biz\account.go`
type accountRepo struct {
	data  *Data
	log   *log.Helper
	table *gorm.DB
}

func NewAccountRepo(data *Data, logger log.Logger) biz.AccountRepo {
	return &accountRepo{
		data:  data,
		log:   log.NewHelper(logger),
		table: data.db.Table("account"),
	}
}

func (ur *accountRepo) FetchByAccountMail(ctx context.Context, accountMail string) (*biz.Account, error) {
	account := &biz.Account{}
	ur.table.WithContext(ctx).First(account, "account_mail = ?", accountMail)
	if account.ID == 0 {
		return nil, biz.ErrAccountNotFound
	}
	return account, nil
}

func (ur *accountRepo) Save(ctx context.Context, account *biz.Account) (int64, error) {
	result := ur.table.WithContext(ctx).Create(account)
	if result.Error != nil {
		return 0, result.Error
	}
	return account.ID, nil
}
