package handlers

import (
	"html/template"
	"net/http"

	"gorm.io/gorm"
)

type PublicServant struct {
	Email     string `gorm:"primaryKey"`
	Department string // Options: Health, Research, Admin
	User      User   `gorm:"foreignKey:Email;constraint:OnDelete:CASCADE;"`
}

func (PublicServant) TableName() string {
    return "publicservant"
}

func PublicServantsPage(db *gorm.DB, tmpl *template.Template) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var publicServants []PublicServant
        var users []User
        db.Preload("User").Find(&publicServants)
        db.Find(&users)
        tmpl.ExecuteTemplate(w, "publicServants.html", struct {
            PublicServants []PublicServant
            Users    []User
        }{PublicServants: publicServants, Users: users})
    }
}

func AddPublicServant(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodPost {
            publicServant := PublicServant{
                Email:   r.FormValue("email"),
                Department: r.FormValue("department"),
            }
            db.Create(&publicServant)
        }
        http.Redirect(w, r, "/publicServants", http.StatusSeeOther)
    }
}

func DeletePublicServant(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodPost {
            email := r.FormValue("email")
            db.Delete(&PublicServant{}, "email = ?", email)
        }
        http.Redirect(w, r, "/publicServants", http.StatusSeeOther)
    }
}

func EditPublicServant(db *gorm.DB, tmpl *template.Template) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        email := r.URL.Query().Get("email")
        var publicServant PublicServant
        result := db.Preload("User").First(&publicServant, "email = ?", email)
        if result.Error != nil {
			http.NotFound(w, r)
			return
		}

        // fmt.Println("here")
        var users []User
		db.Find(&users)

        data := struct {
			PublicServant PublicServant
			Users          []User
		}{
			PublicServant: publicServant,
			Users:          users,
		}

		tmpl.Execute(w, data)

        // tmpl.ExecuteTemplate(w, "edit_publicServants.html", publicServant)
    }
}

func UpdatePublicServant(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodPost {
            originalEmail := r.FormValue("original_email")
            newEmail := r.FormValue("email")
            department := r.FormValue("department")
            db.Model(&PublicServant{}).Where("email = ?", originalEmail).Updates(PublicServant{Email: newEmail, Department: department})
        }
        http.Redirect(w, r, "/publicServants", http.StatusSeeOther)
    }
}