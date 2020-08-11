package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/shakilbd009/go-utils-lib/logger"
)

const (
	mysql_user_username = "mysql_user_username"
	mysql_user_password = "mysql_user_password"
	mysql_user_hostname = "mysql_user_hostname"
	mysql_user_schema   = "mysql_user_schema"
)

var (
	Client   *sql.DB
	username = os.Getenv(mysql_user_username)
	password = os.Getenv(mysql_user_password)
	hostname = os.Getenv(mysql_user_hostname)
	schema   = os.Getenv(mysql_user_schema)
)

func init() {
	var err error
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, hostname, schema)
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err := Client.Ping(); err != nil {
		panic(err)
	}
	mysql.SetLogger(logger.GetLogger())
	log.Println("database successfully configured")
}
