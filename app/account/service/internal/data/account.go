package data

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/mizumoto-cn/bookkeepingo/app/account/service/internal/biz"
	"github.com/mizumoto-cn/bookkeepingo/app/account/service/internal/data/ent"
	"github.com/mizumoto-cn/bookkeepingo/app/account/service/internal/pkg/util"
)

var _ biz.AccountRepo = (*accountRepo)(nil)

var accountCacheKey = func(mail string) string {
	return "account_cache_key_" + mail
}

// type AccountRepo implements the `AccountRepo` interface defined in `biz\account.go`
type accountRepo struct {
	data *Data
	log  *log.Helper
}

func NewAccountRepo(data *Data, logger log.Logger) biz.AccountRepo {
	return &accountRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "account-service/data/account")),
	}
}

func (ur *accountRepo) CreateAccount(ctx context.Context, account *biz.Account) (*biz.Account, error) {
	hashedPassword, err := util.HashPassword(account.PassWord)
	if err != nil {
		return nil, err
	}
	acc, err := ur.data.db.Account.Create().
		SetAccountmail(account.AccountMail).
		SetPasswordHash(hashedPassword).
		SetNickname(account.NickName).
		SetPhone(account.Phone).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Account{
		ID:          acc.ID,
		AccountMail: acc.Accountmail,
	}, nil
}

func (ur *accountRepo) FetchByAccountMail(ctx context.Context, accountMail string) (*biz.Account, error) {
	return nil, nil
}

func (ur *accountRepo) FetchByID(ctx context.Context, id int64) (*biz.Account, error) {
	// try to get account from cache
	cacheKey := accountCacheKey(fmt.Sprintf("%d", id))
	acc, err := ur.getUserFromCache(ctx, cacheKey)
	if err != nil {
		// if not found, get account from db
		acc, err = ur.data.db.Account.Get(ctx, id)
		if err != nil {
			return nil, biz.ErrAccountNotFound
		}
		// save account to cache
		ur.saveAccountToCache(ctx, cacheKey, acc)
	}
	return &biz.Account{
		ID:          acc.ID,
		AccountMail: acc.Accountmail,
	}, nil
}

func (ur *accountRepo) VerifyPassword(ctx context.Context, acc *biz.Account) (bool, error) {
	return false, nil
}

func (ur *accountRepo) getUserFromCache(ctx context.Context, cacheKey string) (*ent.Account, error) {
	result, err := ur.data.rdb.Get(ctx, cacheKey).Result()
	if err != nil {
		return nil, err
	}
	cacheAcc := &ent.Account{}
	err = json.Unmarshal([]byte(result), cacheAcc)
	if err != nil {
		return nil, err
	}
	return cacheAcc, nil
}

func (ur *accountRepo) saveAccountToCache(ctx context.Context, cacheKey string, acc *ent.Account) {
	marshal, err := json.Marshal(acc)
	if err != nil {
		ur.log.Errorf("save account to cache failed: json.Marshal(%v) failed: %v", acc, err)
		return
	}
	err = ur.data.rdb.Set(ctx, cacheKey, marshal, time.Minute*30).Err()
	if err != nil {
		ur.log.Errorf("save account to cache failed: rdb.Set(%s, %s) failed: %v", cacheKey, marshal, err)
		return
	}
}
