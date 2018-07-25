package main

import (
	"github.com/jmoiron/sqlx"
	"log"
	"fmt"
	"os"
	"bufio"
	"strings"
	"io"
	_ "github.com/lib/pq"
)

func main() {
	fd, err := os.OpenFile("sql.sql", os.O_RDONLY, 0755)
	if err != nil {
		fmt.Println("err", err.Error())
		os.Exit(0)
	}
	defer fd.Close()

	buf := bufio.NewReader(fd)
	sqlsmt := ""
	i := 0
	for x := 0; x < 10000000; x++{
		s, e := buf.ReadString('\n')
		if e == io.EOF {
			if err = sql(sqlsmt); err != nil {
				fmt.Println("err", e.Error())
				os.Exit(0)
			}
			os.Exit(0)
		}
		if e != nil {
			fmt.Println("err", e.Error())
			os.Exit(0)
		}

		s = strings.Replace(s, "INSERT INTO `` VALUES ", "", -1)
		s = strings.Replace(s, ";", "", -1)
		s = strings.Replace(s, "\r\n", "", -1)


		if sqlsmt == "" {
			sqlsmt = s
		} else {
			sqlsmt = fmt.Sprintf("%v, %v", sqlsmt, s)
		}
		i++

		if i >= 10 {
			if err = sql(sqlsmt); err != nil {
				fmt.Println("err", e.Error())
				os.Exit(0)
			}
			i = 0
			sqlsmt = ""
			fmt.Println("x = ", x)
		}

	}
}


func sql(vls string) (err error) {
	dsn := "postgresql://edwin@127.0.0.1:26256/tp_user?sslmode=disable"
	var conn *sqlx.DB
	conn, err = sqlx.Open("postgres", dsn)
	if err != nil {
		log.Println("connect", err.Error())
		os.Exit(0)
	}
	defer conn.Close()
	sqlsmt := fmt.Sprintf("insert into tp_user.yz_accounts (uid,nick,gender,mac,imei,imsi,ltype,coin,gem,bank,level,vipexpiry,gifts,win,fail,recharge,uuid) values %v", vls)
	_, err = conn.Exec(sqlsmt)
	if err != nil {
		fmt.Println("err", err.Error(), "sql", sqlsmt)
		os.Exit(0)
	}
	return
}
