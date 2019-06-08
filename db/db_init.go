package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

const (
	sqlDriver    = "sqlite3"
	userDataPath = "./db/userdata.db"
)

func main() {
	_ = os.Remove(userDataPath)
	db, err := sql.Open(sqlDriver, userDataPath)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = db.Close()
	}()

	sqlStmt := `
	CREATE TABLE userinfo (
    uid INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    birthdate TEXT,
    createtime TEXT NOT NULL,
    gender TEXT);

CREATE TABLE posts
(
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    uid        INTEGER,
    createtime TEXT NOT NULL,
    content    TEXT NOT NULL,
    path       TEXT NOT NULL
);

CREATE TABLE boards
(
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    name       TEXT NOT NULL,
    path       TEXT NOT NULL UNIQUE
);

CREATE TABLE response
(
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    uid        INTEGER,
    pid        INTEGER,
    createtime TEXT NOT NULL,
    content    TEXT NOT NULL
);

CREATE TABLE  message
(
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    sender      INTEGER,
    sendee      INTEGER,
    createtime  TEXT NOT NULL,
    content     TEXT NOT NULL 
);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
	log.Print("OK")
	return
}
