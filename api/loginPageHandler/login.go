package loginPageHandler

import (
	"database/sql"
	"github.com/gorilla/sessions"
	"net/http"
)

var tpl *template.Template
var store = sessions.NewCookieStore([]byte("something-very-secret"))

type Model struct {
	Pictures []string
	IsLogged bool
}

var Data Model

func login(res http.ResponseWriter, req *http.Request) {
	db, err := sql.Open("sqlite3", "db/db.db")
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
	session, _ := store.Get(req, "session")
	if session.Values["logged_in"] == true {
		Data.IsLogged = true
	} else {
		Data.IsLogged = false
	}
	if req.Method == "POST" {
		password := req.FormValue("password")
		username := req.FormValue("userName")

		rows, err := db.Query("SELECT username, password FROM users")
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}

		var u string
		var p string

		for rows.Next() {
			err = rows.Scan(&u, &p)
			if err != nil {
				http.Error(res, err.Error(), 500)
				return
			}

			if username == u && password == p {
				session.Values["logged_in"] = true
				session.Save(req, res)
				http.Redirect(res, req, "/", 302)
			}
		}

		rows.Close() //good habit to close
	}
	Data.Pictures = getPicturePaths()
	tpl.ExecuteTemplate(res, "login.gohtml", Data)
}
