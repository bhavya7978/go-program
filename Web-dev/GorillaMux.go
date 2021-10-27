package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/getcustomerdetails", getCustomerDetails).Methods("GET")
	http.ListenAndServe(":8080", router)
}

type customerDetails struct {
	CustomerId       int64  `db:"customer_id" json:"customer_id"`
	CustomerName     string `db:"customer_name" json:"customer_name"`
	CustomerMobile   int64  `db:"customer_mobile" json:"customer_mobile"`
	CustomerTicketId int64  `db:"customer_ticketid" json:"customer_ticketid"`
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "movie"
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

func getCustomerDetails(w http.ResponseWriter, r *http.Request) {
	db := OpenConnection()
	params := r.URL.Query().Get("customer_id")

	var customerDet customerDetails
	str := `SELECT * FROM customer WHERE customer_id = $1`

	err := db.Get(&customerDet, str, params)
	if err != nil {
		panic(err)
	} else {
		byte, _ := json.Marshal(customerDet)
		w.Write(byte)
	}
}
