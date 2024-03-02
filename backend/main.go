// backend for the gocrud
package main

import (
	"database/sql"
	//"encoding/json"
	"log"
	//"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// database for the user which only contains the Id name and the email

// this is the format to define the type for the User it is like the interface in `json:""` and this json format is used for the encoding and decodiing of the json data
type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// defer follow the LIFO pattern and the state ment which is delared by " defer" gives the out put after the execution of the parent function
func main() {
	// target 1 connect to the database using the basic pattern
	db, err := sql.Open("postgres", os.Getenv("database url"))
	// In the go i cannot write the if err {statement} like the typescript
	if err != nil {
		log.Fatalf("Failed to connect to db: %v", err)
	}
	// after all this close the db connections : it can be achieved by using the defer as it uses the lifo principle
	defer db.Close()

	// auto generate the table ( we do not want to manually create the table)
	// it is like the generation of the db from the schema
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT, email TEXT)")
	if err != nil {
		log.Fatal(err)
	}
	// configure the end points
	// after initailizing the router just hit then endpoint by using the router.handleFunc()

	router := mux.NewRouter()

	// get all the users

	// create the users
	// get a single user by the id
	// update the user by the id
	// delete the user by the id
}
