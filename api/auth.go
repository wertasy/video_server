package main

import (
	"net/http"
	"github.com/video_server/api/defs"
	"github.com/video_server/api/session"
)

var HEADER_FAILD_SESSION = "X-Session-Id"
var HEADER_FAILD_UNAME = "X-User-Name"

/* 验证用户的 Session 是否有效 */
func validateUserSession(r *http.Request) bool {
	sid := r.Header.Get(HEADER_FAILD_SESSION)
	if len(sid) == 0 {
		return false
	}

	uname, expired := session.IsSessionExpired(sid)
	if expired {
		return false
	}

	r.Header.Add(HEADER_FAILD_UNAME, uname)
	return true
}
/* 验证用户的身份是否有效 */
func ValidateUser(w http.ResponseWriter, r *http.Request) bool {
	uname := r.Header.Get(HEADER_FAILD_UNAME)
	if len(uname) == 0 {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return false
	}
	return true
}
