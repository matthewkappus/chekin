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

		// remove leading =" from StudentID and trailing "
		sid := record[1][2 : len(record[1])-1]
		c := Checkout{
			SiteName: record[0],
			// remove leading =" from StudentID and trailing "
			StudentID:      sid,
			LastName:       record[2],
			FirstName:      record[3],
			MiddleName:     record[4],
			Grade:          record[5],
			HomeRoom:       record[6],
			StudentNotes:   record[7],
			ProductName:    record[8],
			Model:          record[9],
			ProductType:    record[10],
			SuggestedPrice: record[11],
			Tag:            record[12],
			Serial:         record[13],
			AssetType:      record[14],
			ParentTag:      record[15],
			Status:         record[16],
			// todo: convert to date
			ScanDate:    record[17],
			StatusNotes: record[18],
		}

		// append checkouts to students inventory or create a new inventory
		if cl.Inventory[c.StudentID] == nil {
			cl.Inventory[c.StudentID] = make([]Checkout, 0)
			cl.Inventory[c.StudentID] = append(cl.Inventory[c.StudentID], c)
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
	SiteName       string
	StudentID      string
	LastName       string
	FirstName      string
	MiddleName     string
	Grade          string
	HomeRoom       string
	StudentNotes   string
	ProductName    string
	Model          string
	ProductType    string
	SuggestedPrice string
	Tag            string
	Serial         string
	AssetType      string
	ParentTag      string
	Status         string
	ScanDate       string
	StatusNotes    string
}

func (cl CheckoutList) LookupHandler(w http.ResponseWriter, r *http.Request) {
	cl.tmpl.Lookup("lookup_form").Execute(w, cl)
}

func (cl CheckoutList) ShowLookupHandler(w http.ResponseWriter, r *http.Request) {
	// todo: lookup student id by login email

	// Austin Zimmerman: 980042643
	sid := r.FormValue("studentid")
	fmt.Fprintf(w, "looking up student %s", sid)
	if cl.Inventory[sid] == nil {
		fmt.Fprintf(w, "No inventory for %s", sid)
	} else {
		fmt.Fprintf(w, "Inventory for %s", sid)
		fmt.Fprintf(w, "%v", cl.Inventory[sid])
	}
}
