package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"net/http"
	"github.com/video_server/api/dbops"
	"github.com/video_server/api/defs"
	"github.com/video_server/api/session"
)
/* 创建用户凭证 */
func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body) // 读取POST请求的Body
	ubody := &defs.UserCredential{}
	// 将 json 格式的请求体转换为用户凭证数据模型对象 ubody
	if err := json.Unmarshal(res, ubody); err != nil {
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}
	// 进行数据库操作：添加用户
	if err := dbops.AddUserCredential(ubody.Username, ubody.Pwd); err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	id := session.GenerateNewSessionId(ubody.Username)
	su := &defs.SignedUp{Success: true, SessionId: id}

	if resp, err := json.Marshal(su); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(w, string(resp), 201)
	}
}
/* 用户登陆 */
func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("user_name")
	io.WriteString(w, uname)
}

func GetAllVideos(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func GetVideo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func DeleteVideo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func ShowComments(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func PostComment(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func DeleteComment(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}
