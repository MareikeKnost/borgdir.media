package controller

import (
	//"crypto/rand"
	"html/template"
	"net/http"
	"borgdir.mediaArbeitsversion/app/model"
	"strconv"
	"fmt"
	"time"
	"io"
	"os"
)

type MyEquipmentStruct struct {
Name string
Role bool
Image string
Entliehen []model.Borrowed
Vorgemerkt []model.Reserved
}

type EquipmentPlus struct {
ID int
Name string
Description string
Image string
Category string
Content string
Location string
Status string
Quantity int
Borrowed int
Reserved int
}

type EquipmentaStruct struct{
	Name string
	Role string
	Image string
	EquipmentListe []EquipmentPlus
}

type CartStruct struct{
	Name string
	Role string
	Image string
	Warenkorb []CartEquipment
}

type CartEquipment struct{
	Image string
	Name string
	ID int
	Description string
	Quantity int
	Date string
	MaxQuantity int
}

type AdminEquipments struct{
	Image string
	Name string
	ID int
	Location string
	Description string
	Quantity int
	Status string
	ReturnUntil string
}

type AdminEquipmentStruct struct{
	Name string
	Image string
	Equipment []AdminEquipments
}

type AdminOneEquipmentStruct struct{
	Name string
	Image string
	Equipment model.Equipment
}



func AdminAdd(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); ok || auth {
		if session.Values["admin"].(bool) {
			tmpl, _ := template.ParseFiles("templates/layout.tmpl", "templates/footer.tmpl","templates/admin/add.tmpl")
			tmpl.ExecuteTemplate(w, "layout", nil)
		} else {
			tmpl, _ := template.ParseFiles("templates/layout.tmpl", "templates/footer.tmpl", "templates/login.tmpl")
			tmpl.ExecuteTemplate(w, "layout", nil)
		}
	} else {
		tmpl, _ := template.ParseFiles("templates/layout.tmpl", "templates/footer.tmpl", "templates/login.tmpl")
		tmpl.ExecuteTemplate(w, "layout", nil)
	}
}

// AddEquipment controller BUTTON
func AddEquipment(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
				fmt.Print("---AddEquipment---")
	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); ok || auth {
		if session.Values["admin"].(bool) {

	if r.Method == "POST" {
		name := r.FormValue("name")
		location := r.FormValue("location")
		description:=r.FormValue("description")
		content:=r.FormValue("content")
		category:=r.FormValue("category")
		quantity,_:=strconv.Atoi(r.FormValue("quantity"))
		fmt.Println("Auslesen der Form-Values")

// Bild-Upload
		 _,fileheader, err:= r.FormFile("file")

			if err != nil { fmt.Print("Error beim Auslesen der Bilddatei! ") }

		src, error := fileheader.Open()
			if error != nil { fmt.Print("Error beim Auslesen des Bild-Headers! ") }
			defer src.Close()

		dst, err := os.Create("./static/images/" + fileheader.Filename)
			if err != nil {fmt.Print("Error beim Erstellen der Bild-Datei! ") }
			image := ("./static/images/" + fileheader.Filename)
			defer dst.Close()

		if _ ,err = io.Copy(dst, src)
		err !=nil {fmt.Print("Error beim Kopieren der Bild-Datei! ")}

		if len(name) != 0 {
			equipment := model.Equipment{ Name:name, Description:description, Image:image, Category:category, Content: content, Location:location, Quantity:quantity, MaxQuantity:quantity}
			 fmt.Println("equipment in AddEquipment")
			 fmt.Println(equipment)

			equipment.AddEquipment()
		}
	}
	tmpl, _ := template.ParseFiles("templates/layout.tmpl","templates/footer.tmpl", "templates/admin/add.tmpl")
	tmpl.ExecuteTemplate(w, "layout", nil)
	} else {
		tmpl, _ := template.ParseFiles("templates/layout.tmpl", "templates/footer.tmpl", "templates/login.tmpl")
		tmpl.ExecuteTemplate(w, "layout", nil)
	}
	} else {
	tmpl, _ := template.ParseFiles("templates/layout.tmpl", "templates/footer.tmpl", "templates/login.tmpl")
	tmpl.ExecuteTemplate(w, "layout", nil)
	}
	}


//Button bei Equipment-Übersicht
func EditEquipment(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); ok || auth {
		if session.Values["admin"].(bool) {

	if r.Method == "POST" {
		name := r.FormValue("name")
		location := r.FormValue("location")
		description:=r.FormValue("description")
		content:=r.FormValue("content")
		category:=r.FormValue("category")
		quantity,_:=strconv.Atoi(r.FormValue("quantity"))

			equipment := model.Equipment{ Name: name, Content: content, Location: location, Description: description, Quantity: quantity, Category:category}
			equipment.EditEquipment()

	}

		keys,_:=r.URL.Query()["ID"]
		ID, _ :=strconv.Atoi(keys[0])


		data, _ := model.GetEquipmentByID(ID)
		tmpl, _ := template.ParseFiles("templates/layout.tmpl","templates/footer.tmpl","templates/admin/edit.tmpl")
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

// AdminEditEquipment-BUTTON controller
func AdminEditEquipment(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); ok || auth {
		if session.Values["admin"].(bool) {

	if r.Method == "POST" {
		ID, _ := strconv.Atoi(r.FormValue("ID"))
		name := r.FormValue("name")
		location := r.FormValue("location")
		description:=r.FormValue("description")
		content:=r.FormValue("content")
		category:=r.FormValue("category")
		quantity,_:=strconv.Atoi(r.FormValue("quantity"))

		equipment := model.Equipment{ID:ID, Name: name, Content: content, Location: location, Description: description, Quantity: quantity, Category:category}
		equipment.EditEquipment()
	}

	ID, _ := strconv.Atoi(r.FormValue("ID"))
	data, _ :=model.GetEquipmentByID(ID)
	tmpl, _ := template.ParseFiles("templates/layout.tmpl","templates/footer.tmpl","templates/admin/edit.tmpl")
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

// DeleteEquipment controller
func DeleteEquipment(w http.ResponseWriter, r *http.Request) {
	ID, _ := strconv.Atoi(r.FormValue("ID"))
	equipment, _ := model.GetEquipmentByID(ID)
	equipment.DeleteEquipment()

	data, _:=model.GetAllEquipment()
	tmpl, _ := template.ParseFiles("templates/layout.tmpl","templates/footer.tmpl", "templates/admin/equipment.tmpl")
	tmpl.ExecuteTemplate(w, "layout", data)
}

func AdminEquipment(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		tmpl, _ := template.ParseFiles("templates/layout.tmpl", "templates/footer.tmpl", "templates/login.tmpl")
		tmpl.ExecuteTemplate(w, "layout",nil)
	} else {
		if session.Values["admin"].(bool){
			equipmentArray, _:=model.GetAllEquipment()
			name, _ := ((session.Values["name"].(string)))

			user, _:= model.GetClientByName(name)
			var adminEquipmentArray []AdminEquipments
			for _, equipment:= range equipmentArray{
				returnUntil :=""
				if equipment.Status == "nicht verfügbar"{
					borrow, _ :=model.GetReturnUntil(equipment.ID)
					fmt.Println(borrow)
					returnUntil = (borrow.ReturnUntil)
				}

				adminEquipment :=AdminEquipments{Image:equipment.Image, Name:equipment.Name, ID:equipment.ID, Location:equipment.Location, Description:equipment.Description, Quantity:equipment.Quantity, Status:equipment.Status, ReturnUntil:returnUntil}
				adminEquipmentArray= append(adminEquipmentArray, adminEquipment)
			}
			 data := AdminEquipmentStruct{Name:user.Name, Image:user.Image, Equipment:adminEquipmentArray}
			 tmpl, _ := template.ParseFiles("templates/layout.tmpl","templates/footer2.tmpl", "templates/admin/equipment.tmpl")
			 tmpl.ExecuteTemplate(w, "layout", data)
			}else{
				tmpl, _ := template.ParseFiles("templates/layout.tmpl", "templates/footer.tmpl", "templates/login.tmpl")
				tmpl.ExecuteTemplate(w, "layout",nil)
			}
		}
}

func Cart(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	name, _ := ((session.Values["name"].(string)))
	user, _:= model.GetClientByName(name)


	equipments, _:= session.Values["equipments"].([]model.Equipment)
	currentTime:= time.Now().Local()
	date:= currentTime.AddDate(0,0,7).Format("02.01.2006")
	var warenkorb []CartEquipment

	for _, equipment := range equipments{
		tableE , _:= model.GetEquipmentByID(equipment.ID)
		cartEquipment := CartEquipment{Image:equipment.Image, Name:equipment.Name, ID:equipment.ID, Description:equipment.Description, Quantity:equipment.Quantity, Date:date, MaxQuantity:tableE.Quantity}
			fmt.Print("cartEquipment: ")
			fmt.Println(cartEquipment)
		warenkorb = append(warenkorb, cartEquipment)
	}
	data := CartStruct{Name:name, Role:user.Role, Image:user.Image, Warenkorb:warenkorb}
	//fmt.Print("user.Image: ")
	//fmt.Println(user.Image)
	tmpl, _ := template.ParseFiles("templates/layout.tmpl", "templates/footer.tmpl", "templates/cart.tmpl")
	tmpl.ExecuteTemplate(w, "layout", data)
}

func AddToCart(w http.ResponseWriter, r *http.Request) {
		ID, _ := strconv.Atoi(r.FormValue("ID"))
		equipment, _ := model.GetEquipmentByID(ID)
		fmt.Println("AddToCart - equipment: ")
		fmt.Println(equipment)
		equipment.Quantity=1
		fmt.Println("AddToCart - equipment(nachdem equipment.Quantity=1): ")
		fmt.Println(equipment)
		session, _ := store.Get(r, "session")
		if session.Values["equipments"] == nil {
			equipments = nil
		} else {
			equipments = session.Values["equipments"].([]model.Equipment)
		}

		session.Values["equipments"] = append(equipments,equipment)
		err :=session.Save(r, w)
		if err != nil{
		//fmt.Println(err)
	}
		http.Redirect(w, r, "/cart", http.StatusFound)
}

func CartQuantity(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session")
	if err!=nil {panic(err)}

	if r.Method == "POST" {
		equipments:= session.Values["equipments"].([]model.Equipment)
		var newEquipments []model.Equipment
		for _, equipment := range equipments {
			ID:= strconv.Itoa(equipment.ID)
			//fmt.Print("equipment.ID aus aktueller Session: ")
			//fmt.Println(ID)
			//fmt.Println()

			fieldname:=r.FormValue(("quantity") + ID)
			//fmt.Print("fieldname: ")
			//fmt.Println(fieldname)
			//fmt.Println()
			quantity, _ :=strconv.Atoi(fieldname)
			equipment.Quantity = quantity
			//fmt.Print("neue equipment.Quantity in aktueller Session: ")
			//fmt.Println(equipment.Quantity)
			//fmt.Println()

			fmt.Print("CartQuantity - Equipment in Session mit neuer Quantity: ")
			fmt.Println(equipment)
			fmt.Print("CartQuantity - Equipment in Session Equipment.Quantity: ")
			//fmt.Println(equipment.quantity)
			fmt.Println()
			session.Values["equipments"] = append(newEquipments,equipment)
			newEquipments = append(newEquipments,equipment)
			}
//fmt.Print("equipments (kompletter Cart) in session: ")
//fmt.Println(session.Values["equipments"])
//fmt.Println()
		}

		error :=session.Save(r,w)
		if err != nil{panic(error)}
		http.Redirect(w, r, "/cart", http.StatusFound)
}

func DeleteFromCart(w http.ResponseWriter, r *http.Request) {
	//get a session
	session, _ := store.Get(r, "session")
	equipments:= session.Values["equipments"].([]model.Equipment)

	//fmt.Println("Delete From Cart: (Inhalt des Arrays)")
	//fmt.Println(equipments)

	var newEquipments []model.Equipment

	//fmt.Println("Delete From Cart: (Inhalt des Arrays)")
	//fmt.Println(equipments)
	for _, equipment := range equipments {
		keys,_:=r.URL.Query()["ID"]
		ID, _ :=strconv.Atoi(keys[0])
			if equipment.ID != ID {
				session.Values["equipments"] = append(newEquipments,equipment)
				newEquipments = append(newEquipments,equipment)
			}
	}

	session.Values["equipments"] = newEquipments
	err :=session.Save(r, w)
	if err != nil {panic(err)}
	http.Redirect(w, r, "/cart", http.StatusFound)
}

func Equipmenta(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session")

		searchA,searchBool:=(r.URL.Query()["Search"])
		categoryA,categoryBool:=r.URL.Query()["Category"]
		sortByA,sortByBool:=r.URL.Query()["SortBy"]
		var search string
		var category string
		var sortBy string

		if searchBool {
			search ="%"+searchA[0]+"%"
			fmt.Println(search)
		}
		if categoryBool {
			category =categoryA[0]
			fmt.Println(category)
		}
		if sortByBool {
			sortBy =sortByA[0]
			fmt.Println(sortBy)
		}

		var EquipmentArray []model.Equipment
		//Fallunterscheidung für Sortierung

		if searchBool && categoryBool && sortByBool{
			EquipmentArray, _=model.GetEquipmentSearchCategoryOrderBy(search, category, sortBy)
		}else if searchBool && !categoryBool && !sortByBool{
			EquipmentArray, _=model.GetEquipmentSearch(search)
		}else if searchBool && categoryBool && !sortByBool{
			EquipmentArray, _=model.GetEquipmentSearchCategory(search, category)
		}else if searchBool && !categoryBool && sortByBool{
				EquipmentArray, _=model.GetEquipmentSearchOrderBy(search, sortBy)
		}else if !searchBool && !categoryBool && sortByBool{
					EquipmentArray, _=model.GetEquipmentOrderBy(sortBy)
		}else if !searchBool && categoryBool && sortByBool{
					EquipmentArray, _=model.GetEquipmentCategoryOrderBy(category,sortBy)
		}else if !searchBool && categoryBool && !sortByBool{
					EquipmentArray, _=model.GetEquipmentCategory(category)
		}else {
			EquipmentArray, _=model.GetAllEquipment()
		}

		var EquipmentListe []EquipmentPlus
		var name string
		var role string
		var image string
		borrowed:= 0
		reserved:= 0

		if auth, ok := session.Values["authenticated"].(bool); ok || auth {
			name = (session.Values["name"].(string))
			user, _:=model.GetClientByName(name)
			role = user.Role
			image = user.Image
		}

		for _, equipment := range EquipmentArray{
			if auth, ok := session.Values["authenticated"].(bool); ok || auth {
				borrowed=0
				reserved=0
				name = (session.Values["name"].(string))
				user, _:=model.GetClientByName(name)
				_, err:=model.GetBorrow(equipment.ID, user.ID)
				fmt.Print("Controller.Equipmenta - equipment: ")
				fmt.Println(equipment.ID)
				if err == nil{
					borrowed=1

				}
				_, er:=model.GetReserve(equipment.ID, user.ID)
				if er == nil{
					reserved=1

				}
			}

			equipmentPlus:= EquipmentPlus{ID:equipment.ID, Name:equipment.Name, Description:equipment.Description, Image:equipment.Image, Category:equipment.Category, Content:equipment.Content, Status:equipment.Status, Quantity:equipment.Quantity, Borrowed:borrowed, Reserved:reserved}
			EquipmentListe =append(EquipmentListe , equipmentPlus)
		}
		equipmentaStruct := EquipmentaStruct{Name:name, Role:role, Image:image, EquipmentListe:EquipmentListe}

		tmpl, _ := template.ParseFiles("templates/layout.tmpl", "templates/footer2.tmpl","templates/equipment.tmpl")
		tmpl.ExecuteTemplate(w, "layout", equipmentaStruct)
		}

func MyEquipment(w http.ResponseWriter, r *http.Request) {
	//get a session
	session, _ := store.Get(r, "session")

		// Check if user is authenticated
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			tmpl, _ := template.ParseFiles("templates/layout.tmpl", "templates/footer.tmpl","templates/login.tmpl")
			tmpl.ExecuteTemplate(w, "layout", nil)
		} else {
	 name := (session.Values["name"].(string))
	 role := (session.Values["admin"].(bool))
	 user, _ := (model.GetClientByName(name))
	entliehen, _ :=model.GetEntliehen(user.ID)
	vorgemerkt, _ :=model.GetVorgemerkt(user.ID)
	data:= MyEquipmentStruct{Name:name, Role:role,Image:user.Image, Entliehen:entliehen, Vorgemerkt:vorgemerkt}

	tmpl, _ := template.ParseFiles("templates/layout.tmpl", "templates/footer.tmpl","templates/my-equipment.tmpl")
	tmpl.ExecuteTemplate(w, "layout",data)
	}
}

func Borrow(w http.ResponseWriter, r *http.Request) {
	//get a session
	session, _ := store.Get(r, "session")
	fmt.Println("Im controller.Borrow: ")
	//Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		tmpl, _ := template.ParseFiles("templates/layout.tmpl","templates/footer.tmpl", "templates/login.tmpl")
		tmpl.ExecuteTemplate(w, "layout", nil)
	} else {

		user,_:=model.GetClientByName((session.Values["name"].(string)))
		//fmt.Print("User aus Session: ")
		//fmt.Println(user)
		clientID:= user.ID

		currentTime:= time.Now().Local()
		borrowedOn:=currentTime.Format("02.01.2006")
		returnUntil:= currentTime.AddDate(0,0,7).Format("02.01.2006")
		equipments:= session.Values["equipments"].([]model.Equipment)

		for _, equipment := range equipments {
			equipmentID:= equipment.ID
			quantity:=equipment.Quantity
			fmt.Print("(aus Session) Anzahl der ausgeliehenen Geräte dieses Typs: ")
			fmt.Println(quantity)
			borrow := model.Borrow{EquipmentID:equipmentID, ClientID:clientID, BorrowedOn: borrowedOn, ReturnUntil:returnUntil, Quantity:quantity}
			borrow.Borrow()

			session.Values["equipments"]=nil
			error :=session.Save(r,w)
			if error != nil{panic(error)}

			tableEquipment, _:=model.GetEquipmentByID(equipmentID)
			newQuantity:=tableEquipment.Quantity-quantity
			fmt.Print("Equipment in der Datenbank: ")
			fmt.Println(tableEquipment)
			fmt.Println("bisherige Anzahl in Datenbank: ")
			fmt.Println(tableEquipment.Quantity)
			fmt.Print("neue Anzahl, die in Datenbank geschrieben werden soll: ")
			fmt.Println(newQuantity)

			tableEquipment.SetQuantity(newQuantity)

			if newQuantity <= 0 {
				tableEquipment.SetStatusToNA()
			}
		}

		http.Redirect(w, r, "/cart", http.StatusFound)
	}
}

func Reserve(w http.ResponseWriter, r *http.Request) {
	//get a session
	session, _ := store.Get(r, "session")
	//fmt.Println("controller.Borrow")
	//Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		tmpl, _ := template.ParseFiles("templates/layout.tmpl","templates/footer.tmpl", "templates/login.tmpl")
		tmpl.ExecuteTemplate(w, "layout", nil)
	} else {

		keys,_:=r.URL.Query()["ID"]
		ID, _ :=strconv.Atoi(keys[0])
		//fmt.Println("equipment.ID: ")
		//fmt.Println(ID)

		user,_:=model.GetClientByName((session.Values["name"].(string)))
		clientID:= user.ID

		currentTime:= time.Now().Local()
		reservedOn:=currentTime.Format("02.01.2006")
		borrow, _ :=model.GetReturnUntil(ID)
		returnUntil:= borrow.ReturnUntil
		//fmt.Println("returnUntil: ")
		//fmt.Println(returnUntil)

		reserve := model.Reserve{EquipmentID:ID, ClientID:clientID, ReturnUntil:returnUntil, ReservedOn:reservedOn}
		reserve.Reserve()

		http.Redirect(w, r, "/equipment", http.StatusFound)
	}
}

func DeleteReserve(w http.ResponseWriter, r *http.Request) {
	//get a session
	session, _ := store.Get(r, "session")

	keys,_:=r.URL.Query()["ID"]
	equipmentID, _ :=strconv.Atoi(keys[0])
	user,_:=model.GetClientByName((session.Values["name"].(string)))
	model.DeleteReserve(user.ID, equipmentID)

	http.Redirect(w, r, "/my-equipment", http.StatusFound)
}


func Extend(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	user,_:=model.GetClientByName((session.Values["name"].(string)))
	clientID:= user.ID

	keys,_:=r.URL.Query()["ID"]
	equipmentID, _ :=strconv.Atoi(keys[0])
	borrow, _:=model.GetBorrow(equipmentID, clientID)

	date, err:= time.Parse("02.01.2006", borrow.ReturnUntil)
	if err != nil{
		panic(err)
	}
  extendedDate:=(date.AddDate(0,0,7)).Format("02.01.2006")
	borrow.SetDate(extendedDate)

	http.Redirect(w, r, "/my-equipment", http.StatusFound)
}
