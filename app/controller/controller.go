package controller

import (
	"crypto/rand"
	"html/template"
	"net/http"
	"borgdir.mediaArbeitsversion/app/model"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
	"encoding/base64"
	"fmt"
	"encoding/gob"
)

var tmpl *template.Template
var store *sessions.CookieStore
var equipments []model.Equipment

// Is executed automatically on package load
func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.tmpl"))

	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key := make([]byte, 32)
	_, err:= rand.Read(key)
	if err!= nil {
		panic(err)
	}
	store = sessions.NewCookieStore(key)
	//register the type []Equipment
	gob.Register([]model.Equipment{})
}

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/layout.tmpl", "templates/footer2.tmpl", "templates/index.tmpl")
	tmpl.ExecuteTemplate(w, "layout", nil)
}


func Login(w http.ResponseWriter, r *http.Request) {
	//get a session
	session, _ := store.Get(r, "session")

		// Check if user is authenticated
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			tmpl, _ := template.ParseFiles("templates/layout.tmpl","templates/footer.tmpl", "templates/login.tmpl")
			tmpl.ExecuteTemplate(w, "layout", nil)
		} else {
			tmpl, _ := template.ParseFiles("templates/layout.tmpl","templates/footer2.tmpl", "templates/my-equipment.tmpl")
			tmpl.ExecuteTemplate(w, "layout", nil)
		}
}

func Admin(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")


		// Check if user is authenticated
		if auth, ok := session.Values["authenticated"].(bool); ok || auth {
			if session.Values["admin"].(bool) {
				user, _ :=model.GetClientByName(session.Values["name"].(string))
				tmpl, _ := template.ParseFiles("templates/layout.tmpl", "templates/footer.tmpl","templates/admin.tmpl")
				tmpl.ExecuteTemplate(w, "layout", user)
			} else {
				tmpl, _ := template.ParseFiles("templates/layout.tmpl", "templates/footer.tmpl", "templates/login.tmpl")
				tmpl.ExecuteTemplate(w, "layout", nil)
			}
		} else {
			tmpl, _ := template.ParseFiles("templates/layout.tmpl", "templates/footer.tmpl", "templates/login.tmpl")
			tmpl.ExecuteTemplate(w, "layout", nil)
		}
}

func LoggingIn(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	name := r.FormValue("name")
	password := r.FormValue("password")

	// Authentication
	user, _ := model.GetClientByName(name)

	// decode base64 String to []byte
	passwordDB, _ := base64.StdEncoding.DecodeString(user.Password)
	err := bcrypt.CompareHashAndPassword(passwordDB,[]byte(password))

	if err == nil && user.Status != "gesperrt" {
		user, _ := model.GetClientByName(name)
		// Set user as authenticated
		session.Values["authenticated"] = true
		session.Values["name"] = name
		if user.Role =="admin" {
			session.Values["admin"] = true
		} else {
			session.Values["admin"] = false
		}


		session.Save(r, w)

		if user.Role =="admin" {
			http.Redirect(w, r, "/admin", http.StatusFound)
		} else {
			http.Redirect(w, r, "/my-equipment", http.StatusFound)}

		} else {
			tmpl, _ := template.ParseFiles("templates/layout.tmpl", "templates/footer.tmpl", "templates/login.tmpl")
			tmpl.ExecuteTemplate(w, "layout", nil)
			fmt.Println("Nutzer gesperrt oder Passwort falsch!")
	}}

// Logout controller
func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")

	//Revoke users authentication
	session.Values["authenticated"] = false
	session.Values["name"] = ""
	session.Values["equipments"] = nil
	session.Save(r, w)

	tmpl, _ := template.ParseFiles("templates/layout.tmpl", "templates/footer.tmpl", "templates/login.tmpl")
	tmpl.ExecuteTemplate(w, "layout", nil)
}
