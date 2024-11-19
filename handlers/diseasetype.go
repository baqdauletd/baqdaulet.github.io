package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type DiseaseType struct {
	ID          int    `gorm:"primaryKey"`
	Description string
}

func (DiseaseType) TableName() string {
    return "diseasetype"
}

func DiseaseTypePage(db *gorm.DB, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var diseaseTypes []DiseaseType
		db.Find(&diseaseTypes)
		tmpl.Execute(w, diseaseTypes)
	}
}


func AddDiseaseType(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			id, err := strconv.Atoi(r.FormValue("id"))
			if err != nil {
				http.Error(w, "Invalid ID format", http.StatusBadRequest)
				return
			}
			diseaseType := DiseaseType{
				ID:          id,
				Description: r.FormValue("description"),
			}
			db.Create(&diseaseType)
		}
		http.Redirect(w, r, "/diseasetype", http.StatusSeeOther)
	}
}


func DeleteDiseaseType(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			id := r.FormValue("id")
			db.Delete(&DiseaseType{}, "id = ?", id)
		}
		http.Redirect(w, r, "/diseasetype", http.StatusSeeOther)
	}
}


func EditDiseaseType(db *gorm.DB, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		var diseaseType DiseaseType
		result := db.First(&diseaseType, "id = ?", id)
		if result.Error != nil {
			http.NotFound(w, r)
			return
		}
		tmpl.Execute(w, diseaseType)
	}
}

func UpdateDiseaseType(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			id, err := strconv.Atoi(r.FormValue("id"))
			if err != nil {
				http.Error(w, "Invalid ID format", http.StatusBadRequest)
				return
			}
			diseaseType := DiseaseType{
				ID:          id,
				Description: r.FormValue("description"),
			}
			db.Model(&DiseaseType{}).Where("id = ?", id).Updates(diseaseType)
		}
		http.Redirect(w, r, "/diseasetype", http.StatusSeeOther)
	}
}
