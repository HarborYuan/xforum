package dbop

import (
	"database/sql"
	"encoding/json"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

type Posts struct {
	Pid        int    `json:"pid"`
	Uid        int    `json:"uid"`
	Createtime string `json:"createtime"`
	Content    string `json:"content"`
}

type AllPosts struct {
	Posts []Posts `json:"posts"`
}

// G101 : Cannot connect to DB
// G102 : SQL statement error
// G103 : SQL exe error
// G104 : empty
// G105 : JSON error
func GetPosts(path string) string {
	db, err := sql.Open(sqlDriver, userDataPath)
	if err != nil {
		log.Print(err)
		return "G101"
	}
	defer func() {
		_ = db.Close()
	}()
	stmtUsername, err := db.Prepare(`SELECT id,uid,createtime,content  FROM posts WHERE path = ?`)
	if err != nil {
		log.Print(err)
		return "G102"
	}
	defer func() {
		_ = stmtUsername.Close()
	}()
	rows, err := stmtUsername.Query(path)
	if err == sql.ErrNoRows {
		return "G104"
	} else if err != nil {
		log.Print(err)
		return "G103"
	}
	var result AllPosts
	defer rows.Close()
	for rows.Next() {
		var pid, uid int
		var createtime, content string
		err = rows.Scan(&pid, &uid, &createtime, &content)
		if err != nil {
			log.Print(err)
			return "Unkonwn Error"
		}
		result.Posts = append(result.Posts, Posts{Pid: pid, Uid: uid, Createtime: createtime, Content: content})
	}
	err = rows.Err()
	if err != nil {
		log.Print(err)
		return "Unkonwn Error"
	}

	res, err := json.Marshal(result)
	if err != nil {
		log.Print(err)
		return "G105"
	}
	return string(res)
}

type Board struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type Boards struct {
	Board []Board `json:"board"`
}

// G101 : Cannot connect to DB
// G102 : SQL statement error
// G103 : SQL exe error
// G104 : empty
// G105 : JSON error
func GetBoards() string {
	db, err := sql.Open(sqlDriver, userDataPath)
	if err != nil {
		log.Print(err)
		return "G101"
	}
	defer func() {
		_ = db.Close()
	}()
	stmtUsername, err := db.Prepare(`SELECT name,path FROM boards`)
	if err != nil {
		log.Print(err)
		return "G102"
	}
	defer func() {
		_ = stmtUsername.Close()
	}()
	rows, err := stmtUsername.Query()
	if err == sql.ErrNoRows {
		return "G104"
	} else if err != nil {
		log.Print(err)
		return "G103"
	}
	var result Boards
	defer rows.Close()
	for rows.Next() {
		var name, path string
		err = rows.Scan(&name, &path)
		if err != nil {
			log.Print(err)
			return "Unkonwn Error"
		}
		result.Board = append(result.Board, Board{Name: name, Path: path})
	}
	err = rows.Err()
	if err != nil {
		log.Print(err)
		return "Unkonwn Error"
	}

	res, err := json.Marshal(result)
	if err != nil {
		log.Print(err)
		return "G105"
	}
	return string(res)
}

type Response struct {
	Uid        int    `json:"uid"`
	Createtime string `json:"createtime"`
	Content    string `json:"content"`
}

type Responses struct {
	Response []Response `json:"response"`
}

// G101 : Cannot connect to DB
// G102 : SQL statement error
// G103 : SQL exe error
// G104 : empty
// G105 : JSON error

func GetResponse(pid int) string {
	db, err := sql.Open(sqlDriver, userDataPath)
	if err != nil {
		log.Print(err)
		return "G101"
	}
	defer func() {
		_ = db.Close()
	}()
	stmtUsername, err := db.Prepare(`SELECT uid,createtime,content  FROM response WHERE pid = ?`)
	if err != nil {
		log.Print(err)
		return "G102"
	}
	defer func() {
		_ = stmtUsername.Close()
	}()
	rows, err := stmtUsername.Query(pid)
	if err == sql.ErrNoRows {
		return "G104"
	} else if err != nil {
		log.Print(err)
		return "G103"
	}
	var result Responses
	defer rows.Close()
	for rows.Next() {
		var uid int
		var createtime, content string
		err = rows.Scan(&uid, &createtime, &content)
		log.Print(uid)
		if err != nil {
			log.Print(err)
			return "Unkonwn Error"
		}
		result.Response = append(result.Response, Response{Uid: uid, Createtime: createtime, Content: content})
	}
	err = rows.Err()
	if err != nil {
		log.Print(err)
		return "Unkonwn Error"
	}

	res, err := json.Marshal(result)
	if err != nil {
		log.Print(err)
		return "G105"
	}
	return string(res)
}

// B100 : success
// B101 : Cannot connect to DB
// B102 : sql statement error
// B103 : Unable to get DB handle
// B104 : sql exe error
func AddBoard(name, path string) string {
	// Open database
	db, err := sql.Open(sqlDriver, userDataPath)
	if err != nil {
		log.Print(err)
		return "B101"
	}
	defer func() {
		_ = db.Close()
	}()

	// Prepare for statement
	stmt, err := db.Prepare(`INSERT INTO boards(name, path) values (?, ?)`)
	if err != nil {
		log.Print(err)
		return "B102"
	}

	//createtime := time.Now().Format("2006-01-02 15:04:05")

	tx, err := db.Begin()
	if err != nil {
		log.Print(err)
		return "B103"
	}

	_, err = tx.Stmt(stmt).Exec(name, path)
	if err != nil {
		_ = tx.Rollback()
		log.Print(err)
		return "B104"
	}
	_ = tx.Commit()
	return "B100"
}

// B100 : success
// B101 : Cannot connect to DB
// B102 : sql statement error
// B103 : Unable to get DB handler
// B104 : sql exe error
// B105 : user invalid
func AddPost(content, path string, uid int) string {
	// Open database
	db, err := sql.Open(sqlDriver, userDataPath)
	if err != nil {
		log.Print(err)
		return "B101"
	}
	defer func() {
		_ = db.Close()
	}()

	// Prepare for statement
	stmt, err := db.Prepare(`INSERT INTO posts (uid, createtime, content, path) values (?, ?, ?, ?)`)
	if err != nil {
		log.Print(err)
		return "B102"
	}

	createtime := time.Now().Format("2006-01-02 15:04:05")

	tx, err := db.Begin()
	if err != nil {
		log.Print(err)
		return "B103"
	}

	_, err = tx.Stmt(stmt).Exec(uid, createtime, content, path)
	if err != nil {
		_ = tx.Rollback()
		log.Print(err)
		return "B104"
	}
	_ = tx.Commit()
	return "B100"
}

// R100 : success
// R101 : Cannot connect to DB
// R102 : sql statement error
// R103 : Unable to get DB handler
// R104 : sql exe error
// R105 : user invalid
// R106 : No such pid
func AddResponse(uid, pid int, content string) string {
	// Open database
	db, err := sql.Open(sqlDriver, userDataPath)
	if err != nil {
		log.Print(err)
		return "R101"
	}
	defer func() {
		_ = db.Close()
	}()

	// Prepare for statement
	stmt, err := db.Prepare(`INSERT INTO response (uid, pid, createtime, content) values (?, ?, ?, ?)`)
	if err != nil {
		log.Print(err)
		return "R102"
	}

	createtime := time.Now().Format("2006-01-02 15:04:05")

	tx, err := db.Begin()
	if err != nil {
		log.Print(err)
		return "R103"
	}

	_, err = tx.Stmt(stmt).Exec(uid, pid, createtime, content)
	if err != nil {
		_ = tx.Rollback()
		log.Print(err)
		return "R104"
	}
	_ = tx.Commit()
	return "R100"
}

func IsExistPid(pid int) int {
	db, err := sql.Open(sqlDriver, userDataPath)
	if err != nil {
		log.Print(err)
		return -1
	}
	defer func() {
		_ = db.Close()
	}()
	stmtUsername, err := db.Prepare(`SELECT id FROM posts WHERE id = ?`)
	if err != nil {
		log.Print(err)
		return -1
	}
	defer func() {
		_ = stmtUsername.Close()
	}()
	var res int
	err = stmtUsername.QueryRow(pid).Scan(&res)
	if err == sql.ErrNoRows {
		return 0
	} else if err != nil {
		log.Print(err)
		return -1
	}
	return 1
}
