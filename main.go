package main

import (
	index "Simple-Photo-Blog/api/indexPageHandler"
	"github.com/gorilla/sessions"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"log"
	"net/http"
)

var err error
var tpl *template.Template
var store = sessions.NewCookieStore([]byte("something-very-secret"))

type Model struct {
	Pictures []string
	IsLogged bool
}

var Data Model

func main() {
	tpl, err = tpl.ParseGlob("assets/templates/*.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/register", register)
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServeTLS(":8070", "cert.pem", "key.pem", nil)
}

//
//func register(res http.ResponseWriter, req *http.Request) {
//	db, err := sql.Open("sqlite3", "db/db.db")
//	if err != nil {
//		http.Error(res, err.Error(), 500)
//		return
//	}
//	session, _ := store.Get(req, "session")
//	if session.Values["logged_in"] == true {
//		Data.IsLogged = true
//	} else {
//		Data.IsLogged = false
//	}
//	if req.Method == "POST" {
//		password := req.FormValue("password")
//		username := req.FormValue("userName")
//		stmt, err := db.Prepare("INSERT INTO users (username, password) values(?,?)")
//		if err != nil {
//			http.Error(res, err.Error(), 500)
//			return
//		}
//		stmt.Exec(username, password)
//
//		session.Save(req, res)
//		http.Redirect(res, req, "/login", 302)
//	}
//	Data.Pictures = getPicturePaths()
//	tpl.ExecuteTemplate(res, "register.gohtml", Data)
//}

//func indexPageHandler(res http.ResponseWriter, req *http.Request) {
//	session, _ := store.Get(req, "session")
//	if session.Values["logged_in"] == true {
//		Data.IsLogged = true
//	} else {
//		Data.IsLogged = false
//	}
//	Data.Pictures = getPicturePaths()
//
//	tpl.ExecuteTemplate(res, "indexPageHandler.gohtml", Data)
//}

//func login(res http.ResponseWriter, req *http.Request) {
//	db, err := sql.Open("sqlite3", "db/db.db")
//	if err != nil {
//		http.Error(res, err.Error(), 500)
//		return
//	}
//	session, _ := store.Get(req, "session")
//	if session.Values["logged_in"] == true {
//		Data.IsLogged = true
//	} else {
//		Data.IsLogged = false
//	}
//	if req.Method == "POST" {
//		password := req.FormValue("password")
//		username := req.FormValue("userName")
//
//		rows, err := db.Query("SELECT username, password FROM users")
//		if err != nil {
//			http.Error(res, err.Error(), 500)
//			return
//		}
//
//		var u string
//		var p string
//
//		for rows.Next() {
//			err = rows.Scan(&u, &p)
//			if err != nil {
//				http.Error(res, err.Error(), 500)
//				return
//			}
//
//			if username == u && password == p {
//				session.Values["logged_in"] = true
//				session.Save(req, res)
//				http.Redirect(res, req, "/", 302)
//			}
//		}
//
//		rows.Close() //good habit to close
//	}
//	Data.Pictures = getPicturePaths()
//	tpl.ExecuteTemplate(res, "login.gohtml", Data)
//}

//func logout(res http.ResponseWriter, req *http.Request) {
//	session, _ := store.Get(req, "session")
//	delete(session.Values, "logged_in")
//	session.Save(req, res)
//	http.Redirect(res, req, "/", 302)
//}

//func getPicturePaths() []string {
//	files := []string{}
//	filepath.Walk("./", func(path string, fi os.FileInfo, err error) error {
//
//		if fi.IsDir() {
//			return nil
//		}
//		// path separator replacement fix
//		path = strings.Replace(path, string(filepath.Separator), "/", -1)
//		if strings.HasSuffix(path, ".png") {
//			files = append(files, path)
//		}
//		return nil
//	})
//	return files
//}
//
//func upload(res http.ResponseWriter, req *http.Request) {
//	session, _ := store.Get(req, "session")
//	if session.Values["logged_in"] == true {
//		Data.IsLogged = true
//	} else {
//		Data.IsLogged = false
//		http.Redirect(res, req, "https://localhost:8070/login", 302)
//	}
//	Data.Pictures = getPicturePaths()
//
//	if req.Method == "POST" {
//		src, hdr, err := req.FormFile("my-file")
//		if err != nil {
//			http.Error(res, err.Error(), 500)
//			return
//		}
//		defer src.Close()
//
//		path := "/home/angel/Developer/GoLangProjects/src/Simple-Photo-Blog/assets/imgs/"
//		dst, err := os.Create(path + hdr.Filename)
//		if err != nil {
//			http.Error(res, err.Error(), 500)
//			return
//		}
//		defer dst.Close()
//
//		io.Copy(dst, src)
//		Data.Pictures = getPicturePaths()
//		http.Redirect(res, req, "/", 302)
//	}
//	tpl.ExecuteTemplate(res, "upload-file.gohtml", Data)
//}
