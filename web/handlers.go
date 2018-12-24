package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type HomePage struct {
	Title string
}

type UserPage struct {
	Name string
}

func indexHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//cname, err1 := r.Cookie("username")
	//sid, err2 := r.Cookie("session")

	//if err1 != nil || err2 != nil {
	p := &HomePage{Title: "WERTASY"}
	tmpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		log.Printf("Parsing template index.html error: %s", err)
		return
	}

	tmpl.Execute(w, p)
	//}

	// 这里需要判断是否存在和是否匹配
	//if len(cname.Value) != 0 && len(sid.Value) != 0 {
	//	http.Redirect(w, r, "/userhome", http.StatusFound)
	//}
}

func loginHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("%s\n", string(res))
}

func userHomeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	/*

		cname, err1 := r.Cookie("username")
		_, err2 := r.Cookie("session")

		if err1 != nil || err2 != nil {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		fname := r.FormValue("username")

		var p *UserPage
		if len(cname.Value) != 0 {
			p = &UserPage{ Name: cname.Value }
		} else if len(fname) != 0 {
			p = &UserPage{ Name: fname }
		}

	*/

	p := &UserPage{Name: "wertasy"}
	tmpl, err := template.ParseFiles("./templates/userhome.html")
	if err != nil {
		log.Printf("Parsing userhome.html error: %s", err)
		return
	}

	tmpl.Execute(w, p)
}

func apiHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Method != http.MethodPost {
		re, _ := json.Marshal(ErrorRequestNotRecognized)
		io.WriteString(w, string(re))
		return
	}

	res, _ := ioutil.ReadAll(r.Body)
	apibody := &ApiBody{}
	if err := json.Unmarshal(res, apibody); err != nil {
		re, _ := json.Marshal(ErrorRequestBodyParseFailed)
		io.WriteString(w, string(re))
		return
	}

	request(apibody, w, r)
	defer r.Body.Close()
}
