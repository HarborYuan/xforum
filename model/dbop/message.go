package dbop

import (
	"database/sql"
	"log"
	"time"
)

// M100 : success
// M101 : Cannot connect to DB
// M102 : sql statement error
// M103 : Unable to get DB handler
// M104 : sql exe error
// M105 : user invalid
func AddMessage(sender, sendee int, content string) string {
	// Open database
	db, err := sql.Open(sqlDriver, userDataPath)
	if err != nil {
		log.Print(err)
		return "M101"
	}
	defer func() {
		_ = db.Close()
	}()

	// Prepare for statement
	stmt, err := db.Prepare(`INSERT INTO message (sender,sendee, createtime, content) values (?, ?, ?, ?)`)
	if err != nil {
		log.Print(err)
		return "M102"
	}

	createtime := time.Now().Format("2006-01-02 15:04:05")

	tx, err := db.Begin()
	if err != nil {
		log.Print(err)
		return "M103"
	}

	_, err = tx.Stmt(stmt).Exec(sender, sendee, createtime, content)
	if err != nil {
		_ = tx.Rollback()
		log.Print(err)
		return "M104"
	}
	_ = tx.Commit()
	return "M100"
}
