package dbop

import (
	"database/sql"
	"encoding/json"
	"github.com/HarborYuan/xforum/model/util"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

const (
	sqlDriver    = "sqlite3"
	userDataPath = "./db/userdata.db"
)

// Check Functions have the following return codes :
// C100-1 : There is a same username
// C100-0 : There is no same username
// C101 : Cannot connect to the database.
// C102 : SQL statement error
// C103 : SQL execution error
func CheckUsername(username string) string {
	db, err := sql.Open(sqlDriver, userDataPath)
	if err != nil {
		log.Print(err)
		return "C101"
	}
	defer func() {
		_ = db.Close()
	}()
	stmtUsername, err := db.Prepare(`SELECT uid FROM userinfo WHERE username = ?`)
	if err != nil {
		log.Print(err)
		return "C102"
	}
	defer func() {
		_ = stmtUsername.Close()
	}()
	var name string
	err = stmtUsername.QueryRow(username).Scan(&name)
	if err == sql.ErrNoRows {
		return "C100-0"
	} else if err != nil {
		log.Print(err)
		return "C103"
	}
	return "C100-1"
}

// Check Functions have the following return codes :
// C100-1 : There is a same email
// C100-0 : There is no same email
// C101 : Cannot connect to the database.
// C102 : SQL statement error
// C103 : SQL execution error
func CheckEmail(email string) string {
	db, err := sql.Open(sqlDriver, userDataPath)
	if err != nil {
		log.Print(err)
		return "C101"
	}
	defer func() {
		_ = db.Close()
	}()
	stmtUsername, err := db.Prepare(`SELECT uid FROM userinfo WHERE email = ?`)
	if err != nil {
		log.Print(err)
		return "C102"
	}
	defer func() {
		_ = stmtUsername.Close()
	}()
	var name string
	err = stmtUsername.QueryRow(email).Scan(&name)
	if err == sql.ErrNoRows {
		return "C100-0"
	} else if err != nil {
		log.Print(err)
		return "C103"
	}
	return "C100-1"
}

// AddUser add user
// Return codes:
// U100 : Success!
// U101 : Cannot connect to the database.
// U102 : SQL statement error
// U103 : Unable to get DB handle
// U104 : Execution error
func AddUser(username, password, email, birthdate, gender string) string {
	// Open database
	db, err := sql.Open(sqlDriver, userDataPath)
	if err != nil {
		log.Print(err)
		return "U101"
	}
	defer func() {
		_ = db.Close()
	}()

	// Prepare for statement
	stmt, err := db.Prepare(`INSERT INTO userinfo(username, password, email, birthdate, createtime, gender) values (?, ?, ?, ?, ?, ?)`)
	if err != nil {
		log.Print(err)
		return "U102"
	}

	createtime := time.Now().Format("2006-01-02 15:04:05")

	tx, err := db.Begin()
	if err != nil {
		log.Print(err)
		return "U103"
	}

	_, err = tx.Stmt(stmt).Exec(username, util.Sha(password), email, birthdate, createtime, gender)
	if err != nil {
		_ = tx.Rollback()
		log.Print(err)
		return "U104"
	}
	_ = tx.Commit()
	return "U100"
}

// CheckPass have the following return codes :
// C100-1 : OK
// C100-0 : Wrong password
// C101 : Cannot connect to the database.
// C102 : SQL statement error
// C103 : SQL execution error
// C104 : No user
func CheckPass(username, password string) string {
	db, err := sql.Open(sqlDriver, userDataPath)
	if err != nil {
		log.Print(err)
		return "C101"
	}
	defer func() {
		_ = db.Close()
	}()
	stmtUsername, err := db.Prepare(`SELECT password FROM userinfo WHERE username = ?`)
	if err != nil {
		log.Print(err)
		return "C102"
	}
	defer func() {
		_ = stmtUsername.Close()
	}()
	var name string
	err = stmtUsername.QueryRow(username).Scan(&name)
	if err == sql.ErrNoRows {
		return "C104"
	} else if err != nil {
		log.Print(err)
		return "C103"
	}
	if name == util.Sha(password) {
		return "C100-1"
	}
	return "C100-0"
}

func GetUid(name string) int {
	db, err := sql.Open(sqlDriver, userDataPath)
	if err != nil {
		log.Print(err)
		return -1
	}
	defer func() {
		_ = db.Close()
	}()
	stmtUsername, err := db.Prepare(`SELECT uid FROM userinfo WHERE username = ?`)
	if err != nil {
		log.Print(err)
		return -1
	}
	defer func() {
		_ = stmtUsername.Close()
	}()
	var res int
	err = stmtUsername.QueryRow(name).Scan(&res)
	if err != nil {
		log.Print(err)
		return -1
	}
	return res
}

type UserInfo struct {
	Uid      int    `json:"uid"`
	Username string `json:"username"`
	Gender   string `json:"gender"`
}

// G101 : Cannot connect to DB
// G102 : SQL statement error
// G103 : SQL exe error
// G104 : empty
// G105 : JSON error
func GetUserInfo(id int) string {
	db, err := sql.Open(sqlDriver, userDataPath)
	if err != nil {
		log.Print(err)
		return "G101"
	}
	defer func() {
		_ = db.Close()
	}()
	stmtUsername, err := db.Prepare(`SELECT uid, username,gender FROM userinfo WHERE uid = ?`)
	if err != nil {
		log.Print(err)
		return "G102"
	}
	defer func() {
		_ = stmtUsername.Close()
	}()
	var uid int
	var username, gender string
	err = stmtUsername.QueryRow(id).Scan(&uid, &username, &gender)
	if err == sql.ErrNoRows {
		return "G104"
	} else if err != nil {
		log.Print(err)
		return "G103"
	}

	var result UserInfo
	result = UserInfo{Uid: uid, Username: username, Gender: gender}
	res, err := json.Marshal(result)
	if err != nil {
		log.Print(err)
		return "G105"
	}
	return string(res)
}

type UserDetailInfo struct {
	Uid        int    `json:"uid"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Birthdate  string `json:"birthdate"`
	Createtime string `json:"createtime"`
	Gender     string `json:"gender"`
}

// G101 : Cannot connect to DB
// G102 : SQL statement error
// G103 : SQL exe error
// G104 : empty
// G105 : JSON error
func GetUserDetailInfo(id int) string {
	db, err := sql.Open(sqlDriver, userDataPath)
	if err != nil {
		log.Print(err)
		return "G101"
	}
	defer func() {
		_ = db.Close()
	}()
	stmtUsername, err := db.Prepare(`SELECT uid, username, email, birthdate, createtime, gender FROM userinfo WHERE uid = ?`)
	if err != nil {
		log.Print(err)
		return "G102"
	}
	defer func() {
		_ = stmtUsername.Close()
	}()
	var uid int
	var username, email, birthdate, createtime, gender string
	err = stmtUsername.QueryRow(id).Scan(&uid, &username, &email, &birthdate, &createtime, &gender)
	if err == sql.ErrNoRows {
		return "G104"
	} else if err != nil {
		log.Print(err)
		return "G103"
	}

	var result UserDetailInfo
	result = UserDetailInfo{Uid: uid, Username: username, Email: email, Birthdate: birthdate, Createtime: createtime, Gender: gender}
	res, err := json.Marshal(result)
	if err != nil {
		log.Print(err)
		return "G105"
	}
	return string(res)
}

func GetUserName(id int) string {
	db, err := sql.Open(sqlDriver, userDataPath)
	if err != nil {
		log.Print(err)
		return "@"
	}
	defer func() {
		_ = db.Close()
	}()
	stmtUsername, err := db.Prepare(`SELECT username FROM userinfo WHERE uid = ?`)
	if err != nil {
		log.Print(err)
		return "@"
	}
	defer func() {
		_ = stmtUsername.Close()
	}()
	var username string
	err = stmtUsername.QueryRow(id).Scan(&username)
	if err == sql.ErrNoRows {
		return "!"
	} else if err != nil {
		log.Print(err)
		return "@"
	}
	return string(username)
}

//// DelUser delete user using uid
//func DelUser(uid string) bool {
//	defer printErr()
//	db, err := sql.Open(sqlDriver, userDataPath)
//	checkErr(err)
//	defer db.Close()
//	stmt, err := db.Prepare("DELETE FROM userinfo where uid=?")
//	checkErr(err)
//	tx, err := db.Begin()
//	checkErr(err)
//	_, err = tx.Stmt(stmt).Exec(uid)
//	if err != nil {
//		tx.Rollback()
//		panic(err)
//	} else {
//		tx.Commit()
//	}
//	return true
//}
//
//// ModifyUser modify information of a user
//// format of array attrAndValue: attrName0, value0, attrName1, value1...
//func ModifyUser(uid string, attrAndValue ...string) bool {
//	defer printErr()
//
//	var stmtBuilder strings.Builder
//	stmtBuilder.WriteString("UPDATE userinfo SET ")
//	stmtBuilder.WriteString(fmt.Sprintf("%s='%s'", attrAndValue[0], attrAndValue[1]))
//	for i := 2; i < len(attrAndValue); i += 2 {
//		stmtBuilder.WriteString(fmt.Sprintf(", %s='%s'", attrAndValue[i], attrAndValue[i+1]))
//	}
//	stmtBuilder.WriteString(fmt.Sprintf(" WHERE uid=%s", uid))
//
//	db, err := sql.Open(sqlDriver, userDataPath)
//	checkErr(err)
//	defer db.Close()
//	stmt, err := db.Prepare(stmtBuilder.String())
//	checkErr(err)
//	tx, err := db.Begin()
//	checkErr(err)
//	_, err = tx.Stmt(stmt).Exec()
//	if err != nil {
//		tx.Rollback()
//		panic(err)
//	} else {
//		tx.Commit()
//	}
//	return true
//}
