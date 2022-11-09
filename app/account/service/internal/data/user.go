package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"

	"github.com/mizumoto-cn/bookkeepingo/app/account/service/internal/biz"
)

// type UserRepo implements the `UserRepo` interface defined in `biz\account.go`
type userRepo struct {
	data  *Data
	log   *log.Helper
	table *gorm.DB
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data:  data,
		log:   log.NewHelper(logger),
		table: data.db.Table("user"),
	}
}

func (ur *userRepo) FetchByUserMail(ctx context.Context, userMail string) (*biz.User, error) {
	user := &biz.User{}
	ur.table.WithContext(ctx).First(user, "user_mail = ?", userMail)
	if user.ID == 0 {
		return nil, biz.ErrUserNotFound
	}
	return user, nil
}

func (ur *userRepo) Save(ctx context.Context, user *biz.User) (int64, error) {
	result := ur.table.WithContext(ctx).Create(user)
	if result.Error != nil {
		return 0, result.Error
	}
	return user.ID, nil
}
