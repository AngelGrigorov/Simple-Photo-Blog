package loginPageHandler

import "net/http"

func logout(res http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, "session")
	delete(session.Values, "logged_in")
	session.Save(req, res)
	http.Redirect(res, req, "/", 302)
}
