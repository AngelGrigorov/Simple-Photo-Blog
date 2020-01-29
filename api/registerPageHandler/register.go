package registerPageHandler

import (
	"database/sql"
	"net/http"
)

func register(res http.ResponseWriter, req *http.Request) {
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
		stmt, err := db.Prepare("INSERT INTO users (username, password) values(?,?)")
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		stmt.Exec(username, password)

		session.Save(req, res)
		http.Redirect(res, req, "/login", 302)
	}
	Data.Pictures = getPicturePaths()
	tpl.ExecuteTemplate(res, "register.gohtml", Data)
}
