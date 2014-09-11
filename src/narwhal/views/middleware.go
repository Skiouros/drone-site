package views

import "net/http"

import "narwhal/database"

func authCookie(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	session, _ := store.Get(r, "session")
	user, err := session.Values["user"].(*database.User)

	if err != true || user == nil {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}

	next(w, r)
}
