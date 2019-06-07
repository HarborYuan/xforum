package dbop

import (
	"database/sql"
	"encoding/json"
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

type AllMessages struct {
	Messages []MessageInfo `json:"messages"`
}

type MessageInfo struct {
	Sender     string `json:"sender"`
	Sendee     string `json:"sendee"`
	Createtime string `json:"createtime"`
	Content    string `json:"content"`
}

// G101 : Cannot connect to DB
// G102 : SQL statement error
// G103 : SQL exe error
// G104 : empty
// G105 : JSON error
// G106 : User error
func GetMessage(sender, sendee int) string {
	db, err := sql.Open(sqlDriver, userDataPath)
	if err != nil {
		log.Print(err)
		return "G101"
	}
	defer func() {
		_ = db.Close()
	}()
	stmt1, err := db.Prepare(`SELECT sender,sendee,createtime,content
										FROM message 
										WHERE (sender=? AND sendee=?) OR (sender=? AND sendee=?)`)
	if err != nil {
		log.Print(err)
		return "G102"
	}
	defer func() {
		_ = stmt1.Close()
	}()
	rows, err := stmt1.Query(sender, sendee, sendee, sender)
	if err == sql.ErrNoRows {
		return "G104"
	} else if err != nil {
		log.Print(err)
		return "G103"
	}
	var result AllMessages
	defer rows.Close()
	isNotEmpty := false
	for rows.Next() {
		isNotEmpty = true
		var senderRes, sendeeRes int
		var createtime, content string
		err = rows.Scan(&senderRes, &sendeeRes, &createtime, &content)
		if err != nil {
			log.Print(err)
			return "Unknown Error"
		}
		theSenderName := GetUserName(senderRes)
		if theSenderName == "@" {
			return "Unknown Error"
		} else if theSenderName == "!" {
			return "G103"
		}
		theSendeeName := GetUserName(sendeeRes)
		if theSendeeName == "@" {
			return "Unknown Error"
		} else if theSendeeName == "!" {
			return "G103"
		}
		result.Messages = append(result.Messages,
			MessageInfo{
				Sender:     theSenderName,
				Sendee:     theSendeeName,
				Createtime: createtime,
				Content:    content,
			})
	}
	if !isNotEmpty {
		return "G104"
	}
	err = rows.Err()
	if err != nil {
		log.Print(err)
		return "Unknown Error"
	}

	res, err := json.Marshal(result)
	if err != nil {
		log.Print(err)
		return "G105"
	}
	return string(res)
}

type AllPeople struct {
	MesList []OnePeople `json:"messages"`
}

type OnePeople struct {
	User     string `json:"user"`
	LastTime string `json:"last_time"`
}

// G101 : Cannot connect to DB
// G102 : SQL statement error
// G103 : SQL exe error
// G104 : empty
// G105 : JSON error
// G106 : User error

func GetMesList(userid int) string {
	db, err := sql.Open(sqlDriver, userDataPath)
	if err != nil {
		log.Print(err)
		return "G101"
	}
	defer func() {
		_ = db.Close()
	}()
	stmt1, err := db.Prepare(`
SELECT user,MAX(createtime) AS utime
FROM (
         SELECT sendee AS user, createtime
         FROM message
         WHERE sender = ?

         UNION ALL

         SELECT sender AS user, createtime
         FROM message
         WHERE sendee = ?
     )
GROUP BY user
ORDER BY utime DESC;
`)
	if err != nil {
		log.Print(err)
		return "G102"
	}
	defer func() {
		_ = stmt1.Close()
	}()
	rows, err := stmt1.Query(userid, userid)
	if err == sql.ErrNoRows {
		return "G104"
	} else if err != nil {
		log.Print(err)
		return "G103"
	}
	var result AllPeople
	defer rows.Close()
	isNotEmpty := false
	for rows.Next() {
		isNotEmpty = true
		var uid int
		var lasttime string
		err = rows.Scan(&uid, &lasttime)
		if err != nil {
			log.Print(err)
			return "Unknown Error"
		}
		theSenderName := GetUserName(uid)
		if theSenderName == "@" {
			return "Unknown Error"
		} else if theSenderName == "!" {
			return "G103"
		}
		result.MesList = append(result.MesList,
			OnePeople{
				User:     theSenderName,
				LastTime: lasttime,
			})
	}
	if !isNotEmpty {
		return "G104"
	}
	err = rows.Err()
	if err != nil {
		log.Print(err)
		return "Unknown Error"
	}
	res, err := json.Marshal(result)
	if err != nil {
		log.Print(err)
		return "G105"
	}
	return string(res)
}
