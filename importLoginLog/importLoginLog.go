package main

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
	"github.com/urfave/cli"
	"io"
	"os"
	"strings"
	"time"
)

var (
	dbUser     string
	dbPassword string
)

type importProcesser struct {
	dbUser      string
	dbPassword  string
	logFileName string
	db          *sql.DB
}

type logEntry struct {
	Uid       int64   `json:"uid"`
	Uuid      string  `json:"uuid"`
	Account   string  `json:"account"`
	Password  string  `json:"password"`
	LoginType string  `json:"type"`
	Unionid   string  `json:"unionid"`
	Ip        string  `json:"ip"`
	Imsi      string  `json:"imsi"`
	Mac       string  `json:"mac"`
	Imei      string  `json:"imei"`
	Pmodel    string  `json:"pmodel"`
	Dateline  string  `json:"dateline"`
	Time      float32 `json:"time"`
}

func init() {
	log.SetLevel(log.DebugLevel)
	log.SetLevel(log.InfoLevel)
	// log.SetLevel(log.FatalLevel)
}

func main() {
	app := &cli.App{
		Name: "import user login logs to mysql",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "user,u",
				Value: "root",
				Usage: "db account",
			},
			&cli.StringFlag{
				Name:  "password,p",
				Value: "mysql003@yz.COM",
				Usage: "db password",
			},
			&cli.StringFlag{
				Name:  "path",
				Value: "/var/www/html/wrzjh/service/logs/",
				Usage: "directory where the log files stored",
			},
		},
		Action: func(c *cli.Context) error {
			path := c.String("path")
			pdir := fmt.Sprintf("/%s", time.Now().Format("200601"))
			fname := fmt.Sprintf("/login%s", time.Now().Add(-time.Hour*24).Format("2006-01-02"))
			imp := &importProcesser{
				dbUser:      c.String("user"),
				dbPassword:  c.String("password"),
				logFileName: path + pdir + fname,
			}

			imp.opendb()

			log.Debug(imp.db)

			file, err := os.Open(imp.logFileName)
			if err != nil {
				log.Fatalf("open login file failed: %s", err)
			}
			fileReader := bufio.NewReader(file)

			var counter int

			for {
				logStr, err := fileReader.ReadString('\n')
				if err == io.EOF {
					log.Infof("\n%s, count: %v", err, counter)
					break
				}
				if err != nil {
					log.Error("read log line failed")
				}
				logStr = strings.Replace(logStr, "][", "\",\"ip\":\"", 1)
				logStr = strings.Replace(logStr, "[", "{\"dateline\":\"", 1)
				logStr = strings.Replace(logStr, "]{", "\",", 1)
				logStr = strings.TrimSpace(logStr)

				logRe := logEntry{}
				err = json.Unmarshal([]byte(logStr), &logRe)
				if err != nil {
					log.Error(err)
				}
				logRe.Dateline = fmt.Sprintf("%s%s", time.Now().Format("2006-"), logRe.Dateline)
				log.Debug(logRe)

				imp.insert(&logRe)
				counter++
				if counter%1000 == 0 {
					fmt.Printf(".")
				}

			}

			return nil
		},
	}
	app.Run(os.Args)
}

func (imp *importProcesser) opendb() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/logdb", imp.dbUser, imp.dbPassword))
	checkError(err)
	err = db.Ping()
	checkError(err)
	imp.db = db
}
func (imp *importProcesser) insert(l *logEntry) {
	stmt, err := imp.db.Prepare("INSERT INTO loginLog(uid, uuid,account,password,type,unionid,ip,imsi,mac,imei,pmodel,dateline,time) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		log.Errorf("Prepare insert db error: %s", err)
	}
	defer stmt.Close()
	stmt.Exec(l.Uid, l.Uuid, l.Account, l.Password, l.LoginType, l.Unionid, l.Ip, l.Imsi, l.Mac, l.Imei, l.Pmodel, l.Dateline, l.Time)
}
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
}
