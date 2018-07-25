package main

import (
	"database/sql"
	"fmt"
	//"github.com/jmoiron/sqlx"
	"git.ifunbow.com/tpserver/usercenter/pkg/ucenter"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

//type Account struct {
//	Id          int64
//	Unionid     string
//	Uuid        string
//	Username    string
//	Password    string
//	Nick        string
//	Gender      bool
//	Addr        string
//	Avatar      string
//	Isguest     bool
//	Condays     int
//	Signdate    string
//	Vipsigndate string
//	Status      bool
//	Mtime       string
//	Ctime       string
//	Token       string
//	Bankpwd     string
//	Forbid      string
//	Imsi        string
//	Imei        string
//	Mac         string
//	Did         string
//	Psystem     string
//	Pmodel      string
//}

func main() {
	db, err := sqlx.Open("postgres",
		"postgresql://edwin@192.168.1.205:26257/tp_user?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("ping err: ", err)
	}
	//namedQuery(db)

	//selectQuery(db)

	Exec(db)

}

func Exec(db *sqlx.DB) {
	res, err := db.Exec("update accounts set nick = '一二三四五一一一一' where uid = $1", 230038894379009)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}

func getQuery(db *sqlx.DB) {
	account := new(ucenter.Account)
	udb := db.Unsafe()
	err := udb.Get(account, "select id,unionid,nick,addr,ctime from accounts where uid = 230038894379009")

	if err != nil {
		log.Fatal("select err: ", err)
	}
	log.Info("account: ", account)
}
func tx(db *sql.DB) {
	var id int64
	tx, err := db.Begin()
	if err != nil {
		log.Fatal("22222", err)
	}
	tx.QueryRow("insert into accounts ( unionid) values ('555552vvv3') returning id").Scan(&id)
	if id > 0 {
		tx.Rollback()
	} else {
		tx.Commit()
	}
	fmt.Println(tx)
	fmt.Println(id)

}

func selectQuery(db *sqlx.DB) {
	accounts := []ucenter.Account{}
	err := db.Select(&accounts, "select id,unionid,nick,addr from accounts where id = 2537037475827875")

	if err != nil {
		log.Fatal("select err: ", err)
	}
	log.Info("accounts: ", accounts)
}

//func namedQuery(db *sqlx.DB) {
//	others := "{\"character\":0}"
//	rows, err := db.NamedQuery("insert into accounts ( unionid)"+
//		"values (:unionid) returning id", &ucenter.Account{Unionid: "2222", Others: others})
//	if err != nil {
//		log.Fatal(err)
//	}
//	var id int
//	for rows.Next() {
//		rows.Scan(&id)
//	}
//
//	fmt.Println(id)
//}

func queryRow(db *sqlx.DB) {
	var id int64
	row := db.QueryRow("insert into accounts ( unionid)" +
		"values ('56565') returning id")
	row.Scan(&id)
}

//func stdsql() {
//	db, err := sql.Open("postgres",
//		"postgresql://edwin@192.168.1.51:26257/users?sslmode=disable")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	err = db.Ping()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	//
//	stmt, err := db.Prepare("SELECT id FROM account")
//	defer stmt.Close()
//	if err != nil {
//		log.Fatal(err)
//	}
//	rows, err := stmt.Query()
//	defer rows.Close()
//	for rows.Next() {
//		var id int
//		if err := rows.Scan(&id); err != nil {
//			log.Fatal(err)
//		}
//		log.Info(id)
//	}
//
//	//shortcut for single row
//	var id int
//	db.QueryRow("select id from account").Scan(&id)
//	log.Info(id)
//
//	stmt.QueryRow().Scan(&id)
//	log.Info(id)
//
//	stmt, err = db.Prepare("INSERT INTO account(unionid) VALUES($1) RETURNING id")
//	defer stmt.Close()
//	if err != nil {
//		log.Fatal(err)
//	}
//	res, err := stmt.Exec("dd")
//	if err != nil {
//		log.Fatal(err)
//	}
//	rowCnt, err := res.RowsAffected()
//	if err != nil {
//		log.Fatal(err)
//	}
//	var nid int
//	log.Printf("ID = %d, affected = %d\n", nid, rowCnt)
//}
