package handlers

import (
	"html/template"
	"net/http"

	"gorm.io/gorm"
)

type PatientDisease struct {
	Email       string `gorm:"primaryKey"`
	DiseaseCode string `gorm:"primaryKey"`
}

func (PatientDisease) TableName() string {
    return "patientdisease"
}

func PatientDiseasePage(db *gorm.DB, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var patientDiseases []PatientDisease
		db.Find(&patientDiseases)

		var patients []Patient
		db.Find(&patients)

		var diseases []Disease
		db.Find(&diseases)

		data := struct {
			PatientDiseases []PatientDisease
			Patients           []Patient
			Diseases        []Disease
		}{
			PatientDiseases: patientDiseases,
			Patients:           patients,
			Diseases:        diseases,
		}

		tmpl.Execute(w, data)
	}
}

func AddPatientDisease(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			patientDisease := PatientDisease{
				Email:       r.FormValue("email"),
				DiseaseCode: r.FormValue("disease_code"),
			}
			db.Create(&patientDisease)
		}
		http.Redirect(w, r, "/patientdisease", http.StatusSeeOther)
	}
}

func DeletePatientDisease(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			email := r.FormValue("email")
			diseaseCode := r.FormValue("disease_code")
			db.Delete(&PatientDisease{}, "email = ? AND disease_code = ?", email, diseaseCode)
		}
		http.Redirect(w, r, "/patientdisease", http.StatusSeeOther)
	}
}

func EditPatientDisease(db *gorm.DB, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := r.URL.Query().Get("email")
		diseaseCode := r.URL.Query().Get("disease_code")

		var patientDisease PatientDisease
		result := db.First(&patientDisease, "email = ? AND disease_code = ?", email, diseaseCode)
		if result.Error != nil {
			http.NotFound(w, r)
			return
		}


		var patients []Patient
		db.Find(&patients)

		var diseases []Disease
		db.Find(&diseases)

		data := struct {
			PatientDisease PatientDisease
			Patients          []Patient
			Diseases       []Disease
		}{
			PatientDisease: patientDisease,
			Patients:       patients,
			Diseases:       diseases,
		}

		tmpl.Execute(w, data)
	}
}


func UpdatePatientDisease(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			originalEmail := r.FormValue("original_email")
			newEmail := r.FormValue("email")
			newDiseaseCode := r.FormValue("disease_code")
			db.Model(&PatientDisease{}).Where("email = ?", originalEmail).Updates(PatientDisease{Email: newEmail, DiseaseCode: newDiseaseCode})
		}
		http.Redirect(w, r, "/patientdisease", http.StatusSeeOther)
	}
}
