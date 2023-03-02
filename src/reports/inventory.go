package reports

import (
	"encoding/csv"
	"fmt"
	"html/template"
	"io"
	"net/http"
)

type CheckoutList struct {
	tmpl      *template.Template
	Inventory map[string][]Checkout
}

func CreateCheckoutList(studentaccessorycsv *csv.Reader, templates *template.Template) (CheckoutList, error) {
	cl := CheckoutList{tmpl: templates, Inventory: make(map[string][]Checkout)}

	var readErrors = make([]error, 0)
	for {
		record, err := studentaccessorycsv.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			// TODO: Append read errors to a string and return it
			readErrors = append(readErrors, err)
			continue
		}
		c := Checkout{
			StudentID:       record[0],
			LastName:        record[1],
			FirstName:       record[2],
			MiddleName:      record[3],
			Grade:           record[4],
			Homeroom:        record[5],
			Tag:             record[6],
			ProductName:     record[7],
			AccessoryName:   record[8],
			IssuedDate:      record[9],
			QuantityIssued:  record[10],
			QuantityMissing: record[11],
			MissingValue:    record[12],
		}

		// append checkouts to students inventory or create a new inventory
		if cl.Inventory[c.StudentID] == nil {
			cl.Inventory[c.StudentID] = make([]Checkout, 0)
		} else {
			cl.Inventory[c.StudentID] = append(cl.Inventory[c.StudentID], c)
		}

	}
	var err error
	if len(readErrors) > 0 {
		err = fmt.Errorf("%d errors occurred while reading the inventory file", len(readErrors))
	}

	return cl, err
}

type Checkout struct {
	StudentID       string
	LastName        string
	FirstName       string
	MiddleName      string
	Grade           string
	Homeroom        string
	Tag             string
	ProductName     string
	AccessoryName   string
	IssuedDate      string
	QuantityIssued  string
	QuantityMissing string
	MissingValue    string
}

func (cl CheckoutList) LookupHandler(w http.ResponseWriter, r *http.Request) {
	cl.tmpl.Lookup("lookup_form").Execute(w, cl)
}

func (cl CheckoutList) ShowLookupHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "looking up student %s", r.FormValue("studentid"))
}
