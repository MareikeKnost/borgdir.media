package model

import (
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
)

type Equipment struct {
	ID int
	Name string
	Description string
	Image string
	Category string
	Content string
	Location string
	Quantity int
	Status string
	MaxQuantity int
}

type User struct {
	ID int
	Name string
	Email string
	Password string
	Status string
	ActiveUntil string
	Image string
	Role string
}

type Borrow struct {
	EquipmentID int
	ClientID int
	BorrowedOn string
	ReturnUntil string
	Quantity int
}

type Reserve struct {
	EquipmentID int
	ClientID int
	ReturnUntil string
	ReservedOn string
}

type Borrowed struct {
	Name string
	ID int
	Description string
	BorrowedOn string
	ReturnUntil string
	Image string
}

type Reserved struct {
	Name string
	ID int
	Description string
	ReturnUntil string
	Image string
}

// Db handle
var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("sqlite3", "./data/borgdirmedia.db")
	if err != nil {
		fmt.Println(err)
	}
}

// GetAll Equipment
func GetAllEquipment() (equipments []Equipment, err error) {
	rows, err := Db.Query("select ID, Name, Description, Image, Category, Content, Location, Quantity, Status from Equipment")
	if err != nil {
		log.Println("Fehler Nr. 1 in model.GetAllEquipment")
		panic(err)
		return
	}
	for rows.Next() {
		equipment := Equipment{}
		err = rows.Scan(&equipment.ID, &equipment.Name, &equipment.Description, &equipment.Image, &equipment.Category, &equipment.Content, &equipment.Location, &equipment.Quantity, &equipment.Status)
		if err != nil {
			log.Println("Fehler Nr. 2 in model.GetAllEquipment")
			panic(err)
			return
		}
		equipments = append(equipments, equipment)
	}
	rows.Close()
	return
}

// GetAll Equipment
func GetUserEquipment(userID int) (equipments []Equipment, err error) {
	rows, err := Db.Query("select * from Equipment where ID=$1", userID)
	if err != nil {
		log.Println("Fehler Nr. 1 in model.GetAllEquipment")
		return
	}
	for rows.Next() {
		equipment := Equipment{}
		err = rows.Scan(&equipment.ID, &equipment.Name, &equipment.Description, &equipment.Image, &equipment.Category, &equipment.Content, &equipment.Location, &equipment.Quantity, &equipment.Status)
		if err != nil {
			log.Println("Fehler Nr. 2 in model.GetAllEquipment")
			return
		}
		equipments = append(equipments, equipment)
	}
	rows.Close()
	return
}


// GetEquipmentSearchCategoryOrderBy
func GetEquipmentSearchCategoryOrderBy(search string, category string, orderBy string) (equipments []Equipment, err error) {
	rows, err := Db.Query("select ID, name, description, image, category, content, location, quantity, Status from equipment where category=$1 and name like $2 order by $3 ASC", category, search, orderBy)
	if err != nil {
		log.Println("Fehler Nr. 1 in model.GetAllEquipment")
		return
	}
	for rows.Next() {
		equipment := Equipment{}
		err = rows.Scan(&equipment.ID, &equipment.Name, &equipment.Description, &equipment.Image, &equipment.Category, &equipment.Content, &equipment.Location, &equipment.Quantity, &equipment.Status)
		if err != nil {
			log.Println("Fehler Nr. 2 GetEquipmentSearchCategoryOrderBy")
			return
		}
		equipments = append(equipments, equipment)
	}
	rows.Close()
	return
}

// GetEquipmentSearchCategoryOrderBy
func GetEquipmentCategoryOrderBy(category string, orderBy string) (equipments []Equipment, err error) {
	rows, err := Db.Query("select ID, name, description, image, category, content, location, quantity, Status from equipment where category=$1 order by $3 ASC", category, orderBy)
	if err != nil {
		log.Println("Fehler Nr. 1 in model.GetAllEquipment")
		return
	}
	for rows.Next() {
		equipment := Equipment{}
		err = rows.Scan(&equipment.ID, &equipment.Name, &equipment.Description, &equipment.Image, &equipment.Category, &equipment.Content, &equipment.Location, &equipment.Quantity, &equipment.Status)
		if err != nil {
			log.Println("Fehler Nr. 2 in  GetEquipmentCategoryOrderBy")
			return
		}
		equipments = append(equipments, equipment)
	}
	rows.Close()
	return
}

// GetEquipmentSearchCategoryOrderBy
func GetEquipmentCategory(category string) (equipments []Equipment, err error) {
	rows, err := Db.Query("select ID, name, description, image, category, content, location, quantity, Status from equipment where category=$1", category)
	if err != nil {
		log.Println("Fehler Nr. 1 in model.GetAllEquipment")
		return
	}
	for rows.Next() {
		equipment := Equipment{}
		err = rows.Scan(&equipment.ID, &equipment.Name, &equipment.Description, &equipment.Image, &equipment.Category, &equipment.Content, &equipment.Location, &equipment.Quantity, &equipment.Status)
		if err != nil {
			log.Println("Fehler Nr. 2 in GetEquipmentCategory")
			return
		}
		equipments = append(equipments, equipment)
	}
	rows.Close()
	return
}

// GetEquipmentSearch
func GetEquipmentSearch(search string) (equipments []Equipment, err error) {
	rows, err := Db.Query("select ID, name, description, image, category, content, location, quantity, Status from equipment where name like $1 ", search)
	if err != nil {
		log.Println("Fehler Nr. 1 in model.GetAllEquipment")
		return
	}
	for rows.Next() {
		equipment := Equipment{}
		err = rows.Scan(&equipment.ID, &equipment.Name, &equipment.Description, &equipment.Image, &equipment.Category, &equipment.Content, &equipment.Location, &equipment.Quantity, &equipment.Status)
		if err != nil {
			log.Println("Fehler Nr. 2 in GetEquipmentSearch")
			return
		}
		equipments = append(equipments, equipment)
	}
	rows.Close()
	return
}

// GetEquipmentSearchCategory
func GetEquipmentSearchCategory(search string, category string) (equipments []Equipment, err error) {
	rows, err := Db.Query("select ID, name, description, image, category, content, location, quantity, Status from equipment where category=$1 and name like $2", category, search)
	if err != nil {
		log.Println("Fehler Nr. 1 in model.GetAllEquipment")
		return
	}
	for rows.Next() {
		equipment := Equipment{}
		err = rows.Scan(&equipment.ID, &equipment.Name, &equipment.Description, &equipment.Image, &equipment.Category, &equipment.Content, &equipment.Location, &equipment.Quantity, &equipment.Status)
		if err != nil {
			log.Println("Fehler Nr. 2 in model.GetAllEquipment")
			return
		}
		equipments = append(equipments, equipment)
	}
	rows.Close()
	return
}

// GetEquipmentSearchCategoryOrderBy
func GetEquipmentSearchOrderBy(search string, orderBy string) (equipments []Equipment, err error) {
	rows, err := Db.Query("select ID, name, description, image, category, content, location, quantity, Status from equipment where name like $1 order by $2 ASC", search, orderBy)
	if err != nil {
		log.Println("Fehler Nr. 1 in model.GetAllEquipment")
		return
	}
	for rows.Next() {
		equipment := Equipment{}
		err = rows.Scan(&equipment.ID, &equipment.Name, &equipment.Description, &equipment.Image, &equipment.Category, &equipment.Content, &equipment.Location, &equipment.Quantity, &equipment.Status)
		if err != nil {
			log.Println("Fehler Nr. 2 in model.GetAllEquipment")
			return
		}
		equipments = append(equipments, equipment)
	}
	rows.Close()
	return
}

// GetEquipmentSearchCategoryOrderBy
func GetEquipmentOrderBy(orderBy string) (equipments []Equipment, err error) {
qtext := fmt.Sprintf("select ID, name, description, image, category, content, location, quantity, Status from equipment order by %s ASC", orderBy)
	rows, err := Db.Query(qtext)
	if err != nil {
		panic(err)
		return
	}
	for rows.Next() {
		equipment := Equipment{}
		err = rows.Scan(&equipment.ID, &equipment.Name, &equipment.Description, &equipment.Image, &equipment.Category, &equipment.Content, &equipment.Location, &equipment.Quantity, &equipment.Status)
		if err != nil {
			log.Println("Fehler Nr. 2 in model.GetAllEquipment")
			return
		}
		log.Println(equipment.Name)
		equipments = append(equipments, equipment)
	}
	rows.Close()
	return
}


// Add Equipment
func (equipment *Equipment) AddEquipment() (err error) {
	statement := "insert into equipment (name, description, image, category, content, location, quantity, maxQuantity) values ($1, $2, $3, $4, $5, $6, $7, $8)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		log.Println("Fehler in model.AddEquipment()")
		return
	}

	defer stmt.Close()
	_, err = stmt.Exec(equipment.Name, equipment.Description, equipment.Image, equipment.Category, equipment.Content, equipment.Location, equipment.Quantity, equipment.MaxQuantity)
	return
}

// GetEntliehen with the provided id
func GetEntliehen(id int) (borroweds []Borrowed, err error) {
	//Abfragen für Array Entliehen
	rows, err := Db.Query("select equipment.ID, equipment.name, equipment.description, equipment.Image, borrow.returnUntil, borrow.borrowedOn from equipment, borrow where borrow.clientID=$1 and borrow.equipmentID=equipment.ID", id)
	if err != nil {	fmt.Println(err)}

	for rows.Next() {
		equipment := Equipment{}
		borrow := Borrow{}
		err = rows.Scan(&equipment.ID, &equipment.Name, &equipment.Description, &equipment.Image, &borrow.ReturnUntil, &borrow.BorrowedOn)

		if err != nil {
			log.Println("Fehler Nr. 2 in model.GetAllEquipment")
			return
		}

		borrowed := Borrowed{Name:equipment.Name, ID:equipment.ID, Description:equipment.Description, Image:equipment.Image, BorrowedOn:borrow.BorrowedOn, ReturnUntil:borrow.ReturnUntil}
		fmt.Print("borrowed: ")
		fmt.Print(borrowed)
		borroweds = append(borroweds, borrowed)
	}
	rows.Close()
	return
}


// GetVorgemerkt with the provided id
func GetVorgemerkt(id int) (reserveds []Reserved, err error) {
	//Abfragen für Array Entliehen
	rows, err := Db.Query("select equipment.ID, equipment.name, equipment.description, equipment.Image, reserve.returnUntil from equipment, reserve where reserve.clientID=$1 and reserve.equipmentID=equipment.ID", id)
	if err != nil {	fmt.Println(err)}

	for rows.Next() {
		equipment := Equipment{}
		reserve := Reserve{}
		err = rows.Scan(&equipment.ID, &equipment.Name, &equipment.Description, &equipment.Image, &reserve.ReturnUntil)

		if err != nil {
			log.Println("Fehler Nr. 2 in model.GetAllEquipment")
			return
		}

		reserved := Reserved{Name:equipment.Name, ID:equipment.ID, Description:equipment.Description, Image:equipment.Image, ReturnUntil:reserve.ReturnUntil}
		fmt.Print("reserved: ")
		fmt.Print(reserved)
		reserveds = append(reserveds, reserved)
	}
	rows.Close()
	return
}

// Get Equipment with the provided id
func GetEquipmentByID(id int) (equipment Equipment, err error) {
	equipment = Equipment{}
	err = Db.QueryRow("select ID, name, description, image, category, content, location, quantity, status from equipment where id=$1", id).Scan(&equipment.ID, &equipment.Name, &equipment.Description, &equipment.Image, &equipment.Category, &equipment.Content, &equipment.Location, &equipment.Quantity, &equipment.Status)
	fmt.Println("Model.GetEquipmentByID - Equipment")
	fmt.Println(equipment)
	return
}


// Edit Equipment
func (equipment *Equipment) EditEquipment() (err error) {
	_, err = Db.Exec("update equipment set name=$1, description=$2, category=$3, content=$4, location=$5, quantity=$6 where id=$7",equipment.Name, equipment.Description, equipment.Category, equipment.Content, equipment.Location, equipment.Quantity, equipment.ID)

fmt.Println(err)
return
}

// Delete Equipment with the provided id from the list of Equipment
func (equipment *Equipment) DeleteEquipment() (err error) {
	_, err = Db.Exec("delete from equipment where id = $1", equipment.ID)
	return
}

// Add Client
func (user *User) AddClient() (err error) {
	statement := "insert into user (name, email, password) values ($1, $2, $3)"
	stmt, err := Db.Prepare(statement)

	if err != nil {
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Email, user.Password)
	return
}

// GetAll Clients
func GetAllClients() (users []User, err error) {
	rows, err := Db.Query("select * from user")
	if err != nil {
		log.Println("Fehler Nr. 1 in model.GetAllClients")
		return
	}
	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Status, &user.ActiveUntil, &user.Image, &user.Role )
		if err != nil {
			log.Println("Fehler Nr. 2 in model.GetAllClients")
			return
		}
		users = append(users, user)
	}
	rows.Close()
	return
}

//GetBorrowedEquipment
func (user *User)GetBorrowedEquipment() (borrowedEquipment []Equipment, err error) {
	rows, err := Db.Query("select equipment.id, equipment.name from equipment, borrow where equipment.ID=borrow.equipmentID and borrow.clientID=$1", user.ID)
	if err != nil {
		log.Println("Fehler Nr. 1 in model.GetAllClients")
		return
	}
	for rows.Next() {
		equipment := Equipment{}
		//borrow := Borrow{}
		err = rows.Scan(&equipment.ID, &equipment.Name)
		if err != nil {
			panic(err)
			return
		}
		borrowedEquipment = append(borrowedEquipment, equipment)
	}
	rows.Close()
	return
}


// Get Client with the provided id
func GetClientByID(id int) (user User, err error) {
	user = User{}
	err = Db.QueryRow("select * from user where id = $1", id).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Status, &user.ActiveUntil, &user.Image, &user.Role)
fmt.Println(user)

	return
}

// GetUserByUsername retrieve User by username
func GetClientByName(name string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("select * from user where name = $1", name).Scan(&user.ID, &user.Name, &user.Email ,&user.Password, &user.Status,  &user.ActiveUntil,  &user.Image, &user.Role)
	return
}

// Edit Client Mail, Name
func (user *User) EditClient() (err error) {
	_, err = Db.Exec("update user set name=$1, email=$2, image=$3 where id=$4",user.Name, user.Email, user.Image, user.ID)

fmt.Println(err)
return
}

// Edit Client mit Passwort
func (user *User) EditClientPw() (err error) {
	_, err = Db.Exec("update user set name=$1, email=$2 , password=$3 , image=$4 where id=$5",user.Name, user.Email, user.Password, user.Image, user.ID)

fmt.Println(err)
return
}

// Lock Client
func (user *User) LockClient() (err error) {
	_, err = Db.Exec("update user set status = 'gesperrt' where id=$1", user.ID)
	return
}

//Unlock Client
func (user *User) UnlockClient() (err error) {
	_, err = Db.Exec("update user set status = 'aktiv' where id=$1", user.ID)
	return
}

func (borrow *Borrow) Borrow() (err error) {
	statement := "insert into borrow (equipmentID, clientID, borrowedOn, returnUntil, quantity) values ($1, $2, $3, $4, $5)"
	stmt, err := Db.Prepare(statement)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer stmt.Close()
	_, err = stmt.Exec(borrow.EquipmentID, borrow.ClientID, borrow.BorrowedOn, borrow.ReturnUntil, borrow.Quantity)
	return
}

// SetQuantity
func (equipment *Equipment) SetQuantity(newQuantity int) (err error) {
	_, err = Db.Exec("update equipment set quantity=$1 where id=$2", newQuantity, equipment.ID)

fmt.Println(err)
return
}

// SetDate
func (borrow *Borrow) SetDate(date string) (err error) {
	_, err = Db.Exec("update borrow set returnUntil=$1 where equipmentID=$2 and clientID=$3", date, borrow.EquipmentID, borrow.ClientID)

fmt.Println(err)
return
}

// SetStatusToNA
func (equipment *Equipment) SetStatusToNA() (err error) {
	_, err = Db.Exec("update equipment set status='nicht verfügbar' where id=$1", equipment.ID)
return
}

//GetBorrow
func GetBorrow(equipmentID int, clientID int) (borrow Borrow, err error) {
	borrow = Borrow{}
err= Db.QueryRow("select * from borrow where equipmentID=$1 and clientID=$2", equipmentID, clientID).Scan(&borrow.EquipmentID, &borrow.ClientID, &borrow.BorrowedOn, &borrow.ReturnUntil, &borrow.Quantity)
if err!=nil{fmt.Println(err)}
	return
}




//GetReserve
func GetReserve(equipmentID int, clientID int) (reserve Reserve, err error) {
	reserve = Reserve{}
err= Db.QueryRow("select * from reserve where equipmentID=$1 and clientID=$2", equipmentID, clientID).Scan(&reserve.EquipmentID, &reserve.ClientID, &reserve.ReservedOn, &reserve.ReturnUntil)
if err!=nil{fmt.Println(err)}
	return
}

//GetReturnUntil
func GetReturnUntil(ID int) (borrow Borrow, err error) {
	borrow = Borrow{}
	err= Db.QueryRow("select * from borrow where equipmentID=$1 order by strftime('%d.%m.%Y', returnUntil) DESC limit 1", ID).Scan(&borrow.EquipmentID, &borrow.ClientID, &borrow.BorrowedOn, &borrow.ReturnUntil, &borrow.Quantity)
	if err!=nil{fmt.Println(err)}
	return
}

func (reserve *Reserve) Reserve() (err error) {
	statement := "insert into reserve (equipmentID, clientID, returnUntil, reservedOn) values ($1, $2, $3, $4)"
	stmt, err := Db.Prepare(statement)

	if err != nil {
		fmt.Println(err)
		return
	}

		defer stmt.Close()
		_, err = stmt.Exec(reserve.EquipmentID, reserve.ClientID, reserve.ReturnUntil,reserve.ReservedOn)
		return
	}

func	DeleteReserve(userID int, equipmentID int) (err error)	{
_, err = Db.Exec("delete from reserve where equipmentID=$1 and clientID=$2", equipmentID, userID)
	return
}

func	DeleteClient(userID int) (err error)	{
_, err = Db.Exec("delete from user where ID=$1", userID)
	return
}
