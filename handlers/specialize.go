package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type Specialize struct {
	ID          int    `gorm:"primaryKey"`
	DoctorEmail string `gorm:"column:email"`
}

func (Specialize) TableName() string {
    return "specialize"
}

func SpecializePage(db *gorm.DB, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var specializations []Specialize
		var doctors []Doctor
		var diseaseTypes []DiseaseType

		db.Find(&specializations)
		db.Find(&doctors)
		db.Find(&diseaseTypes)

		data := struct {
			Specializations []Specialize
			DiseaseTypes    []DiseaseType
			Doctors         []Doctor
		}{
			Specializations: specializations,
			DiseaseTypes:    diseaseTypes,
			Doctors:         doctors,
		}

		tmpl.Execute(w, data)
	}
}

func AddSpecialize(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			id, err := strconv.Atoi(r.FormValue("id"))
			if err != nil {
				http.Error(w, "Invalid Disease Type ID format", http.StatusBadRequest)
				return
			}
			doctorEmail := r.FormValue("doctor_email")

			specialize := Specialize{
				ID: id,
				DoctorEmail:   doctorEmail,
			}
			db.Create(&specialize)
		}
		http.Redirect(w, r, "/specialize", http.StatusSeeOther)
	}
}

func DeleteSpecialize(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			id := r.FormValue("id")
			db.Delete(&Specialize{}, "id = ?", id)
		}
		http.Redirect(w, r, "/specialize", http.StatusSeeOther)
	}
}

func EditSpecialize(db *gorm.DB, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		var specialize Specialize
		var doctors []Doctor
		var diseaseTypes []DiseaseType

		result := db.First(&specialize, "id = ?", id)
		db.Find(&doctors)
		db.Find(&diseaseTypes)

		if result.Error != nil {
			http.NotFound(w, r)
			return
		}

		data := struct {
			Specialize    Specialize
			Doctors       []Doctor
			DiseaseTypes  []DiseaseType
		}{
			Specialize:    specialize,
			Doctors:       doctors,
			DiseaseTypes:  diseaseTypes,
		}

		tmpl.Execute(w, data)
	}
}

func UpdateSpecialize(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			originalId, err := strconv.Atoi(r.FormValue("original_id"))
			if err != nil {
				http.Error(w, "Invalid ID format", http.StatusBadRequest)
				return
			}
			newId, err := strconv.Atoi(r.FormValue("id"))
			if err != nil {
				http.Error(w, "Invalid ID format", http.StatusBadRequest)
				return
			}
			doctorEmail := r.FormValue("doctor_email")

			db.Model(&Specialize{}).Where("id = ?", originalId).Updates(Specialize{ID: newId, DoctorEmail: doctorEmail})
		}
		http.Redirect(w, r, "/specialize", http.StatusSeeOther)
	}
}
