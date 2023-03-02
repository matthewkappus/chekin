package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/matthewkappus/chekin/src/reports"
)

var CheckoutList reports.CheckoutList
var Roster reports.Roster

func init() {
	// Open the file
	studentInventoryStatus, err := os.Open("data/StudentInventoryStatus.csv")
	if err != nil {
		log.Fatalln("Couldn't open data/StudentInventoryStatus.csv", err)
	}

	CheckoutList, err = reports.CreateCheckoutList(csv.NewReader(studentInventoryStatus))
	if err != nil {
		log.Println(err)
	}

	studentRoster, err := os.Open("data/StudentRoster.csv")
	if err != nil {
		log.Fatalln("Couldn't open data/StudentRoster.csv", err)
	}
	Roster, err = reports.CreateRosterList(csv.NewReader(studentRoster))
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Data loaded. now listening on port localhost:8080")

}

func main() {

	http.HandleFunc("/students", Roster.ListHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world!")
	})
	http.ListenAndServe(":8080", nil)
}
