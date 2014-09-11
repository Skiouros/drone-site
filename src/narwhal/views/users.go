package views

import "net/http"
import "narwhal/server"
import "narwhal/database"

func GetUser(r *http.Request) *database.User {
	session, _ := store.Get(r, "session")
	user, _ := session.Values["user"].(*database.User)
	return user
}

// TODO: Validate form
func userNew(w http.ResponseWriter, r *http.Request) {
	msg := Message{}
	username := r.FormValue("username")
	if username == "" {
		msg.Errors = []string{ "Must have a username" }
		ServeJson(w, msg)
		return
	}
	pass1, pass2 := r.FormValue("password"), r.FormValue("password2")
	if (pass1 != pass2) || pass1 == "" {
		msg.Errors = []string{ "Passwords don't match" }
		ServeJson(w, msg)
		return

	}

	user, err := database.CreateUser(username, pass1)
	if err != nil {
		msg.Errors = []string{ "Incorrect password" }
		ServeJson(w, msg)
		// http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}
	session, _ := store.Get(r, "session")
	session.Values["user"] = user
	session.Save(r, w)

	msg.Content = "OK"
	ServeJson(w, msg)
}

func userLogin(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	user := database.GetUser("name", username)
	session, _ := store.Get(r, "session")
	session.Values["user"] = user
	if user == nil || user.ValidatePassword(password) != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func userLogout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	session.Options.MaxAge = -1
	session.Save(r, w)

	msg := Message{}
	msg.Content = "Logged out"
	ServeJson(w, msg)
	// http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func init() {
	server := server.GetServer()

	server.Route("/users", userNew).Methods("POST")
	server.Route("/users/logout", userLogout)
}
