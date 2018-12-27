package dbops

import (
	"github.com/video_server/api/defs"
	"strconv"
	"log"
	"sync"
)

func InsertSession(sid string, ttl int64, uname string) error {
	ttlstr := strconv.FormatInt(ttl, 10)

	stmt, err := dbConn.Prepare("INSERT INTO sesstions (session_id, ttl, login_name) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(sid, ttlstr, uname)
	if err != nil {
		return err
	}

	return nil
}

func RetrieveSession(sid string) (*defs.SimpleSession, error) {
	ss := &defs.SimpleSession{}
	stmt, err := dbConn.Prepare("SELECT ttl, login_name FROM sessions WHERE session_id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var ttl, uname string
	err = stmt.QueryRow(sid).Scan(&ttl, &uname)
	if err != nil {
		return nil, err
	}

	if res, err := strconv.ParseInt(ttl, 10, 64); err == nil {
		ss.TTL = res
		ss.Username = uname
	} else {
		return nil, err
	}

	return ss, nil
}

func RetriveAllSessions() (*sync.Map, error) {
	m := &sync.Map{}
	stmt, err := dbConn.Prepare("SELECT * FROM sessions")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	for rows.Next() {
		var id, ttlstr, login_name string
		err = rows.Scan(&id, &ttlstr, &login_name)
		if err != nil {
			log.Printf("retrive sessions error: %s", err)
			break
		}

		ttl, err := strconv.ParseInt(ttlstr, 10, 64)
		if err != nil {
			ss := &defs.SimpleSession{Username: login_name, TTL: ttl}
			m.Store(id, ss)
			log.Printf("sessions id: %s, ttl: %d", id, ss.TTL)
		}
	}

	return m, nil
}

func DeleteSession(sid string) error {
	stmt, err := dbConn.Prepare("DELETE FROM sessions WHERE session_id = ?")
	if err != nil {
		log.Printf("Delete Session Error: %s", err)
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(sid); err !=  nil {
		return err
	}

	return nil
}