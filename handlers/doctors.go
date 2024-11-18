package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"gorm.io/gorm"
)

type Doctor struct {
	Email  string `gorm:"primaryKey"`
	Degree string // Either "MD" or "PhD"
	User   User   `gorm:"foreignKey:Email;constraint:OnDelete:CASCADE;"`
}

func (Doctor) TableName() string {
    return "doctor"
}


func DoctorsPage(db *gorm.DB, tmpl *template.Template) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var doctors []Doctor
        var users []User
        db.Preload("User").Find(&doctors)
        db.Find(&users)
        tmpl.ExecuteTemplate(w, "doctors.html", struct {
            Doctors []Doctor
            Users    []User
        }{Doctors: doctors, Users: users})
    }
}

func AddDoctor(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodPost {
            doctor := Doctor{
                Email:   r.FormValue("email"),
                Degree: r.FormValue("degree"),
            }
            db.Create(&doctor)
        }
        http.Redirect(w, r, "/doctors", http.StatusSeeOther)
    }
}

func DeleteDoctor(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodPost {
            email := r.FormValue("email")
            db.Delete(&Doctor{}, "email = ?", email)
        }
        http.Redirect(w, r, "/doctors", http.StatusSeeOther)
    }
}

func EditDoctor(db *gorm.DB, tmpl *template.Template) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        email := r.URL.Query().Get("email")
        var doctor Doctor
        // db.Preload("User").First(&doctor, "email = ?", email)
        result := db.Preload("User").First(&doctor, "email = ?", email)
		if result.Error != nil {
			http.NotFound(w, r)
			return
		}

        // fmt.Println("here")
        var users []User
		db.Find(&users)

        data := struct {
			Doctor Doctor
			Users          []User
		}{
			Doctor: doctor,
			Users:          users,
		}

		tmpl.Execute(w, data)
        fmt.Print()
        // fmt.Printf("Doctor: %+v\n", doctor)
        // fmt.Printf("Users: %+v\n", users)


        // tmpl.ExecuteTemplate(w, "edit_doctors.html", struct {
        //     Doctor Doctor
        //     Users    []User
        // }{Doctor: doctor, Users: users})
        // tmpl.ExecuteTemplate(w, "edit_doctors.html", doctor)
    }
}

func UpdateDoctor(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodPost {
            originalEmail := r.FormValue("original_email")
            newEmail := r.FormValue("email")
            degree := r.FormValue("degree")
            db.Model(&Doctor{}).Where("email = ?", originalEmail).Updates(Doctor{Email: newEmail, Degree: degree})
        }
        http.Redirect(w, r, "/doctors", http.StatusSeeOther)
    }
}