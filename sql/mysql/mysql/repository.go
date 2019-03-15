package mysql

import (
	"fmt"

	"database/sql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

const (
	DbMq           = "z_mq"
	DbActivity     = "z_activity"
	TableConsumers = "consumers"
	TableAccount   = "account"
	TableTasks     = "tasks"
)

const (
	TaskKindGeneral = iota + 1
	TaskKindShare
)

type DatabaseAdminInfo struct {
	AdminUser string
	Pwd       string
	Addr      string
	DbName    string
}
type DbRepo struct {
	conn   *sqlx.DB
	logger log.Logger
	info   *DatabaseAdminInfo
}

type StringField struct {
	Uid       int64
	Vipexpiry string
	Others    string
	Props     string
	Gifts     string
	Friends   string
	Records   string
	Tags      string
	Medals    string
	Strings   string
}

func NewDbRepo(info *DatabaseAdminInfo, logger log.Logger) (*DbRepo, error) {
	var err error
	dbRepo := &DbRepo{
		logger: logger,
		info:   info,
	}
	dsn := fmt.Sprintf("%v:%v@tcp(%v:13306)/%v", info.AdminUser, info.Pwd, info.Addr, info.DbName)
	if dbRepo.conn, err = sqlx.Open("mysql", dsn); err != nil {
		return nil, err
	}
	if err = dbRepo.conn.Ping(); err != nil {
		return nil, err
	}
	return dbRepo, nil
}

func (r *DbRepo) CreateTableIfNotExist(db string, table string, template string) (err error) {
	var (
		res sql.Result
		aff int64
	)
	res, err = r.conn.Exec(fmt.Sprintf("create table IF NOT EXISTS %s.%s (%s)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;", db, table, template))
	if err != nil {
		level.Error(r.logger).Log("DbRepo", "createTableIfNotExist", "err", err.Error())
		return
	}
	aff, err = res.RowsAffected()
	if aff > 0 {
		sqlstmt := fmt.Sprintf("grant all privileges on  %s.%s to %s", db, table, r.info.AdminUser)
		_, err = r.conn.Exec(sqlstmt)
		level.Info(r.logger).Log("DbRepo", "createTableIfNotExist", "sql", sqlstmt)
		if err != nil {
			level.Error(r.logger).Log("DbRepo", "createTableIfNotExist", "err", err.Error())
			return
		}
	}
	return
}
