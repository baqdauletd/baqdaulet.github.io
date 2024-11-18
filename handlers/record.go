package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type Record struct {
    Email        string `gorm:"primaryKey"` // References PublicServant
    Cname        string `gorm:"primaryKey"` // References Country
    DiseaseCode  string `gorm:"primaryKey"` // References Disease
    TotalDeaths  int
    TotalPatients int
}

func (Record) TableName() string {
    return "record"
}

func RecordPage(db *gorm.DB, tmpl *template.Template) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var records []Record
        var diseases []Disease
        var countries []Country
        var publicServants []PublicServant

        db.Find(&records)
        db.Find(&diseases)
        db.Find(&countries)
        db.Find(&publicServants)

        data := struct {
            Records       []Record
            Diseases      []Disease
            Countries     []Country
            PublicServants []PublicServant
        }{
            Records:       records,
            Diseases:      diseases,
            Countries:     countries,
            PublicServants: publicServants,
        }
        tmpl.Execute(w, data)
    }
}

func AddRecord(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
		totDeaths, err := strconv.Atoi(r.FormValue("total_deaths"))
		if err!= nil {
			totDeaths = 0
		}

		totPatients, err := strconv.Atoi(r.FormValue("total_patients"))
		if err!= nil {
			totPatients = 0
		}

        if r.Method == http.MethodPost {
            record := Record{
                Email:         r.FormValue("email"),
                Cname:         r.FormValue("cname"),
                DiseaseCode:   r.FormValue("disease_code"),
                TotalDeaths:   totDeaths,
                TotalPatients: totPatients,
            }
            db.Create(&record)
        }
        http.Redirect(w, r, "/record", http.StatusSeeOther)
    }
}

func DeleteRecord(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodPost {
            email := r.FormValue("email")
            cname := r.FormValue("cname")
            diseaseCode := r.FormValue("disease_code")
            db.Delete(&Record{}, "email = ? AND cname = ? AND disease_code = ?", email, cname, diseaseCode)
        }
        http.Redirect(w, r, "/record", http.StatusSeeOther)
    }
}

func EditRecordPage(db *gorm.DB, tmpl *template.Template) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodGet {
            email := r.URL.Query().Get("email")
            cname := r.URL.Query().Get("cname")
            diseaseCode := r.URL.Query().Get("disease_code")

            var record Record
            var diseases []Disease
            var countries []Country
            var publicServants []PublicServant

            db.First(&record, "email = ? AND cname = ? AND disease_code = ?", email, cname, diseaseCode)
            db.Find(&diseases)
            db.Find(&countries)
            db.Find(&publicServants)

            data := struct {
                Record        Record
                Diseases      []Disease
                Countries     []Country
                PublicServants []PublicServant
            }{
                Record:        record,
                Diseases:      diseases,
                Countries:     countries,
                PublicServants: publicServants,
            }
            tmpl.Execute(w, data)
        }
    }
}

func UpdateRecord(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
		totDeaths, err := strconv.Atoi(r.FormValue("total_deaths"))
		if err!= nil {
			totDeaths = 0
		}

		totPatients, err := strconv.Atoi(r.FormValue("total_patients"))
		if err!= nil {
			totPatients = 0
		}
        if r.Method == http.MethodPost {
            originalEmail := r.FormValue("original_email")
            originalCname := r.FormValue("original_cname")
            originalDiseaseCode := r.FormValue("original_disease_code")

            var record Record
            db.First(&record, "email = ? AND cname = ? AND disease_code = ?", originalEmail, originalCname, originalDiseaseCode)

            record.Email = r.FormValue("email")
            record.Cname = r.FormValue("cname")
            record.DiseaseCode = r.FormValue("disease_code")
            record.TotalDeaths = totDeaths
            record.TotalPatients = totPatients

            db.Save(&record)
        }
        http.Redirect(w, r, "/record", http.StatusSeeOther)
    }
}
