package dbop

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	// Sqlite driver for go
	_ "github.com/mattn/go-sqlite3"
)

const (
	sqlDriver = "sqlite3"
	// relative path of databases
	userDataPath = "./db/userdata.db"
)

// AddUser add user
func AddUser(uid, username, password, email, phone, birthdate, gender string) bool {
	defer printErr()
	db, err := sql.Open(sqlDriver, userDataPath)
	checkErr(err)
	defer db.Close()
	stmt, err := db.Prepare(`INSERT INTO userinfo(uid, 
												  username, 
												  password, 
												  email, 
												  phone, 
												  birthdate, 
												  createtime, 
												  gender) values (?, ?, ?, ?, ?, ?, ?, ?)`)
	checkErr(err)
	createtime := time.Now().Format("2006-01-02 15:04:05")
	tx, err := db.Begin()
	checkErr(err)
	_, err = tx.Stmt(stmt).Exec(uid, username, password, email, phone, birthdate, createtime, gender)
	if err != nil {
		tx.Rollback()
		panic(err)
	} else {
		tx.Commit()
	}
	return true
}

// DelUser delete user using uid
func DelUser(uid string) bool {
	defer printErr()
	db, err := sql.Open(sqlDriver, userDataPath)
	checkErr(err)
	defer db.Close()
	stmt, err := db.Prepare("DELETE FROM userinfo where uid=?")
	checkErr(err)
	tx, err := db.Begin()
	checkErr(err)
	_, err = tx.Stmt(stmt).Exec(uid)
	if err != nil {
		tx.Rollback()
		panic(err)
	} else {
		tx.Commit()
	}
	return true
}

// ModifyUser modify information of a user
// format of array attrAndValue: attrName0, value0, attrName1, value1...
func ModifyUser(uid string, attrAndValue ...string) bool {
	defer printErr()

	var stmtBuilder strings.Builder
	stmtBuilder.WriteString("UPDATE userinfo SET ")
	stmtBuilder.WriteString(fmt.Sprintf("%s='%s'", attrAndValue[0], attrAndValue[1]))
	for i := 2; i < len(attrAndValue); i += 2 {
		stmtBuilder.WriteString(fmt.Sprintf(", %s='%s'", attrAndValue[i], attrAndValue[i+1]))
	}
	stmtBuilder.WriteString(fmt.Sprintf(" WHERE uid=%s", uid))

	db, err := sql.Open(sqlDriver, userDataPath)
	checkErr(err)
	defer db.Close()
	stmt, err := db.Prepare(stmtBuilder.String())
	checkErr(err)
	tx, err := db.Begin()
	checkErr(err)
	_, err = tx.Stmt(stmt).Exec()
	if err != nil {
		tx.Rollback()
		panic(err)
	} else {
		tx.Commit()
	}
	return true
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func printErr() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
}
