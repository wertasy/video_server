package main

import (
	"encoding/json"
	"io"
	"net/http"
	"github.com/video_server/api/defs"
)
/* 发送一个错误回复 */
func sendErrorResponse(w http.ResponseWriter, errResp defs.ErrResponse) {
	w.WriteHeader(errResp.HttpSC)

	resStr, _ := json.Marshal(&errResp.Error)
	io.WriteString(w, string(resStr))
}
/* 发送一个正常回复 */
func sendNormalResponse(w http.ResponseWriter, resp string, sc int) {
	w.WriteHeader(sc)

}
