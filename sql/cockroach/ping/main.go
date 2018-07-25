package main

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	dsn := "postgresql://edwin@cockroach-public:26257/tp_mall?sslmode=disable"
	log.Println("edwin #1")
	for {
		log.Println("edwin #2")
		conn, err := sqlx.Open("postgres", dsn)
		log.Println("edwin #3")
		if err != nil {
			log.Println("connect", err.Error())
		}
		log.Println("edwin #4")
		err = conn.Ping()
		log.Println("edwin #5")
		if err != nil {
			log.Println("ping", err.Error())
		} else {
			var node int64
			if err = conn.Get(&node, "SELECT node_id FROM crdb_internal.node_build_info LIMIT 1"); err != nil {
				log.Println("get", err.Error())
			}
			log.Println("pong ", node)
		}
		conn.Close()

		time.Sleep(time.Second)
	}
}
