package dbop

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

// G101 : Cannot connect to DB
// G102 : SQL statement error
// G103 : SQL exe error
// G104 : empty
func getPosts(path string) string {
	db, err := sql.Open(sqlDriver, userDataPath)
	if err != nil {
		log.Print(err)
		return "G101"
	}
	defer func() {
		_ = db.Close()
	}()
	stmtUsername, err := db.Prepare(`SELECT uid,createtime,content  FROM posts WHERE path = ?`)
	if err != nil {
		log.Print(err)
		return "G102"
	}
	defer func() {
		_ = stmtUsername.Close()
	}()
	var name string
	err = stmtUsername.QueryRow(path).Scan(&name)
	if err == sql.ErrNoRows {
		return "G104"
	} else if err != nil {
		log.Print(err)
		return "G103"
	}
	return name
}
