package handlers

import (
	// "fmt"
	"html/template"
	"net/http"

	"gorm.io/gorm"
)

type Patient struct {
	Email   string `gorm:"primaryKey"`
	User    User `gorm:"foreignKey:Email;constraint:OnDelete:CASCADE;"`
}

func PatientsPage(db *gorm.DB, tmpl *template.Template) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var patients []Patient
        var users []User
        db.Preload("User").Find(&patients)
        db.Find(&users)
        tmpl.ExecuteTemplate(w, "patients.html", struct {
            Patients []Patient
            Users    []User
        }{Patients: patients, Users: users})
    }
}

func AddPatient(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodPost {
            patient := Patient{
                Email:   r.FormValue("email"),
            }
            db.Create(&patient)
        }
        http.Redirect(w, r, "/patients", http.StatusSeeOther)
    }
}

func DeletePatient(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodPost {
            email := r.FormValue("email")
            db.Delete(&Patient{}, "email = ?", email)
        }
        http.Redirect(w, r, "/patients", http.StatusSeeOther)
    }
}

func EditPatient(db *gorm.DB, tmpl *template.Template) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        email := r.URL.Query().Get("email")
        var patient Patient
        // db.Preload("User").First(&patient, "email = ?", email)
        result := db.Preload("User").First(&patient, "email = ?", email)
		if result.Error != nil {
			http.NotFound(w, r)
			return
		}

        var users []User
		db.Find(&users)

        data := struct {
			Patient Patient
			Users          []User
		}{
			Patient: patient,
			Users:          users,
		}

        tmpl.Execute(w, data)

    }
}

func UpdatePatient(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodPost {
            originalEmail := r.FormValue("original_email")
            newEmail := r.FormValue("email")
            db.Model(&Patient{}).Where("email = ?", originalEmail).Updates(Patient{Email: newEmail})
        }
        http.Redirect(w, r, "/patients", http.StatusSeeOther)
    }
}
