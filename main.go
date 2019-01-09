package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type InputData struct {
	Host  string
	Table string
	User  string
	Pwd   string
	Db    string
}

func checkErr(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		panic(err)
	}
}

func (i *InputData) parseFlags() {
	hostPtr := flag.String("host", "nope", "The name of the host mysql instance to connect with")
	tablePtr := flag.String("table", "nope", "The name of the mysql table we want to alter to InnoDB")
	userPtr := flag.String("user", "nope", "The username of the mysql instance to connect with")
	pwdPtr := flag.String("pwd", "nope", "This is the password to get into the mysql instance")
	dbPtr := flag.String("db", "", "This is the database to be used --optional")

	flag.Parse()
	i.Host = *hostPtr
	i.Table = *tablePtr
	i.User = *userPtr
	i.Pwd = *pwdPtr
	i.Db = *dbPtr

	if i.Host == "nope" || i.Table == "nope" || i.User == "nope" || i.Pwd == "nope" {
		log.Fatal("We are missing a flag. Need host, table, user, pwd. Did you include all of them?")
	}
}

func main() {

	id := new(InputData)
	id.parseFlags()

	connstr := fmt.Sprintf("%s:%s@%s/%s", id.User, id.Pwd, id.Host, id.Db)
	db, err := sql.Open("mysql", connstr)
	checkErr(err, "Error connecting to the database!")

	stmt, err := db.Prepare(fmt.Sprintf("ALTER TABLE %s ENGINE=InnoDB;", id.Table))
	checkErr(err, "Error with preparing statement to alter table")

	res, err := stmt.Exec()
	checkErr(err, "Error on query")

	affect, err := res.RowsAffected()
	checkErr(err, "Error on getting the number of rows affected")

	fmt.Println(affect)

	fmt.Sprintf("We have completed altering table %s from myIsam to InnoDb", id.Table)
	defer db.Close()
}
