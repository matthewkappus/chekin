package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/matthewkappus/src/chekin/reports"
)

var CheckoutList reports.CheckoutList

func init() {
	// Open the file
	studentInventoryStatus, err := os.Open("data/StudentInventoryStatus.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	CheckoutList, err = reports.CreateCheckoutList(*csv.NewReader(studentInventoryStatus))
	if err != nil {
		log.Fatal(err)
	}

}
func greet(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, CheckoutList)
}

func main() {

	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
}
