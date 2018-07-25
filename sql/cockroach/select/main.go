package main

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"fmt"
)

func main() {
	dsn := "postgresql://edwin@127.0.0.1:26256/tp_mall?sslmode=disable"
	conn, err := sqlx.Open("postgres", dsn)
	if err != nil {
		log.Println("connect", err.Error())
	}
	type Sum struct {
		Date  string
		Count int32
	}
	start, err := time.Parse("2006-01-02", "2018-01-01")
	if err != nil {
		panic(err.Error())
	}
	end, err := time.Parse("2006-01-02", "2018-01-12")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("########## mol molPlayersAmount ######################")
	molPlayersAmount(conn ,"tp_user", start, end )
	fmt.Println("########## playersAmount ######################")
	playersAmount(conn ,"tp_user", start, end )
	fmt.Println("########## allOrders ######################")
	allOrders(conn ,"tp_mall", start, end )
	fmt.Println("########## paidOrders ######################")
	paidOrders(conn ,"tp_mall", start, end )
	fmt.Println("########## ordersAmount ######################")
	ordersAmount(conn ,"tp_mall", start, end )
}

func allOrders(conn *sqlx.DB,db string,  start, end time.Time ) (err error) {
	_, err = conn.Exec(fmt.Sprintf("set database = '%s'", db))
	if err != nil {
		return
	}
	for  {
		var count int32
		sql := fmt.Sprintf("select count(id) from uc_orders where ctime >= %d and ctime < %d", start.Unix(), start.Add(time.Hour * 24).Unix())
		//fmt.Println(sql)
		err = conn.Get(&count, sql)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("%s\t%d\n", start.Format("2006-01-02"), count)
		start = start.Add(time.Hour * 24)
		if start.Equal(end) {
			break
		}
	}
	return
}


func paidOrders(conn *sqlx.DB,db string, start, end time.Time ) (err error) {
	_, err = conn.Exec(fmt.Sprintf("set database = '%s'", db))
	if err != nil {
		return
	}
	for  {
		var count int32
		sql := fmt.Sprintf("select count(id) from uc_orders where ctime >= %d and ctime < %d and payorder != ''", start.Unix(), start.Add(time.Hour * 24).Unix())
		//fmt.Println(sql)
		err = conn.Get(&count, sql)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("%s\t%d\n", start.Format("2006-01-02"), count)
		start = start.Add(time.Hour * 24)
		if start.Equal(end) {
			break
		}
	}
	return
}
func ordersAmount(conn *sqlx.DB,db string, start, end time.Time ) (err error) {
	_, err = conn.Exec(fmt.Sprintf("set database = '%s'", db))
	if err != nil {
		return
	}
	for  {
		var count float32
		sql := fmt.Sprintf("select sum(amt) as a from uc_orders where ctime >= %d and ctime < %d and payorder != ''", start.Unix(), start.Add(time.Hour * 24).Unix())
		//fmt.Println(sql)
		err = conn.Get(&count, sql)
		if err != nil {
			//fmt.Println(err.Error())
		}
		fmt.Printf("%s\t%v\n", start.Format("2006-01-02"), count)
		start = start.Add(time.Hour * 24)
		if start.Equal(end) {
			break
		}
	}
	return
}
func playersAmount(conn *sqlx.DB,db string, start, end time.Time ) (err error) {
	_, err = conn.Exec(fmt.Sprintf("set database = '%s'", db))
	if err != nil {
		return
	}
	for  {
		var count float32
		sql := fmt.Sprintf("select count(id) as a from accounts where ctime >= %d and ctime < %d and unionid='uc'", start.Unix(), start.Add(time.Hour * 24).Unix())
		//fmt.Println(sql)
		err = conn.Get(&count, sql)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Printf("%s\t%v\n", start.Format("2006-01-02"), count)
		start = start.Add(time.Hour * 24)
		if start.Equal(end) {
			break
		}
	}
	return
}
func molPlayersAmount(conn *sqlx.DB,db string, start, end time.Time ) (err error) {
	_, err = conn.Exec(fmt.Sprintf("set database = '%s'", db))
	if err != nil {
		return
	}
	for  {
		var count float32
		sql := fmt.Sprintf("select count(id) as a from accounts where ctime >= %d and ctime < %d and unionid='mol'", start.Unix(), start.Add(time.Hour * 24).Unix())
		//fmt.Println(sql)
		err = conn.Get(&count, sql)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Printf("%s\t%v\n", start.Format("2006-01-02"), count)
		start = start.Add(time.Hour * 24)
		if start.Equal(end) {
			break
		}
	}
	return
}