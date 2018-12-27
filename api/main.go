package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

/* httprouter 中间件 */
type midWareHandler struct {
	r *httprouter.Router
}

func NewMidWareHandler(r *httprouter.Router) http.Handler {
	m := midWareHandler{}
	m.r = r
	return m
}
/* 将 httprouter 的 ServeHTTP 进行封装，每次请求都要进行用户 Session 验证 */
func (m midWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 先验证会话
	validateUserSession(r)
	// 再提供服务
	m.r.ServeHTTP(w, r)
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.POST("/user", CreateUser)

	router.POST("/user/:user_name", Login)

	router.GET("/user/:user_name/videos", GetAllVideos)

	router.GET("/user/:user_name/videos/:vid", GetVideo)

	router.DELETE("/user/:user_name/videos/:vid", DeleteVideo)

	router.GET("/videos/:vid/comments", ShowComments)

	router.POST("/videos/:vid/comments", PostComment)

	router.DELETE("/videos/:vid/comments/:comment-id", DeleteComment)

	return router
}

func main() {
	r := RegisterHandlers()
	mh := NewMidWareHandler(r)
	http.ListenAndServe(":8000", mh)
}
