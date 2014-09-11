package views

import "os"
import "net/http"
import "html/template"

import "github.com/daemonl/go_sweetpl"
import "github.com/gorilla/sessions"

import _"narwhal/database"
import "narwhal/server"

var store = sessions.NewCookieStore([]byte("vanilla lemon cake"))
var templates = make(map[string]*template.Template)
var tpl sweetpl.SweeTpl

type Message struct {
	Content string
	Errors []string
}

func serveTemplate(w http.ResponseWriter, page string, context interface{}) {
	err := tpl.Render(w, page, context)
	if err != nil {
		http.Error(w, "Template not found", http.StatusInternalServerError)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	user := GetUser(r)

	context := make(map[string]interface{})
	context["user"] = user
	serveTemplate(w, "index.html", context)
}

func login(w http.ResponseWriter, r *http.Request) {
	user := GetUser(r)
	if user != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// count, err := database.DbMap.SelectInt("select count(*) from users")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// if count < 1 {
	// 	http.Redirect(w, r, "/register", http.StatusSeeOther)
	// 	return
	// }

	serveTemplate(w, "login.html", nil)
}

func init() {
	tpl = sweetpl.SweeTpl{
		Loader: &sweetpl.DirLoader{
			BasePath: "views/",
		},
		FuncMap: template.FuncMap{
		},
	}

	store.Options = &sessions.Options{
		Path: "/",
		MaxAge: 3600 * 8, // 8 Hours
	}
}

func StartServer() {
	server := server.GetServer()

	server.Route("/", index)

	println("Listening " + os.Getenv("PORT"))
	server.Start(os.Getenv("PORT"))
}

