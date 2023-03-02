package main

import (
	"encoding/csv"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/matthewkappus/chekin/src/reports"
)

var CheckoutList reports.CheckoutList
var Roster reports.Roster

func init() {
	templates, err := template.ParseGlob("templates/*.tmpl.html")
	if err != nil {
		log.Fatalln("Couldn't parse templates", err)
	}
	// Open the file
	studentInventoryStatus, err := os.Open("data/StudentInventoryStatus.csv")
	if err != nil {
		log.Fatalln("Couldn't open data/StudentInventoryStatus.csv", err)
	}

	CheckoutList, err = reports.CreateCheckoutList(csv.NewReader(studentInventoryStatus), templates)
	if err != nil {
		log.Println(err)
	}

	studentRoster, err := os.Open("data/StudentRoster.csv")
	if err != nil {
		log.Fatalln("Couldn't open data/StudentRoster.csv", err)
	}
	Roster, err = reports.CreateRosterList(csv.NewReader(studentRoster), templates)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Data loaded. now listening on port localhost:8080")

}

func main() {
	http.HandleFunc("/lookup", CheckoutList.LookupHandler)
	http.HandleFunc("/lookup_inventory", CheckoutList.ShowLookupHandler)
	http.HandleFunc("/students", Roster.ListHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world!")
	})
	http.ListenAndServe(":8080", nil)
}
