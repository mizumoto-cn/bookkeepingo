// data implements the `repo` interface defined in `biz`

package data

import (
	"github.com/mizumoto-cn/bookkeepingo/app/account/service/internal/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewAccountRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	db  *gorm.DB
	rdb *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	db, err := gorm.Open(mysql.Open(c.GetDatabase().GetSource()), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.GetRedis().GetAddr(),
		Password:     c.GetRedis().GetPassword(),
		DB:           int(c.GetRedis().GetDb()),
		DialTimeout:  c.GetRedis().GetDialTimeout().AsDuration(),
		ReadTimeout:  c.GetRedis().GetReadTimeout().AsDuration(),
		WriteTimeout: c.GetRedis().GetWriteTimeout().AsDuration(),
	})
	rdb.AddHook(redisotel.TracingHook{})
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{db: db, rdb: rdb}, cleanup, nil
}
