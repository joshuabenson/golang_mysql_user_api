package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func getUsers() []*User {
	// Open up our database connection.
	db, err := sql.Open("mysql", "tester:secret@tcp(db:3306)/josh_db")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	// Execute the query
	results, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	var users []*User
	for results.Next() {
		var u User
		// for each row, scan the result into our tag composite object
		err = results.Scan(&u.ID, &u.Name)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		users = append(users, &u)
	}

	return users
}

func insertUser(w http.ResponseWriter, r *http.Request) (int64, error) {

	// Open up our database connection.
	db, err := sql.Open("mysql", "tester:secret@tcp(db:3306)/josh_db")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}

	var requestBody User
	defer db.Close()
	jsonErr := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, jsonErr.Error(), http.StatusBadRequest)
	}
	if requestBody.Name == "" {
		http.Error(w, "request body should contain `name` property", http.StatusBadRequest)
		return -1, errors.New("error")
	}
	// Insert the name passed in the body of the request
	results, queryErr := db.Exec("INSERT INTO users (`name`) VALUES (?);", requestBody.Name)
	if queryErr != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	return results.LastInsertId()
}

//func checkExistingUser(w http.ResponseWriter, r *http.Request) (int64, error) {
// Open up our database connection.
//db, err := sql.Open("mysql", "tester:secret@tcp(db:3306)/josh_db")
//
//// if there is an error opening the connection, handle it
//if err != nil {
//	log.Print(err.Error())
//}
//
//var requestBody User
//defer db.Close()
//jsonErr := json.NewDecoder(r.Body).Decode(&requestBody)
//if err != nil {
//	http.Error(w, jsonErr.Error(), http.StatusBadRequest)
//}
//
//// Insert the name passed in the body of the request
//println(requestBody.Name)
//results, queryErr := db.Query("SELECT * FROM users WHERE `name` = \"test\"")
//
//
////if queryErr != nil {
////	panic(err.Error()) // proper error handling instead of panic in your app
////}
//
//return results.LastInsertId()
//}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage Bitch!")
	fmt.Println("Endpoint Hit: homePage")
}

func userPage(w http.ResponseWriter, r *http.Request) {
	users := getUsers()
	fmt.Println("Endpoint Hit: /users")
	json.NewEncoder(w).Encode(users)
}

func addUser(w http.ResponseWriter, r *http.Request) {
	users, error := insertUser(w, r)
	fmt.Println("Endpoint Hit: /addUser")
	if error == nil {
		json.NewEncoder(w).Encode(users)
	}
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/users", userPage)
	http.HandleFunc("/addUser", addUser)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
