package main

import (
	"github.com/go-kit/kit/log/level"
	"github.com/go-kit/kit/log"
	"github.com/yiv/go-lab/sql/mysql/mysql"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var logger log.Logger

func main() {
	logger = log.NewLogfmtLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	dbInfo := &mysql.DatabaseAdminInfo{AdminUser: "root", Pwd: "root", Addr: "192.168.1.12", DbName: mysql.DbMq}
	dbRepo, err  := mysql.NewDbRepo(dbInfo, logger)
	if err != nil {
		level.Error(logger).Log("err", err.Error())
		return
	}
	if err = dbRepo.CreateTableIfNotExist(mysql.DbMq, mysql.TableConsumers, mysql.TableConsumersTemplate); err != nil {
		level.Error(logger).Log("err", err.Error())
		return
	}

}
