package controller

import (
	"fmt"
	//"crypto/rand"
	"html/template"
	"net/http"
	"borgdir.mediaArbeitsversion/app/model"
	"strconv"
	"os"
	"io"
//"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
	"encoding/base64"

)

type AdminClientStruct struct{
	Name string
	Image string
	Clients []AdminClient
}

type AdminClient struct{
	ID int
	Name string
	Image string
	Role string
	Status string
	ActiveUntil string
	Borrowed []model.Equipment
}

// Übersicht Clients
func AdminClients(w http.ResponseWriter, r *http.Request) {
  session, _ := store.Get(r, "session")

		// Check if user is authenticated
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			tmpl, _ := template.ParseFiles("templates/layout.tmpl","templates/footer2.tmpl", "templates/login.tmpl")
			tmpl.ExecuteTemplate(w, "layout", nil)
		} else {
			if session.Values["admin"].(bool){

			 name, _ := ((session.Values["name"].(string)))
			 admin, _:= model.GetClientByName(name)

			 var clients []AdminClient
			 userArray, _:=model.GetAllClients()
			 for _, user := range userArray{
			 	borrowed, _ := user.GetBorrowedEquipment()// Array mit Equipmentnamen holen
				 fmt.Println("User.ID , Equipment:")
				 fmt.Print(user.ID)
				 fmt.Println(borrowed)
				 adminClient := AdminClient{ID:user.ID, Name:user.Name, Image:user.Image,Role:user.Role, Status:user.Status, ActiveUntil:user.ActiveUntil, Borrowed:borrowed}
				 clients= append(clients,adminClient)
			 }

			 data :=AdminClientStruct{Name:name, Image:admin.Image, Clients:clients}

	     tmpl, _ := template.ParseFiles("templates/layout.tmpl","templates/footer2.tmpl", "templates/admin/clients.tmpl")
	     tmpl.ExecuteTemplate(w, "layout", data)
       }else{
				 tmpl, _ := template.ParseFiles("templates/layout.tmpl","templates/footer.tmpl", "templates/login.tmpl")
	 			tmpl.ExecuteTemplate(w, "layout", nil)
				 }}
}

func AdminEditClient(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); ok || auth {
		if session.Values["admin"].(bool) {

      keys,_:=r.URL.Query()["ID"]
      ID, _ :=strconv.Atoi(keys[0])
      data, _ :=model.GetClientByID(ID)

      tmpl, _ := template.ParseFiles("templates/layout.tmpl","templates/footer2.tmpl","templates/admin/edit-client.tmpl")
	    tmpl.ExecuteTemplate(w, "layout", data)
			} else {
				tmpl, _ := template.ParseFiles("templates/layout.tmpl", "templates/footer.tmpl", "templates/login.tmpl")
				tmpl.ExecuteTemplate(w, "layout", nil)
			}
			} else {
			tmpl, _ := template.ParseFiles("templates/layout.tmpl", "templates/footer.tmpl", "templates/login.tmpl")
			tmpl.ExecuteTemplate(w, "layout", nil)
			}
			}

//Für Admin
func AdminEditClientButton(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); ok || auth {
		if session.Values["admin"].(bool) {

			if r.Method == "POST" {
				ID, _ := strconv.Atoi(r.FormValue("ID"))
				name := r.FormValue("name")
				email := r.FormValue("email")
				password:=r.FormValue("password")

				user := model.User{ID:ID, Name: name, Email:email, Password:password}
				user.EditClient()
			}

			ID, _ := strconv.Atoi(r.FormValue("ID"))
			data, _ :=model.GetClientByID(ID)
			tmpl, _ := template.ParseFiles("templates/layout.tmpl","templates/footer.tmpl","templates/admin/edit-client.tmpl")
			tmpl.ExecuteTemplate(w, "layout", data)
		} else {
			tmpl, _ := template.ParseFiles("templates/layout.tmpl", "templates/footer.tmpl", "templates/login.tmpl")
			tmpl.ExecuteTemplate(w, "layout", nil)
		}
	} else {
		tmpl, _ := template.ParseFiles("templates/layout.tmpl", "templates/footer.tmpl", "templates/login.tmpl")
		tmpl.ExecuteTemplate(w, "layout", nil)
	}
}
//Für Clienten
func EditClientButton(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); ok || auth {
	//	if r.Method == "POST" {
			keys,_:=r.URL.Query()["ID"]
			ID, _ :=strconv.Atoi(keys[0])
			name := r.FormValue("name")
			email := r.FormValue("email")
			oldPassword:=r.FormValue("oldPassword")
			newPassword:=r.FormValue("newPassword")
			user, _ := model.GetClientByName(name)
			fmt.Println("EditClientButton, authenticated")
			image :=user.Image

			fmt.Print("user")
			fmt.Println(user)

				 // Bild-Upload
				 _ ,fileheader, err:= r.FormFile("file")

					if err != nil { fmt.Print("Error beim Auslesen der Bilddatei! ") }

				src, error := fileheader.Open()
					if error != nil { fmt.Print("Error beim Auslesen des Bild-Headers! ") }
					defer src.Close()

				dst, err := os.Create("./static/images/" + fileheader.Filename)
					if err != nil {fmt.Print("Error beim Erstellen der Bild-Datei! ") }
					image = ("./static/images/" + fileheader.Filename)
					defer dst.Close()

				if _ ,err = io.Copy(dst, src)
				err !=nil {fmt.Print("Error beim Kopieren der Bild-Datei! ")}


			//ohne Passwortänderung
			if newPassword == ""{
				fmt.Println("newPassword == ")
				fmt.Println(email)
				user := model.User{ Name:name, Email:email, ID:ID, Image:image}
				user.EditClient() //nur Name und Mail werden bearbeitet
			//mit Passwortänderung
			}else{
				fmt.Println("newPassword != ")
				// decode base64 String to []byte
				passwordDB, _ := base64.StdEncoding.DecodeString(user.Password)
				err := bcrypt.CompareHashAndPassword(passwordDB,[]byte(oldPassword))
				if err == nil{
					password :=[]byte(newPassword)
	 				hashedPwd, err := bcrypt.GenerateFromPassword(password, 14)
	 				if err!= nil{
	 					 panic(err)
	 				 }
	 			//hashed password to string
	 			b64HashPwd := base64.StdEncoding.EncodeToString(hashedPwd)
				user := model.User{ Name: name, Email:email, Password: b64HashPwd, ID:ID, Image:image}
				user.EditClientPw()
			}
		}
	//	}//Post
			//ID, _ := strconv.Atoi(r.FormValue("ID"))
			data, _ :=model.GetClientByID(ID)
			tmpl, _ := template.ParseFiles("templates/layout.tmpl","templates/footer2.tmpl","templates/profil.tmpl")
			tmpl.ExecuteTemplate(w, "layout", data)
	} else {
		tmpl, _ := template.ParseFiles("templates/layout.tmpl", "templates/footer.tmpl", "templates/login.tmpl")
		tmpl.ExecuteTemplate(w, "layout", nil)
	}//falls nicht authentifiziert
}//Funktion

func Profil(w http.ResponseWriter, r *http.Request) {
  session, _ := store.Get(r, "session")

		// Check if user is authenticated
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			tmpl, _ := template.ParseFiles("templates/layout.tmpl","templates/footer.tmpl", "templates/login.tmpl")
			tmpl.ExecuteTemplate(w, "layout", nil)
		} else {
			name :=session.Values["name"].(string)
			data, _:=model.GetClientByName(name)
	     tmpl, _ := template.ParseFiles("templates/layout.tmpl", "templates/footer2.tmpl", "templates/profil.tmpl")
	     tmpl.ExecuteTemplate(w, "layout", data)
     }
}

func Register(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/layout.tmpl", "templates/footer.tmpl", "templates/register.tmpl")
	tmpl.ExecuteTemplate(w, "layout", nil)
}

// AddClient controller
func AddClient(w http.ResponseWriter, r *http.Request) {
      if r.Method == "POST" {
	       name := r.FormValue("name")
	       email:=r.FormValue("email")
				 password :=[]byte(r.FormValue("password"))
				 hashedPwd, err := bcrypt.GenerateFromPassword(password, 14)
				 if err!= nil{
					 panic(err)
				 }
				 //hashed password to string
				 b64HashPwd := base64.StdEncoding.EncodeToString(hashedPwd)

		     user := model.User{ Name: name, Email:email, Password: b64HashPwd}
		     user.AddClient()

        	tmpl, _ := template.ParseFiles("templates/layout.tmpl","templates/footer.tmpl", "templates/login.tmpl")
          tmpl.ExecuteTemplate(w, "layout", nil)
             }
}

// DeleteClient controller
func DeleteClient(w http.ResponseWriter, r *http.Request) {
	keys,_:=r.URL.Query()["ID"]
	ID, _ :=strconv.Atoi(keys[0])

	model.DeleteClient(ID)
	http.Redirect(w, r, "/logout", http.StatusFound)
}

// LockUser controller
func AdminLockClient(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); ok || auth {
		if session.Values["admin"].(bool) {

				fmt.Println("AdminLockClient session")
				keys,_:=r.URL.Query()["ID"]
				fmt.Print("keys: ")
				fmt.Println(keys)
				ID, _ :=strconv.Atoi(keys[0])
					fmt.Print("ID: ")
					fmt.Print(ID)
					user, _ := model.GetClientByID(ID)
					fmt.Println(user)
					user.LockClient()
					StringID:= strconv.Itoa(ID)
						http.Redirect(w, r, "/admin/edit-client?ID=" + StringID , http.StatusFound)
		} else {
			tmpl, _ := template.ParseFiles("templates/layout.tmpl", "templates/footer.tmpl", "templates/login.tmpl")
			tmpl.ExecuteTemplate(w, "layout", nil)
		}
		} else {
		tmpl, _ := template.ParseFiles("templates/layout.tmpl", "templates/footer.tmpl", "templates/login.tmpl")
		tmpl.ExecuteTemplate(w, "layout", nil)
		}
		}

// UnlockUser controller
func AdminUnlockClient(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
		// Check if user is authenticated
		if auth, ok := session.Values["authenticated"].(bool); ok || auth {
			if session.Values["admin"].(bool) {
				fmt.Println("AdminUnlockClient session")
				keys,_:=r.URL.Query()["ID"]
				fmt.Print("keys: ")
				fmt.Println(keys)
				ID, _ :=strconv.Atoi(keys[0])
				fmt.Print("ID: ")
				fmt.Print(ID)
				user, _ := model.GetClientByID(ID)
				fmt.Println(user)
				user.UnlockClient()
				StringID:= strconv.Itoa(ID)
				http.Redirect(w, r, "/admin/edit-client?ID=" + StringID , http.StatusFound)
				} else {
					tmpl, _ := template.ParseFiles("templates/layout.tmpl", "templates/footer.tmpl", "templates/login.tmpl")
					tmpl.ExecuteTemplate(w, "layout", nil)
				}
				} else {
				tmpl, _ := template.ParseFiles("templates/layout.tmpl", "templates/footer.tmpl", "templates/login.tmpl")
				tmpl.ExecuteTemplate(w, "layout", nil)
				}
}
