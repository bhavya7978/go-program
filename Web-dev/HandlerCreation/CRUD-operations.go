package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {

	http.HandleFunc("/getcustomerdetails", getCustomerDetails)
	http.HandleFunc("/getsalesmandetails", getsalesmanDetails)
	http.HandleFunc("/createcustomerdetails", createCustomerDetails)
	http.HandleFunc("/updatecustomerdetails", updateCustomerDetails)
	http.HandleFunc("/deletecustomerdetails", deleteCustomerDetails)
	http.ListenAndServe(":8080", nil)
}

type customerDetails struct {
	CustomerId       int64  `db:"customer_id" json:"customer_id"`
	CustomerName     string `db:"customer_name" json:"customer_name"`
	CustomerMobile   int64  `db:"customer_mobile" json:"customer_mobile"`
	CustomerTicketId int64  `db:"customer_ticketid" json:"customer_ticketid"`
}

type salesmanDetails struct {
	SalesmanID   int64  `db:"salesman_id" json:"salesman_id"`
	SalesmanName string `db:"salesman_name" json:"salesman_name"`
	SalesmanCity string `db:"salesman_city" json:"salesman_city"`
	CustomerId   int64  `db:"customer_id" json:"customer_id"`
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

//ParsingByURL
func getCustomerDetails(w http.ResponseWriter, r *http.Request) {
	db := OpenConnection()

	param := r.URL.Query().Get("customer_id")
	param1 := r.URL.Query().Get("customer_name")

	var custDet customerDetails
	str := `SELECT * FROM customer WHERE customer_id = $1 AND customer_name= $2`

	err := db.Get(&custDet, str, param, param1)
	if err != nil {
		fmt.Println(err)
		fmt.Errorf("BAD REQUEST", err)

	} else {
		byte, _ := json.Marshal(custDet)
		w.Write(byte)
	}

}

//ParsingByBody
func getsalesmanDetails(w http.ResponseWriter, r *http.Request) {
	db := OpenConnection()
	//var salesDet = salesmanDetails{}
	var salesDet1 = []salesmanDetails{}

	/*err := json.NewDecoder(r.Body).Decode(&salesDet1)

	if err != nil {
		panic(err)
	}*/

	str := `SELECT * FROM salesman`

	//err = db.Get(&salesDet, str, salesDet.SalesmanID)
	err := db.Select(&salesDet1, str)
	if err != nil {
		fmt.Println(err)
	} else {
		byte, _ := json.Marshal(&salesDet1)
		w.Write(byte)
	}
}

//CREATE=POST Request
func createCustomerDetails(w http.ResponseWriter, r *http.Request) {
	db := OpenConnection()
	var custDet customerDetails
	err := json.NewDecoder(r.Body).Decode(&custDet)
	if err != nil {
		fmt.Println(err)
		fmt.Errorf("BAD REQUEST:", err)
		os.Exit(http.StatusBadRequest)
	}

	str := `INSERT INTO customer(customer_name,customer_mobile,customer_ticketid)
	          values($1,$2,$3)`

	_, err = db.Exec(str, custDet.CustomerName, custDet.CustomerMobile, custDet.CustomerTicketId)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Fprintln(w, "CREATED SUCCESSFULLY!!")
	}
}

//UPDATE=PUT Request
func updateCustomerDetails(w http.ResponseWriter, r *http.Request) {
	db := OpenConnection()
	var custDet customerDetails
	err := json.NewDecoder(r.Body).Decode(&custDet)
	if err != nil {
		fmt.Println(err)
		fmt.Errorf("BAD REQUEST:", err)
	}

	str := `UPDATE customer SET  customer_name = $1 WHERE customer_id= $2`
	_, err = db.Exec(str, custDet.CustomerName, custDet.CustomerId)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Fprintln(w, "UPDATED SUCCESSFULLY!!")
	}
}

//DELETE
func deleteCustomerDetails(w http.ResponseWriter, r *http.Request) {
	db := OpenConnection()
	var custDet customerDetails
	err := json.NewDecoder(r.Body).Decode(&custDet)
	if err != nil {
		fmt.Println(err)
		fmt.Errorf("BAD REQUEST:", err)
	}

	str := `DELETE from  customer WHERE customer_name=$1`
	_, err = db.Exec(str, custDet.CustomerName)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Fprintln(w, "DELETED SUCCESSFULLY!!")
	}
}
