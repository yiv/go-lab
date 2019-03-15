package mongo

import (
	"database/sql"
	"fmt"
	"github.com/go-kit/kit/log/level"
	"time"
	"context"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/readpref"

	"github.com/go-kit/kit/log"
)

const (
	DatabaseUser = "z_user"
	TableAccount = "account"
)

type MongoOptions struct {
	DBUser string
	DBPwd  string
	DBName string
	DBAddr string
	DBPort string
}

type MongoRepo struct {
	MongoOptions
	Logger log.Logger
}

func NewSWPRepo(options MongoOptions, Logger log.Logger) (*MongoRepo, error) {
	repo := &MongoRepo{
		Logger:       Logger,
		MongoOptions: options,
	}
	dsn := fmt.Sprintf("mongodb://%v:%v", repo.DBAddr, repo.DBPort)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, dsn)
	if err != nil {
		return nil, err
	}
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}
	return repo, nil
}

func (r *MongoRepo) FindUserById(uid string) (user *ucenter.User, err error) {
	user = &ucenter.User{}
	sqlsmt := fmt.Sprintf("select id, uid, api, app, unionid, nick, gender, avatar, status, mtime , ctime , online, ipaddr, gem, coin, score, win, lose from %s.%s where uid = '%v'", DatabaseUser, TableAccount, uid)
	err = r.conn.Get(user, sqlsmt)
	if err != nil {
		level.Error(r.logger).Log("DbRepo", "FindUserById", "err", err.Error(), "uid", uid)
		if err == sql.ErrNoRows {
			return nil, ucenter.ErrNotFound
		}
		return nil, err
	}
	return user, nil
}


