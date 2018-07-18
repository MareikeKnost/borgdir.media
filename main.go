package main

import (
	"log"
	"net/http"
	"borgdir.mediaArbeitsversion/app/controller"
	"github.com/gorilla/context"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	//controller
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/admin", controller.Admin) //Auswahlseite
	http.HandleFunc("/loggingIn", controller.LoggingIn)
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/logout", controller.Logout)

	//Equipment
	http.HandleFunc("/equipment", controller.Equipmenta) //Übersicht Client
	http.HandleFunc("/admin/equipment", controller.AdminEquipment)// Übersicht Admin
	http.HandleFunc("/admin/add", controller.AdminAdd) //PLUS-BUTTON: Eingabemaske aufrufen
	http.HandleFunc("/add-equipment", controller.AddEquipment)// Daten insert
	http.HandleFunc("/admin/edit", controller.EditEquipment) //STIFT_BUTTON: Eingabemaske zu bestimmtem Gerät aufrufen
	http.HandleFunc("/admin/edit-equipment", controller.AdminEditEquipment) // BUTTON_ÄNDERUNG_SPEICHERN: Daten updaten
	http.HandleFunc("/admin/delete-equipment", controller.DeleteEquipment)//BIN_BUTTON
	http.HandleFunc("/my-equipment", controller.MyEquipment)
	http.HandleFunc("/cart", controller.Cart)
	http.HandleFunc("/add-to-cart", controller.AddToCart)
	http.HandleFunc("/cart-quantity", controller.CartQuantity)
	http.HandleFunc("/borrow", controller.Borrow)
	http.HandleFunc("/deleteFromCart", controller.DeleteFromCart)
	http.HandleFunc("/deleteReserved", controller.DeleteReserve)
	http.HandleFunc("/extend", controller.Extend)

	//user
	http.HandleFunc("/admin/clients", controller.AdminClients)//Übersicht Clients
	http.HandleFunc("/register", controller.Register)//Eingabemaske Client
	http.HandleFunc("/add-client", controller.AddClient)//Daten insert
	http.HandleFunc("/profil", controller.Profil)//Eingabemaske für User
	http.HandleFunc("/admin/edit-client", controller.AdminEditClient)//Eingabemaske Admin
	http.HandleFunc("/admin/edit-client-button", controller.AdminEditClientButton)//BUTTON_ÄNDERUNG_SPEICHERN
	http.HandleFunc("/admin/lock-client", controller.AdminLockClient)//Client sperren
	http.HandleFunc("/admin/unlock-client", controller.AdminUnlockClient)//Client entsperren
	http.HandleFunc("/reserve", controller.Reserve)//Client entsperren
	http.HandleFunc("/delete-client", controller.DeleteClient)//Client löschen
	http.HandleFunc("/edit-client-button", controller.EditClientButton)//Client bearbeiten

	log.Println("Listening...")

	//against memory leak use context.clearHandler, instead of gorilla/mux
	http.ListenAndServe(":80", context.ClearHandler(http.DefaultServeMux))
}
