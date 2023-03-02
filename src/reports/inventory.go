package reports

import (
	"encoding/csv"
)

type CheckoutList struct {
	CheckoutList []Checkout
}

func CreateCheckoutList(studentaccessorycsv csv.Reader) (CheckoutList, error) {
	return CheckoutList{}, nil
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
