package handlers

import (
	"html/template"
	"net/http"

	"gorm.io/gorm"
	"strconv"
)

type User struct {
	Email   string `gorm:"primaryKey"`
	Name    string
	Surname string
	Salary  int
	Phone   string
	Cname   string
}

func UsersPage(db *gorm.DB, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var users []User
		db.Find(&users)
		tmpl.Execute(w, users)
	}
}

func AddUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			sal, err := strconv.Atoi(r.FormValue("salary"))
			if err!= nil {
        		sal = 0
  			}
			user := User{
				Email:   r.FormValue("email"),
				Name:    r.FormValue("name"),
				Surname: r.FormValue("surname"),
				Salary:  sal,
				Phone:   r.FormValue("phone"),
				Cname:   r.FormValue("cname"),
			}
			db.Create(&user)
		}
		http.Redirect(w, r, "/users", http.StatusSeeOther)
	}
}

func DeleteUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			email := r.FormValue("email")
			db.Delete(&User{}, "email = ?", email)
		}
		http.Redirect(w, r, "/users", http.StatusSeeOther)
	}
}

func EditUserPage(db *gorm.DB, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := r.URL.Query().Get("email")
		if email == "" {
			http.Error(w, "Email is required", http.StatusBadRequest)
			return
		}

		var user User
		result := db.First(&user, "email = ?", email)
		if result.Error != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		tmpl.Execute(w, user)
	}
}

func UpdateUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			originalEmail := r.FormValue("original_email")
			sal, err := strconv.Atoi(r.FormValue("salary"))
			if err != nil {
				sal = 0
			}

			updatedUser := User{
				Email:   r.FormValue("email"),
				Name:    r.FormValue("name"),
				Surname: r.FormValue("surname"),
				Salary:  sal,
				Phone:   r.FormValue("phone"),
				Cname:   r.FormValue("cname"),
			}

			result := db.Model(&User{}).Where("email = ?", originalEmail).Updates(updatedUser)
			if result.Error != nil {
				http.Error(w, "Unable to update user", http.StatusInternalServerError)
				return
			}
		}
		http.Redirect(w, r, "/users", http.StatusSeeOther)
	}
}

