package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type Disease struct {
	DiseaseCode string `gorm:"primaryKey"`
	Pathogen    string
	Description string
	ID          int
}

func (Disease) TableName() string {
    return "disease"
}

// Displays the Disease page
func DiseasePage(db *gorm.DB, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var diseases []Disease
		var diseaseTypes []DiseaseType
		db.Find(&diseases)
		db.Find(&diseaseTypes)

		data := struct {
			Diseases     []Disease
			DiseaseTypes []DiseaseType
		}{
			Diseases:     diseases,
			DiseaseTypes: diseaseTypes,
		}
		tmpl.Execute(w, data)
	}
}

// Adds a new Disease
func AddDisease(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			id, err := strconv.Atoi(r.FormValue("id"))
			if err != nil {
				http.Error(w, "Invalid DiseaseType ID", http.StatusBadRequest)
				return
			}
			disease := Disease{
				DiseaseCode: r.FormValue("disease_code"),
				Pathogen:    r.FormValue("pathogen"),
				Description: r.FormValue("description"),
				ID:          id,
			}
			db.Create(&disease)
		}
		http.Redirect(w, r, "/disease", http.StatusSeeOther)
	}
}

// Deletes a Disease
func DeleteDisease(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			diseaseCode := r.FormValue("disease_code")
			db.Delete(&Disease{}, "disease_code = ?", diseaseCode)
		}
		http.Redirect(w, r, "/disease", http.StatusSeeOther)
	}
}

// Edits an existing Disease
func EditDisease(db *gorm.DB, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		diseaseCode := r.URL.Query().Get("disease_code")
		var disease Disease
		var diseaseTypes []DiseaseType

		if err := db.First(&disease, "disease_code = ?", diseaseCode).Error; err != nil {
			http.NotFound(w, r)
			return
		}

		db.Find(&diseaseTypes)

		data := struct {
			Disease      Disease
			DiseaseTypes []DiseaseType
		}{
			Disease:      disease,
			DiseaseTypes: diseaseTypes,
		}

		tmpl.Execute(w, data)
	}
}

// Updates a Disease
func UpdateDisease(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			id, err := strconv.Atoi(r.FormValue("id"))
			if err != nil {
				http.Error(w, "Invalid DiseaseType ID", http.StatusBadRequest)
				return
			}
			disease := Disease{
				DiseaseCode: r.FormValue("disease_code"),
				Pathogen:    r.FormValue("pathogen"),
				Description: r.FormValue("description"),
				ID:          id,
			}
			db.Model(&Disease{}).Where("disease_code = ?", r.FormValue("original_disease_code")).Updates(disease)
		}
		http.Redirect(w, r, "/disease", http.StatusSeeOther)
	}
}
