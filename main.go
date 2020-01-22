package main

import (
	"net/http"
	"github.com/gorilla/sessions"
	"html/template"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var err error
var tpl *template.Template
var store = sessions.NewCookieStore([]byte("something-very-secret"))

type Model struct {
	Pictures []string
}

func main()  {
	tpl, err = tpl.ParseGlob("assets/templates/*.gohtml")
	if err != nil{
		log.Fatalln(err)
	}

	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/logout", logout)
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServeTLS(":8070", "cert.pem", "key.pem", nil)
}

func index(res http.ResponseWriter, req *http.Request)  {
	tpl.ExecuteTemplate(res, "index.gohtml", Data)
}

func login(res http.ResponseWriter, req *http.Request){
	session, _ := store.Get(req, "session")

	if req.Method == "POST" {
		password := req.FormValue("password")
			if password == "secret" {
				session.Values["logged_in"] = true
			} else {
				http.Error(res, "invalid credentials", 401)
				return
			}
			session.Save(req, res)
			http.Redirect(res, req, "/", 302)
	}
	tpl.ExecuteTemplate(res, "login.gohtml", nil)
}

func logout(res http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, "session")
	delete(session.Values, "logged_in")
	session.Save(req, res)
	http.Redirect(res, req, "/", 302)
}
var Data Model = getPicturePaths()

func getPicturePaths() Model  {
	files := []string{}
	filepath.Walk("./", func(path string, fi os.FileInfo, err error) error {

		if fi.IsDir() {
			return nil
		}

		path = strings.Replace(path, "//", "/", -1)

		if strings.HasSuffix(path, ".jpg") {
			files = append(files, path)
		}
		return nil
	})
	return Model{Pictures:files}
}

func upload(res http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, "session")
	if session.Values["logged_in"] == false || session.Values["logged_in"] == nil {
		http.Redirect(res, req, "https://localhost:8070/login", 302)
	}

	if req.Method == "POST" {
		src, hdr, err := req.FormFile("my-file")
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		defer src.Close()

		path := "/Users/8770W/Desktop/Simple-Photo-Blog/assets/imgs/"
		dst, err := os.Create(path + hdr.Filename)
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		defer dst.Close()

		io.Copy(dst,src)
		Data = getPicturePaths()
		http.Redirect(res,req,"/", 302)
	}
	tpl.ExecuteTemplate(res, "upload-file.gohtml", nil)
}