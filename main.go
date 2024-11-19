package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"as3/handlers"
	"fmt"

	"database/sql"
	_ "github.com/lib/pq"
)

func main() {
	//databaseQueriesRun()
	appRun()
}

func databaseQueriesRun(){
	// connStr := "host=localhost port=5432 user=posttest password=postgres dbname=csci341_ass3 sslmode=disable"
	dsn := "host=dpg-csu4kkt2ng1s73cceer0-a.singapore-postgres.render.com user=posttest password=uzUSpYHW2eqPNd9MT8WyeqAkEAUM3bbX dbname=testdb_q0pg port=5432 sslmode=require"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		fmt.Printf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	fmt.Println("Query1")
	Basic1(db)

	fmt.Println("Query2")
	Basic2(db)

	fmt.Println("Query3")
	Basic3(db)

	fmt.Println("Query4")
	Basic4(db)

	fmt.Println("Query5")
	Basic5(db)

	fmt.Println("Query6")
	Basic6(db)

	fmt.Println("Query7")
	Basic7(db)

	fmt.Println("Query8")
	Basic8(db)

	fmt.Println("Query8")
	Basic9(db)

	fmt.Println("Query10")
	Basic10(db)

	fmt.Println("Query11")
	Basic11(db)

	fmt.Println("Query12")
	Basic12(db)

	fmt.Println("Query13")
	Basic13(db)
}

func appRun(){
	//dsn := "host=localhost user=posttest password=uzUSpYHW2eqPNd9MT8WyeqAkEAUM3bbX dbname=testdb_q0pg port=5432 sslmode=require"
	dsn := "host=dpg-csu4kkt2ng1s73cceer0-a.singapore-postgres.render.com user=posttest password=uzUSpYHW2eqPNd9MT8WyeqAkEAUM3bbX dbname=testdb_q0pg port=5432 sslmode=require"


	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	tmpl := template.Must(template.ParseGlob(filepath.Join("templates", "*.html")))
	editUserTmpl := template.Must(template.ParseFiles("templates/edit_users.html"))
	editPatientTmpl := template.Must(template.ParseFiles("templates/edit_patients.html"))
	editDoctorTmpl := template.Must(template.ParseFiles("templates/edit_doctors.html"))
	editPublicServantTmpl := template.Must(template.ParseFiles("templates/edit_publicServants.html"))
	tmplDiseaseType := template.Must(template.ParseFiles("templates/edit_diseasetype.html"))
	tmplDisease := template.Must(template.ParseFiles("templates/edit_disease.html"))
	tmplCountry := template.Must(template.ParseFiles("templates/edit_country.html"))
	tmplDiscover := template.Must(template.ParseFiles("templates/edit_discover.html"))
	tmplRecord := template.Must(template.ParseFiles("templates/edit_record.html"))
	tmplSpecialize := template.Must(template.ParseFiles("templates/edit_specialize.html"))
	tmplPatientDisease := template.Must(template.ParseFiles("templates/edit_patientdisease.html"))

	indexTmpl := template.Must(template.ParseFiles("templates/index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := indexTmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			log.Println("Template rendering error:", err)
		}
	})

	// handlers
	http.HandleFunc("/users", handlers.UsersPage(db, tmpl.Lookup("users.html")))
	http.HandleFunc("/users/add", handlers.AddUser(db))
	http.HandleFunc("/users/delete", handlers.DeleteUser(db))
	http.HandleFunc("/users/edit", handlers.EditUserPage(db, editUserTmpl))
	http.HandleFunc("/users/update", handlers.UpdateUser(db))

	http.HandleFunc("/patients", handlers.PatientsPage(db, tmpl.Lookup("patients.html")))
    http.HandleFunc("/patients/add", handlers.AddPatient(db))
    http.HandleFunc("/patients/delete", handlers.DeletePatient(db))
	http.HandleFunc("/patients/edit", handlers.EditPatient(db, editPatientTmpl))
	http.HandleFunc("/patients/update", handlers.UpdatePatient(db))

    http.HandleFunc("/doctors", handlers.DoctorsPage(db, tmpl.Lookup("doctors.html")))
    http.HandleFunc("/doctors/add", handlers.AddDoctor(db))
    http.HandleFunc("/doctors/delete", handlers.DeleteDoctor(db))
	http.HandleFunc("/doctors/edit", handlers.EditDoctor(db, editDoctorTmpl))
	http.HandleFunc("/doctors/update", handlers.UpdateDoctor(db))

    http.HandleFunc("/publicServants", handlers.PublicServantsPage(db, tmpl.Lookup("publicServants.html")))
    http.HandleFunc("/publicServants/add", handlers.AddPublicServant(db))
    http.HandleFunc("/publicServants/delete", handlers.DeletePublicServant(db))
	http.HandleFunc("/publicServants/edit", handlers.EditPublicServant(db, editPublicServantTmpl))
	http.HandleFunc("/publicServants/update", handlers.UpdatePublicServant(db))

	http.HandleFunc("/diseasetype", handlers.DiseaseTypePage(db, tmpl.Lookup("diseasetype.html")))
	http.HandleFunc("/diseasetype/add", handlers.AddDiseaseType(db))
	http.HandleFunc("/diseasetype/delete", handlers.DeleteDiseaseType(db))
	http.HandleFunc("/diseasetype/edit", handlers.EditDiseaseType(db, tmplDiseaseType))
	http.HandleFunc("/diseasetype/update", handlers.UpdateDiseaseType(db))

	http.HandleFunc("/disease", handlers.DiseasePage(db, tmpl.Lookup("disease.html")))
	http.HandleFunc("/disease/add", handlers.AddDisease(db))
	http.HandleFunc("/disease/delete", handlers.DeleteDisease(db))
	http.HandleFunc("/disease/edit", handlers.EditDisease(db, tmplDisease))
	http.HandleFunc("/disease/update", handlers.UpdateDisease(db))

	http.HandleFunc("/country", handlers.CountryPage(db, tmpl.Lookup("country.html")))
	http.HandleFunc("/country/add", handlers.AddCountry(db))
	http.HandleFunc("/country/delete", handlers.DeleteCountry(db))
	http.HandleFunc("/country/edit", handlers.EditCountry(db, tmplCountry))
	http.HandleFunc("/country/update", handlers.UpdateCountry(db))

	http.HandleFunc("/discover", handlers.DiscoverPage(db, tmpl.Lookup("discover.html")))
	http.HandleFunc("/discover/add", handlers.AddDiscover(db))
	http.HandleFunc("/discover/delete", handlers.DeleteDiscover(db))
	http.HandleFunc("/discover/edit", handlers.EditDiscoverPage(db, tmplDiscover))
	http.HandleFunc("/discover/update", handlers.UpdateDiscover(db))

	http.HandleFunc("/record", handlers.RecordPage(db, tmpl.Lookup("record.html")))
	http.HandleFunc("/record/add", handlers.AddRecord(db))
	http.HandleFunc("/record/delete", handlers.DeleteRecord(db))
	http.HandleFunc("/record/edit", handlers.EditRecordPage(db, tmplRecord))
	http.HandleFunc("/record/update", handlers.UpdateRecord(db))
	
	http.HandleFunc("/specialize", handlers.SpecializePage(db, tmpl.Lookup("specialize.html")))
	http.HandleFunc("/specialize/add", handlers.AddSpecialize(db))
	http.HandleFunc("/specialize/delete", handlers.DeleteSpecialize(db))
	http.HandleFunc("/specialize/edit", handlers.EditSpecialize(db, tmplSpecialize))
	http.HandleFunc("/specialize/update", handlers.UpdateSpecialize(db))
	
	http.HandleFunc("/patientdisease", handlers.PatientDiseasePage(db, tmpl.Lookup("patientdisease.html")))
	http.HandleFunc("/patientdisease/add", handlers.AddPatientDisease(db))
	http.HandleFunc("/patientdisease/delete", handlers.DeletePatientDisease(db))
	http.HandleFunc("/patientdisease/edit", handlers.EditPatientDisease(db, tmplPatientDisease))
	http.HandleFunc("/patientdisease/update", handlers.UpdatePatientDisease(db))
	
	log.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}


type Disease struct {
	DiseaseCode string
	Description string
}
type Doctor struct {
	Name    string
	Surname string
	Degree  string
}
type CountryAvgSalary struct {
	Country       string
	AverageSalary float64
}
type DepartmentReport struct {
	Department    string
	EmployeeCount int
}

type CountryPatientCount struct {
	Country       string
	TotalPatients int
}
type PatientDisease struct {
	Name    string
	Surname string
	Disease string
}

func queries(qu string) string{
	query1 := `
		SELECT d.disease_code, d.description
		FROM disease d
		JOIN discover di ON d.disease_code = di.disease_code
		WHERE d.pathogen = 'Bacteria' AND di.first_enc_date < '2020-01-01';
	`

	query2 := `
		SELECT DISTINCT u.name, u.surname, d.degree
		FROM doctor d
		JOIN users u ON d.email = u.email
		LEFT JOIN specialize s ON d.email = s.email
		LEFT JOIN diseasetype dt ON s.id = dt.id
		WHERE d.email NOT IN (
			SELECT DISTINCT d.email
			FROM doctor d
			JOIN specialize s ON d.email = s.email
			JOIN diseasetype dt ON s.id = dt.id
			WHERE dt.description = 'Infectious Diseases'
		);
	`

	query3 := `
		SELECT u.name, u.surname, d.degree
		FROM doctor d
		JOIN users u ON d.email = u.email
		JOIN specialize s ON d.email = s.email
		GROUP BY u.name, u.surname, d.degree
		HAVING COUNT(s.id) > 2;
	`

	query4 := `
		SELECT u.cname, AVG(u.salary) AS avg_salary
		FROM doctor d
		JOIN users u ON d.email = u.email
		JOIN specialize s ON d.email = s.email
		JOIN diseasetype dt ON s.id = dt.id
		WHERE dt.description = 'Viral Diseases'
		GROUP BY u.cname;
	`

	query5 := `
		SELECT ps.department, COUNT(DISTINCT ps.email) AS employee_count
		FROM publicservant ps
		JOIN record r ON ps.email = r.email
		JOIN disease d ON r.disease_code = d.disease_code
		WHERE d.disease_code = 'COVID-19'
		GROUP BY ps.department;
	`

	query6 := `
		UPDATE users
		SET salary = salary * 2
		WHERE email IN (
			SELECT ps.email
			FROM publicservant ps
			JOIN record r ON ps.email = r.email
			JOIN disease d ON r.disease_code = d.disease_code
			WHERE d.disease_code = 'COVID-19'
			GROUP BY ps.email
			HAVING SUM(r.total_patients) > 3
		);
	`

	query7 := `
		DELETE FROM users
		WHERE name ILIKE '%bek%'OR name ILIKE '%gul%';
	`

	query8 := `
		ALTER TABLE users
		ADD CONSTRAINT pk_users_email PRIMARY KEY (email);
	`

	query9 := `
		CREATE INDEX idx_disease_code
		ON disease (disease_code);
	`

	query10 := `
		SELECT r.cname, SUM(r.total_patients) AS total_patients
		FROM record r
		GROUP BY r.cname
		ORDER BY total_patients DESC
		LIMIT 2;
	`

	query11 := `
		SELECT SUM(r.total_patients) AS total_covid_patients
		FROM record r
		JOIN disease d ON r.disease_code = d.disease_code
		WHERE d.disease_code = 'COVID-19';
	`
	query12 := `
		CREATE VIEW patient_disease_view AS
		SELECT u.name, u.surname, d.description AS disease
		FROM patients p
		JOIN users u ON p.email = u.email
		JOIN patientdisease pd ON p.email = pd.email
		JOIN disease d ON pd.disease_code = d.disease_code;
	`
	query13 := `
		SELECT name, surname, disease
		FROM patient_disease_view;
	`
	
	if qu == "query1"{
		return query1
	} else if qu == "query2"{
		return query2
	} else if qu == "query3"{
		return query3 
	} else if qu == "query4"{
		return query4
	} else if qu == "query5"{
		return query5
	} else if qu == "query6"{
		return query6
	} else if qu == "query7"{
		return query7
	} else if qu == "query8"{
		return query8
	} else if qu == "query9"{
		return query9
	} else if qu == "query10"{
		return query10
	} else if qu == "query11"{
		return query11
	} else if qu == "query12"{
		return query12
	} else if qu == "query13"{
		return query13
	}
	
	return qu
}


func Basic1(db *sql.DB) {
	query1 := queries("query1")
	rows, err := db.Query(query1)
	if err != nil {
		fmt.Printf("Failed to execute query: %v", err)
	}
	defer rows.Close()


	var diseases []Disease
	for rows.Next() {
		var disease Disease
		err := rows.Scan(&disease.DiseaseCode, &disease.Description)
		if err != nil {
			fmt.Printf("Failed to scan row: %v", err)
		}
		diseases = append(diseases, disease)
	}

	if err := rows.Err(); err != nil {
		fmt.Printf("Error iterating over rows: %v", err)
	}

	fmt.Println("Diseases caused by 'bacteria' discovered before 2020:")
	for _, disease := range diseases {
		fmt.Printf("Disease Code: %s, Description: %s\n", disease.DiseaseCode, disease.Description)
	}
	fmt.Println()
}


func Basic2(db *sql.DB) {
	query2 := queries("query2")
	rows, err := db.Query(query2)
	if err != nil {
		fmt.Printf("Failed to execute query: %v", err)
	}
	defer rows.Close()

	var doctors []Doctor
	for rows.Next() {
		var doctor Doctor
		err := rows.Scan(&doctor.Name, &doctor.Surname, &doctor.Degree)
		if err != nil {
			fmt.Printf("Failed to scan row: %v", err)
		}
		doctors = append(doctors, doctor)
	}

	fmt.Println("Doctors not specialized in 'infectious diseases':")
	for _, doctor := range doctors {
		fmt.Printf("Doctor: %s %s, Degree: %s\n", doctor.Name, doctor.Surname, doctor.Degree)
	}
	fmt.Println()
}

func Basic3(db *sql.DB) {
	query3 := queries("query3")
	rows, err := db.Query(query3)
	if err != nil {
		fmt.Printf("Failed to execute query: %v", err)
	}
	defer rows.Close()

	var doctors []Doctor
	for rows.Next() {
		var doctor Doctor
		err := rows.Scan(&doctor.Name, &doctor.Surname, &doctor.Degree)
		if err != nil {
			fmt.Printf("Failed to scan row: %v", err)
		}
		doctors = append(doctors, doctor)
	}

	fmt.Println("Doctors specialized in more than 2 disease types:")
	for _, doctor := range doctors {
		fmt.Printf("Doctor: %s %s, Degree: %s\n", doctor.Name, doctor.Surname, doctor.Degree)
	}
	fmt.Println()
}


func Basic4(db *sql.DB) {
	query4 := queries("query4")
	rows, err := db.Query(query4)
	if err != nil {
		fmt.Printf("Failed to execute query: %v", err)
	}
	defer rows.Close()

	var results []CountryAvgSalary
	for rows.Next() {
		var result CountryAvgSalary
		err := rows.Scan(&result.Country, &result.AverageSalary)
		if err != nil {
			fmt.Printf("Failed to scan row: %v", err)
		}
		results = append(results, result)
	}

	fmt.Println("Average salary of doctors specialized in 'virology' by country:")
	for _, result := range results {
		fmt.Printf("Country: %s, Average Salary: %.2f\n", result.Country, result.AverageSalary)
	}
	fmt.Println()
}


func Basic5(db *sql.DB) {
	query5 := queries("query5")
	rows, err := db.Query(query5)
	if err != nil {
		fmt.Printf("Failed to execute query: %v", err)
	}
	defer rows.Close()

	var departments []DepartmentReport
	for rows.Next() {
		var dept DepartmentReport
		err := rows.Scan(&dept.Department, &dept.EmployeeCount)
		if err != nil {
			fmt.Printf("Failed to scan row: %v", err)
		}
		departments = append(departments, dept)
	}

	fmt.Println("Departments with public servants reporting 'covid-19' cases:")
	for _, dept := range departments {
		fmt.Printf("Department: %s, Employee Count: %d\n", dept.Department, dept.EmployeeCount)
	}
	fmt.Println()
}

func Basic6(db *sql.DB) {
	query6 := queries("query6")
	result, err := db.Exec(query6)
	if err != nil {
		fmt.Printf("Failed to execute query: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("Updated salaries for %d public servants.\n", rowsAffected)
	fmt.Println()
}

func Basic7(db *sql.DB) {
	query7 := queries("query7")
	result, err := db.Exec(query7)
	if err != nil {
		fmt.Printf("Failed to execute query: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("Deleted %d users with names containing 'bek' or 'gul'.\n", rowsAffected)
	fmt.Println()
}

func Basic8(db *sql.DB) {
	query8 := queries("query8")
	_, err := db.Exec(query8)
	if err != nil {
		fmt.Printf("Failed to create primary index: %v", err)
	} else {
		fmt.Println("Primary index on 'email' field in Users table created successfully.")
	}
	fmt.Println()
}

func Basic9(db *sql.DB) {
	query9 := queries("query9")
	_, err := db.Exec(query9)
	if err != nil {
		fmt.Printf("Failed to create secondary index: %v", err)
	} else {
		fmt.Println("Secondary index on 'disease_code' field in Disease table created successfully.")
	}
	fmt.Println()
}

func Basic10(db *sql.DB) {
	query10 := queries("query10")
	rows, err := db.Query(query10)
	if err != nil {
		fmt.Printf("Failed to execute query: %v", err)
	}
	defer rows.Close()

	var results []CountryPatientCount
	for rows.Next() {
		var result CountryPatientCount
		err := rows.Scan(&result.Country, &result.TotalPatients)
		if err != nil {
			fmt.Printf("Failed to scan row: %v", err)
		}
		results = append(results, result)
	}

	fmt.Println("Top 2 countries with the highest number of total patients recorded:")
	for _, result := range results {
		fmt.Printf("Country: %s, Total Patients: %d\n", result.Country, result.TotalPatients)
	}
	fmt.Println()
}


func Basic11(db *sql.DB) {
	query11 := queries("query11")
	var totalPatients int
	err := db.QueryRow(query11).Scan(&totalPatients)
	if err != nil {
		fmt.Printf("Failed to execute query: %v", err)
	}

	fmt.Printf("Total number of patients with 'covid-19': %d\n", totalPatients)
	fmt.Println()
}

func Basic12(db *sql.DB) {
	query12 := queries("query12")
	_, err := db.Exec(query12)
	if err != nil {
		fmt.Printf("Failed to create view: %v", err)
		fmt.Println()
	} else {
		fmt.Println("View 'patient_disease_view' created successfully.")
	}
	fmt.Println()
}

func Basic13(db *sql.DB) {
	query13 := queries("query13")
	rows, err := db.Query(query13)
	if err != nil {
		fmt.Printf("Failed to execute query: %v", err)
	}
	defer rows.Close()

	var results []PatientDisease
	for rows.Next() {
		var result PatientDisease
		err := rows.Scan(&result.Name, &result.Surname, &result.Disease)
		if err != nil {
			fmt.Printf("Failed to scan row: %v", err)
		}
		results = append(results, result)
	}

	fmt.Println("List of patients and their diseases from the view:")
	for _, result := range results {
		fmt.Printf("Patient: %s %s, Disease: %s\n", result.Name, result.Surname, result.Disease)
	}
	fmt.Println()
}
