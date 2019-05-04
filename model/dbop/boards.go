package dbop

import (
	"database/sql"
	"encoding/json"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Posts struct {
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
	stmtUsername, err := db.Prepare(`SELECT uid,createtime,content  FROM posts WHERE path = ?`)
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
		var uid int
		var createtime, content string
		err = rows.Scan(&uid, &createtime, &content)
		if err != nil {
			log.Print(err)
			return "Unkonwn Error"
		}
		result.Posts = append(result.Posts, Posts{Uid: uid, Createtime: createtime, Content: content})
	}
	err = rows.Err()
	if err != nil {
		log.Print(err)
		return "Unkonwn Error"
	}

	res, err := json.Marshal(result)
	log.Print(result)
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
	log.Print(result)
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
	log.Print(result)
	if err != nil {
		log.Print(err)
		return "G105"
	}
	return string(res)
}