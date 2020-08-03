package databases

import (
	"database/sql"
	"log"
)
import _ "github.com/go-sql-driver/mysql"

var MySql *sql.DB

func init() {
	dataSource := "root:123456@tcp(localhost:3306)/test?parseTime=true"
	var err error
	//MySql, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?parseTime=true")

	MySql, err = sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatal(err)
	}
	err = MySql.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}
