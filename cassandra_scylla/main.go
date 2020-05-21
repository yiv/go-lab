package main

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
)

type Row struct {
	id          int32   `json:"id"`
	t_tinyint   byte    `json:"t_tinyint"`
	t_smallint  int32   `json:"t_smallint"`
	t_bigint    int64   `json:"t_bigint"`
	t_ascii     string  `json:"t_ascii"`
	t_text      string  `json:"t_text"`
	t_varchar   string  `json:"t_varchar"`
	t_boolean   bool    `json:"t_boolean"`
	t_time      int64   `json:"t_time"`
	t_timestamp int64   `json:"t_timestamp"`
	t_float     float32 `json:"t_float"`
	t_double    float64 `json:"t_double"`
	// t_uuid uuid
	// t_blob blob
	t_list []string          `json:"t_list"`
	t_set  []int32           `json:"t_set"`
	t_map  map[string]string `json:"t_map"`
}

func main() {
	cluster := gocql.NewCluster("106.52.187.25:9042")
	cluster.Keyspace = "ks_edwin"
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: "cassandra",
		Password: "Brysj@1gsycl",
	}
	//cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()

	var (
		id int64
		//row = Row{}
		myset = []int32{}
	)

	//iter := session.Query(`SELECT id,t_tinyint,t_smallint,t_bigint,t_ascii,t_text,t_varchar,t_boolean,t_time,t_timestamp,t_float,t_double,t_list,t_set,t_map FROM ks_edwin.my_test_table where id = 2`).Iter()
	//smt := `SELECT id,t_tinyint,t_smallint,t_bigint,t_ascii,t_text,t_varchar,t_boolean,t_time,t_timestamp,t_float,t_double,t_list,t_set,t_map FROM ks_edwin.my_test_table where id = 2`
	smt := `select id,t_set from ks_edwin.my_test_table where id = 2`
	if err := session.Query(smt).Scan(&id, &myset); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Tweet:", id, myset)

	fmt.Println("heel")
}
