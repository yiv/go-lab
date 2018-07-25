package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/go-kit/kit/log/level"
	"github.com/go-kit/kit/log"
	"os"
)

const (
	DatabaseAnalytics = "tp_analytics"
	TableConcurrency  = "concurrency_"
)
const ConcurrencyTableTemplate = `
	id serial NOT NULL,
	game int NOT NULL DEFAULT 0,
	room int NOT NULL DEFAULT 0,
	desk int NOT NULL DEFAULT 0,
	ctime int NOT NULL DEFAULT 0,
	count int NOT NULL DEFAULT 0,
	INDEX (game,ctime),
	PRIMARY KEY (id)
`

type DbRepo struct {
	conn   *sqlx.DB
	logger log.Logger
}
type Schema struct {
	Table_schema string
	Table_name   string
}

func main() {
	dsn := "postgresql://edwin@103.207.164.43:26256/tp_mall?sslmode=disable"
	repo, err := NewDbUserRepo(dsn)
	if err != nil {
		level.Error(repo.logger).Log("err", err.Error())
		return
	}
	table := TableConcurrency + time.Now().Format("200601")
	if err = repo.createTableIfNotExist(DatabaseAnalytics, table, ConcurrencyTableTemplate); err != nil {
		level.Error(repo.logger).Log("err", err.Error())
		return
	}

}

func NewDbUserRepo(userDataSourceName string) (*DbRepo, error) {
	conn, err := sqlx.Open("postgres", userDataSourceName)
	if err != nil {
		return nil, err
	}
	dbRepo := &DbRepo{
		conn:   conn,
		logger: log.NewLogfmtLogger(os.Stdout),
	}
	return dbRepo, nil
}

func (r *DbRepo) isTableExist(db string, table string) (exist bool, err error) {
	if err = r.setDatabase(db); err != nil {
		return
	}
	tb := Schema{}
	//sqlsmt := fmt.Sprintf("select table_schema,table_name from information_schema.tables where information_schema.tables.table_schema = '%s' and information_schema.tables.table_name = '%s'", db, table)
	sqlsmt := fmt.Sprintf("select table_schema,table_name from information_schema.tables where information_schema.tables.table_schema = 'tp_analytics' and information_schema.tables.table_name = 'concurrency_201803'")
	err = r.conn.Get(&tb, sqlsmt)
	if err != nil {
		level.Error(r.logger).Log("DbRepo", "isTableExist", "sqlsmt", sqlsmt, "err", err.Error())
		if err == sql.ErrNoRows {
			return false, nil
		}
		return
	}
	level.Debug(r.logger).Log("DbRepo", "isTableExist", "tb", fmt.Sprintf("%#v", tb))
	return true, nil
}
func (r *DbRepo) createTableIfNotExist(db string, table string, template string) (err error) {
	exist, e := r.isTableExist(db, table)
	if e != nil {
		return e
	}
	if exist {
		return nil
	}
	if err = r.setDatabase(db); err != nil {
		return
	}

	if !exist {
		_, err = r.conn.Exec(fmt.Sprintf("create table %s.%s (%s)", db, table, template))
		if err != nil {
			level.Error(r.logger).Log("err", err.Error())
			return
		}
		_, err = r.conn.Exec(fmt.Sprintf("grant all on  %s.%s to edwin", db, table))
		if err != nil {
			level.Error(r.logger).Log("err", err.Error())
			return
		}
		_, err = r.conn.Exec(fmt.Sprintf("insert into  %s.%s (server, mtime, count) values (%d, %d, %d)", db, table, 8, 55, 10))
		if err != nil {
			level.Error(r.logger).Log("err", err.Error())
			return
		}
	}
	return
}
func (r *DbRepo) setDatabase(db string) (err error) {
	_, err = r.conn.Exec(fmt.Sprintf("set database = '%s'", db))
	if err != nil {
		level.Error(r.logger).Log("err", err.Error())
		return
	}
	return
}
