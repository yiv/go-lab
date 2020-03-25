package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

const (
	//MongoAddr = "192.168.1.12"
	//MongoPort = "27017"
	AccountId = "8876543"

	//DBAddr = "10.72.17.30"
	//DBPort = "13306"
	//DBUser = "root"
	//DBPwd  = "123@mysql"

	//DBAddr             = "119.29.1.133"
	//DBPort             = "30257"

	DBAddr     = "10.72.17.30"
	DBPort     = "26257"
	DBUserName = "root"
	DBPwd      = "root"
	DBIPUser   = "ip_user"
)

func main() {
	ops := CRDBOptions{DBAddr: DBAddr, DBPort: DBPort, DBUser: DBUserName, DBPwd: DBPwd, DBName: dbName}
	dbRepo, err := NewCRDBRepo(ops)
	if err != nil {
		log.Println("main", "mongoRepo", "err", err.Error())
		os.Exit(1)
	}

}

func SMSSend(result *Result) error {

}

type Result struct {
	Id       int64  `xorm:"id"    json:"id"`
	NickName string `xorm:"nick_name"    json:"nick_name"`
	Mobile   string `xorm:"mobile"    json:"mobile"`
	Error    string `xorm:"error"    json:"error"`
	Sent     bool   `xorm:"sent"    json:"sent"`
}

type CRDBOptions struct {
	DBUser string
	DBPwd  string
	DBName string
	DBAddr string
	DBPort string
}

type CRDBRepo struct {
	CRDBOptions
	engine *xorm.Engine
}

func NewCRDBRepo(ops CRDBOptions) (*CRDBRepo, error) {
	repo := &CRDBRepo{
		CRDBOptions: ops,
	}
	// user:password@tcp(localhost:5555)/dbname
	dsn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", repo.DBUser, repo.DBPwd, repo.DBAddr, repo.DBPort, repo.DBName)
	engine, err := xorm.NewEngine("postgres", dsn)
	if err != nil {
		log.Fatal("39", err.Error())
		return nil, err
	}
	engine.SetMaxOpenConns(200)
	repo.engine = engine
	if err = engine.Ping(); err != nil {
		log.Fatal("45", err.Error())
		return nil, err
	}
	return repo, nil
}

func (r *CRDBRepo) ReadPage(pageNum, pageSize int32) (results []*Result, err error) {
	results = []*Result{}
	if pageSize <= 0 {
		pageSize = 100
	}
	if pageNum <= 0 {
		pageNum = 1
	}
	pageStart := (pageNum - 1) * pageSize
	order := "id desc"
	if err = r.engine.Where("online_status = 1").OrderBy(order).Limit(int(pageSize), int(pageStart)).Find(&results); err != nil {
		log.Println("pageNum", pageNum, "pageSize", pageSize, "pageStart", pageStart, "order", order, "err", err.Error())
		return nil, err
	}
	return
}

func (r *CRDBRepo) IsTableExist(tableName string) (yes bool, err error) {
	if yes, err = r.engine.IsTableExist(tableName); err != nil {
		log.Println("err", err.Error())
	}
	return
}
