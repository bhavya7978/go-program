package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	http.HandleFunc("/userdetails", getUserDetails)
	http.ListenAndServe(":8080", nil)

}

type userDetails struct {
	User_id      int64  `db:"user_id" json:"user_id"`
	First_Name   string `db:"first_name" json:"first_name"`
	Last_Name    string `db:"last_name" json:"last_name"`
	Age          int64  `db:"age" json:"age"`
	SearchString string `json:"searchString"`
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "user"
)

func OpenConnection() *sqlx.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	//defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println(psqlInfo)
	fmt.Println("Successfully connected!")
	return db
}

func getUserDetails(w http.ResponseWriter, r *http.Request) {
	db := OpenConnection()
	//Declaration of variable for storing the data from database
	var userDet = []userDetails{}

	//Declaration of variable for storing decoded value of searchstring
	var SearchstringfromBody userDetails

	//Decoding  and storing the value seachstring
	err := json.NewDecoder(r.Body).Decode(&SearchstringfromBody)

	//Decoding & storing the values in userDet
	err = json.NewDecoder(r.Body).Decode(&userDet)
	if err != nil {
		fmt.Println(err)
	}

	Str := `SELECT * FROM userdetails WHERE first_name LIKE '%%%s%%'`
	Str = fmt.Sprintf(Str, SearchstringfromBody.SearchString)
	fmt.Println(Str)

	err = db.Select(&userDet, Str)
	//fmt.Println(userDet)
	for i := range userDet {
		userDet[i].First_Name = strings.ReplaceAll(userDet[i].First_Name, " ", "")
		userDet[i].Last_Name = strings.ReplaceAll(userDet[i].Last_Name, " ", "")

	}
	fmt.Println(userDet)

	if err != nil {
		fmt.Println(err)
	} else {
		byte, _ := json.Marshal(&userDet)
		w.Write(byte)

	}
}
