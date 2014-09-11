package views

import "net/http"
import "narwhal/server"
import "narwhal/database"

func GetUser(r *http.Request) *database.User {
	session, _ := store.Get(r, "session")
	user, _ := session.Values["user"].(*database.User)
	return user
}

func userNew(w http.ResponseWriter, r *http.Request) {
	msg := Message{}
	username := r.FormValue("username")
	if username == "" {
		msg.Errors = []string{ "Must have a username" }
		ServeJson(w, msg)
		return
	}

	var usr database.User
	database.DbMap.Where("name = ?", username).First(&usr)

	if usr.Id != 0 {
		msg.Errors = []string{ "Username taken" }
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
		return
	}
	session, _ := store.Get(r, "session")
	session.Values["user"] = user
	session.Save(r, w)

	msg.Content = "OK"
	ServeJson(w, msg)
}

func userLogin(w http.ResponseWriter, r *http.Request) {
	msg := Message{}
	username := r.FormValue("username")
	password := r.FormValue("password")

	user := database.GetUserByName(username)
	if user == nil || user.ValidatePassword(password) != nil {
		msg.Errors = []string{ "Incorrect password" }
		ServeJson(w, msg)
		return
	}

	session, _ := store.Get(r, "session")
	session.Values["user"] = user
	session.Save(r, w)

	msg.Content = "OK"
	ServeJson(w, msg)
}

func userLogout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	session.Options.MaxAge = -1
	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func userInfo(w http.ResponseWriter, r *http.Request) {
	msg := Message{}
	msg.Content = "here is ur info"
	ServeJson(w, msg)
}

func userProfile(w http.ResponseWriter, r *http.Request) {
	user := GetUser(r)
	context := make(map[string]interface{})
	context["user"] = user

	serveTemplate(w, "profile.html", context)
}

func userChangePass(w http.ResponseWriter, r *http.Request) {
	msg := Message{}
	pass1, pass2 := r.FormValue("password"), r.FormValue("password2")
	if (pass1 != pass2) || pass1 == "" {
		msg.Errors = []string{ "Passwords don't match" }
		ServeJson(w, msg)
		return

	}

	user := GetUser(r)
	user.ChangePass(pass1)

	msg.Content = "OK"
	ServeJson(w, msg)
}

func init() {
	server := server.GetServer()

	server.Route("/users", userNew).Methods("POST")
	server.Route("/users/password", userChangePass, authCookie).Methods("POST")
	server.Route("/users", userInfo, authCookie).Methods("GET")
	server.Route("/users/profile", userProfile, authCookie)
	server.Route("/users/login", userLogin).Methods("POST")
	server.Route("/users/logout", userLogout)
}
