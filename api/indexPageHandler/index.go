package indexPageHandler

import (
	"github.com/gorilla/sessions"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"net/http"
)

var tpl *template.Template
var store = sessions.NewCookieStore([]byte("something-very-secret"))

type Model struct {
	Pictures []string
	IsLogged bool
}

var Data Model

func index(res http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, "session")
	if session.Values["logged_in"] == true {
		Data.IsLogged = true
	} else {
		Data.IsLogged = false
	}
	Data.Pictures = getPicturePaths()

	tpl.ExecuteTemplate(res, "indexPageHandler.gohtml", Data)
}

func getPicturePaths() []string {
	files := []string{}
	filepath.Walk("./", func(path string, fi os.FileInfo, err error) error {

		if fi.IsDir() {
			return nil
		}
		// path separator replacement fix
		path = strings.Replace(path, string(filepath.Separator), "/", -1)
		if strings.HasSuffix(path, ".png") {
			files = append(files, path)
		}
		return nil
	})
	return files
}
