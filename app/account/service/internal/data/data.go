// data implements the `repo` interface defined in `biz`

package data

import (
	"context"

	"github.com/go-redis/redis/v8"

	"github.com/mizumoto-cn/bookkeepingo/app/account/service/internal/conf"
	"github.com/mizumoto-cn/bookkeepingo/app/account/service/internal/data/ent"
	"github.com/mizumoto-cn/bookkeepingo/app/account/service/internal/data/ent/migrate"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewEntClient, NewRedisCmd, NewAccountRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	db  *ent.Client
	rdb redis.Cmdable
}

// EntClient .
func NewEntClient(conf *conf.Data, logger log.Logger) *ent.Client {
	log := log.NewHelper(log.With(logger, "module", "account-service/data/ent"))

	client, err := ent.Open(
		conf.Database.Driver,
		conf.Database.Source,
	)
	if err != nil {
		log.Fatalf("failed opening connection to db: %v", err)
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background(), migrate.WithForeignKeys(false)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}

// RedisCmd .
func NewRedisCmd(conf *conf.Data, logger log.Logger) redis.Cmdable {
	log := log.NewHelper(log.With(nil, "module", "account-service/data/redis"))
	rdb := redis.NewClient(&redis.Options{
		Addr:         conf.Redis.Addr,
		Password:     conf.Redis.Password,
		DB:           int(conf.Redis.Db),
		DialTimeout:  conf.Redis.DialTimeout.AsDuration(),
		ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
		PoolSize:     10,
	})
	timeout, cancelFunc := context.WithTimeout(context.Background(), conf.Redis.DialTimeout.AsDuration())
	defer cancelFunc()
	_, err := rdb.Ping(timeout).Result()
	if err != nil {
		log.Fatalf("failed connecting to redis: %v", err)
	}
	return rdb
}

// NewData .
func NewData(logger log.Logger, db *ent.Client, rdb redis.Cmdable) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "account-service/data"))
	d := &Data{
		db:  db,
		rdb: rdb,
	}
	return d, func() {
		log.Info("closing the data resources")
		if err := d.db.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}
