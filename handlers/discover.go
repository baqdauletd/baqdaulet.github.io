package handlers

import (
    "html/template"
    "net/http"
    // "time"

    "gorm.io/gorm"
)

type Discover struct {
    Cname        string    `gorm:"primaryKey"` // References Country
    DiseaseCode  string    `gorm:"primaryKey"` // References Disease
    FirstEncDate string    // Date of discovery
}

func (Discover) TableName() string {
    return "discover"
}

func DiscoverPage(db *gorm.DB, tmpl *template.Template) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var discovers []Discover
        var diseases []Disease
        var countries []Country

        db.Find(&discovers)
        db.Find(&diseases)
        db.Find(&countries)

        data := struct {
            Discovers []Discover
            Diseases  []Disease
            Countries []Country
        }{
            Discovers: discovers,
            Diseases:  diseases,
            Countries: countries,
        }
        tmpl.Execute(w, data)
    }
}

func AddDiscover(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodPost {
            discover := Discover{
                Cname:        r.FormValue("cname"),
                DiseaseCode:  r.FormValue("disease_code"),
                FirstEncDate: r.FormValue("first_enc_date"),
            }
            db.Create(&discover)
        }
        http.Redirect(w, r, "/discover", http.StatusSeeOther)
    }
}

func DeleteDiscover(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodPost {
            cname := r.FormValue("cname")
            diseaseCode := r.FormValue("disease_code")
            db.Delete(&Discover{}, "cname = ? AND disease_code = ?", cname, diseaseCode)
        }
        http.Redirect(w, r, "/discover", http.StatusSeeOther)
    }
}

func EditDiscoverPage(db *gorm.DB, tmpl *template.Template) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodGet {
            cname := r.URL.Query().Get("cname")
            diseaseCode := r.URL.Query().Get("disease_code")

            var discover Discover
            var diseases []Disease
            var countries []Country

            db.First(&discover, "cname = ? AND disease_code = ?", cname, diseaseCode)
            db.Find(&diseases)
            db.Find(&countries)

            data := struct {
                Discover  Discover
                Diseases  []Disease
                Countries []Country
            }{
                Discover:  discover,
                Diseases:  diseases,
                Countries: countries,
            }
            tmpl.Execute(w, data)
        }
    }
}

func UpdateDiscover(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodPost {
            originalCname := r.FormValue("original_cname")
            // originalDiseaseCode := r.FormValue("original_disease_code")

            newCname := r.FormValue("cname")
            diseaseCode := r.FormValue("disease_code")
            first_enc_date := r.FormValue("first_enc_date")

            db.Model(&Discover{}).Where("cname = ?", originalCname).Updates(Discover{Cname: newCname, DiseaseCode: diseaseCode, FirstEncDate: first_enc_date})
        }
        http.Redirect(w, r, "/discover", http.StatusSeeOther)
    }
}
