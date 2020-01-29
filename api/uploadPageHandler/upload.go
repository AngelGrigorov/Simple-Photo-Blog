package uploadPageHandler

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

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

func upload(res http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, "session")
	if session.Values["logged_in"] == true {
		Data.IsLogged = true
	} else {
		Data.IsLogged = false
		http.Redirect(res, req, "https://localhost:8070/login", 302)
	}
	Data.Pictures = getPicturePaths()

	if req.Method == "POST" {
		src, hdr, err := req.FormFile("my-file")
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		defer src.Close()

		path := "/home/angel/Developer/GoLangProjects/src/Simple-Photo-Blog/assets/imgs/"
		dst, err := os.Create(path + hdr.Filename)
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		defer dst.Close()

		io.Copy(dst, src)
		Data.Pictures = getPicturePaths()
		http.Redirect(res, req, "/", 302)
	}
	tpl.ExecuteTemplate(res, "upload-file.gohtml", Data)
}
