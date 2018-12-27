package session

import (
	"time"
	"sync"
	"github.com/video_server/api/defs"
	"github.com/video_server/api/dbops"
	"github.com/video_server/api/utils"
)


var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func nowInMili() int64 {
	return time.Now().UnixNano()/1000000
}

func deleteExpiredSession(sid string) {
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}

func LoadSessionsFromDB() {
	r, err := dbops.RetriveAllSessions()
	if err != nil {
		return
	}

	r.Range(func(k, v interface{}) bool {
		ss := v.(*defs.SimpleSession)
		sessionMap.Store(k, ss)
		return true
	})
}
/* 生成一个新的 Session ID，Session ID 是一个 string 类型的 UUID */
func GenerateNewSessionId(uname string) string {
	id, _ := utils.NewUUID()
	ctime := nowInMili()
	ttl := ctime + 30 * 60 * 1000 // Serverside session valid time: 30 min
	ss := &defs.SimpleSession{Username: uname, TTL:ttl}
	sessionMap.Store(id, ss)
	dbops.InsertSession(id, ttl, uname)

	return id
}
/*
  通过 session id 判断一个 session 是否已经过期。
  如果没有过期将返回对应的用户名和 true，
  否则返回空字符串和 false
*/
func IsSessionExpired(sid string) (string, bool) {
	ss, ok := sessionMap.Load(sid)
	if ok {
		ctime := nowInMili()
		if ss.(*defs.SimpleSession).TTL < ctime {
			deleteExpiredSession(sid)
			return "", true
		}
		return ss.(*defs.SimpleSession).Username, false
	}
	return "", true
}