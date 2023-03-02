package reports

import (
	"encoding/csv"
	"fmt"
	"html/template"
	"io"
	"net/http"
)

type Stu415 struct {
	OrganizationName string
	SchoolYear       string
	StudentName      string
	PermID           string
	Gender           string
	Grade            string
	TermName         string
	Per              string
	Term             string
	SectionID        string
	CourseIDAndTitle string
	MeetDays         string
	Teacher          string
	Room             string
	PreScheduled     string
}

type Roster struct {
	tmpl     *template.Template
	Students []Stu415
}

func CreateRosterList(stu415s *csv.Reader, templates *template.Template) (Roster, error) {
	list := Roster{tmpl: templates}

	var readErrors = make([]error, 0)
	for {
		record, err := stu415s.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			readErrors = append(readErrors, err)
			continue
		}
		list.Students = append(list.Students, Stu415{
			OrganizationName: record[0],
			SchoolYear:       record[1],
			StudentName:      record[2],
			PermID:           record[3],
			Gender:           record[4],
			Grade:            record[5],
			TermName:         record[6],
			Per:              record[7],
			Term:             record[8],
			SectionID:        record[9],
			CourseIDAndTitle: record[10],
			MeetDays:         record[11],
			Teacher:          record[12],
			Room:             record[13],
			PreScheduled:     record[14],
		})
	}
	var err error
	if len(readErrors) > 0 {
		err = fmt.Errorf("%d errors occurred while reading roster file", len(readErrors))
	}
	return list, err
}

func (rl Roster) ListHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, rl.Students)
}
