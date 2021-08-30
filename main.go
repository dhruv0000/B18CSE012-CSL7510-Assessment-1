package main

import (
	"encoding/json"
	"fmt"
	"log"

	"net/http"

	"github.com/gorilla/mux"

	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/rs/cors"
)

var db *gorm.DB

var err error

//----- Models -----

type User struct {
	gorm.Model

	Name string

	RollNo string

	Attendance string
}

var users = []User{
	{Name: "John", RollNo: "1", Attendance: "Present"},
	{Name: "Paul", RollNo: "2", Attendance: "Absent"},
	{Name: "George", RollNo: "3", Attendance: "Present"},
	{Name: "Ringo", RollNo: "4", Attendance: "Absent"},
}

//----- Routes -----

func main() {

	router := mux.NewRouter()

	// Postgres
	dsn := "host=localhost user=postgres password=Pass2020! dbname=postgres port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// SQLite
	// db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {

		fmt.Print(err.Error())
		panic("failed to connect database")

	}

	err = db.AutoMigrate(&User{})
	if err != nil {

		fmt.Print(err.Error())
		panic("failed to Auto Migrate")

	}

	// for index := range users {

	// 	db.Create(&users[index])

	// }

	staticFileDirectory := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	router.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

	router.HandleFunc("/users", GetUsers).Methods("GET")

	router.HandleFunc("/post-users", PostUsers).Methods("POST")

	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":8080", handler))

}

//----- Routes Handlers -----

func GetUsers(w http.ResponseWriter, r *http.Request) {

	var users []User

	db.Find(&users)

	// log.Default().Print(users)

	_ = json.NewEncoder(w).Encode(&users)

}

func PostUsers(w http.ResponseWriter, r *http.Request) {

	var user User

	user.Name = r.FormValue("Name")
	user.RollNo = r.FormValue("RollNo")
	user.Attendance = r.FormValue("Attendance")

	db.Create(&user)

	// _ = json.NewEncoder(w).Encode(&user)
}
