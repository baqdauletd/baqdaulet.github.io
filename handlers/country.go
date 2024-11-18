package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type Country struct {
	Cname      string `gorm:"primaryKey"`
	Population int
}

func (Country) TableName() string {
    return "country"
}

// Displays the Country page
func CountryPage(db *gorm.DB, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var countries []Country
		db.Find(&countries)
		tmpl.Execute(w, countries)
	}
}

// Adds a new Country
func AddCountry(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			pop, err := strconv.Atoi(r.FormValue("population"))
			if err != nil {
				http.Error(w, "Invalid population format", http.StatusBadRequest)
				return
			}
			country := Country{
				Cname:      r.FormValue("cname"),
				Population: pop,
			}
			db.Create(&country)
		}
		http.Redirect(w, r, "/country", http.StatusSeeOther)
	}
}

// Deletes a Country
func DeleteCountry(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			cname := r.FormValue("cname")
			db.Delete(&Country{}, "cname = ?", cname)
		}
		http.Redirect(w, r, "/country", http.StatusSeeOther)
	}
}

// Edits an existing Country
func EditCountry(db *gorm.DB, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cname := r.URL.Query().Get("cname")
		var country Country
		result := db.First(&country, "cname = ?", cname)
		if result.Error != nil {
			http.NotFound(w, r)
			return
		}
		tmpl.Execute(w, country)
	}
}

// Updates a Country
func UpdateCountry(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			pop, err := strconv.Atoi(r.FormValue("population"))
			if err != nil {
				http.Error(w, "Invalid population format", http.StatusBadRequest)
				return
			}
			cname := r.FormValue("cname")
			country := Country{
				Cname:      cname,
				Population: pop,
			}
			db.Model(&Country{}).Where("cname = ?", r.FormValue("original_cname")).Updates(country)
		}
		http.Redirect(w, r, "/country", http.StatusSeeOther)
	}
}
