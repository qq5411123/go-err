package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"log"
)

type Person struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
}

var Persons []Person

var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatalf("%+v", errors.Wrap(err, "open mysql failed!"))
	}
	Db = database
}

func GetPerson() error {
	err := Db.Select(&Persons, "select *from person")

	switch {

	//空数据属于业务逻辑 return
	case err == sql.ErrNoRows:

		return errors.Wrap( err,  "GetPerson No Data" )

	//datebase exception stop run
	case err != nil:

		wrapErr := errors.Wrap(err, "func GetPerson")
		log.Printf("%+v", wrapErr)

	}

	return nil

}



func IsErrNoRows(err error) bool {
	if err == sql.ErrNoRows {
		return true
	}

	return false
}