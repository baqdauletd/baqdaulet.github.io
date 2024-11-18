package main

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	// time"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"as3/handlers"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// func main() {

// 	connStr := "host=localhost port=5432 user=posttest password=postgres dbname=csci341_ass3 sslmode=disable"
// 	db, err := sql.Open("postgres", connStr)
// 	if err != nil {
// 		fmt.Printf("Failed to connect to the database: %v", err)
// 	}
// 	defer db.Close()

// 	fmt.Println("Basic1")
// 	noBacteriaFunc(db)

// 	fmt.Println("Basic2")
// 	queryNonInfectious(db)

// 	fmt.Println("Basic3")
// 	queryDoctorsSpecializedInMoreThanTwoDiseases(db)

// 	fmt.Println("Basic4")
// 	queryAverageSalaryByCountryForVirology(db)

// 	fmt.Println("Basic5")
// 	queryDepartmentsWithCovidCases(db)

// 	//fmt.Println("Basic6")
// 	//doubleSalaryForCovidReporters(db)

// 	// fmt.Println("Basic7")
// 	// deleteUsersWithSpecificNames(db)

// 	// fmt.Println("Basic8")
// 	// createPrimaryIndexOnUsers(db)

// 	// fmt.Println("Basic9")
// 	// createSecondaryIndexOnDiseaseCode(db)

// 	fmt.Println("Basic10")
// 	queryTopCountriesByPatients(db)

// 	fmt.Println("Basic11")
// 	queryTotalCovidPatients(db)

// 	// fmt.Println("Basic12")
// 	// createPatientDiseaseView(db)

// 	fmt.Println("Basic13")
// 	queryPatientsFromView(db)

// }
func main() {
	// Initialize database
	dbHost := "dpg-cstj47t6147c73eli48g-a"
	dbUser := "posttest"
	dbPassword :="WdAf2MsIodvbj0vxd0qRS0oqfSS88Sy"
	dbName := "csci341_as3"
	dbPort := "5432"
	// dsn := "host=localhost user=posttest password=postgres dbname=csci341_ass3 port=5432 sslmode=disable"
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Parse templates
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

	// Register handlers
	http.HandleFunc("/users", handlers.UsersPage(db, tmpl.Lookup("users.html")))
	http.HandleFunc("/users/add", handlers.AddUser(db))
	http.HandleFunc("/users/delete", handlers.DeleteUser(db))
	http.HandleFunc("/users/edit", handlers.EditUserPage(db, editUserTmpl))
	http.HandleFunc("/users/update", handlers.UpdateUser(db))

	// http.HandleFunc("/patients", handlers.PatientsPage(db, tmpl))
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

	http.HandleFunc("/record", handlers.DiscoverPage(db, tmpl.Lookup("record.html")))
	http.HandleFunc("/record/add", handlers.AddDiscover(db))
	http.HandleFunc("/record/delete", handlers.DeleteDiscover(db))
	http.HandleFunc("/record/edit", handlers.EditDiscoverPage(db, tmplRecord))
	http.HandleFunc("/record/update", handlers.UpdateDiscover(db))
	
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

func noBacteriaFunc(db *sql.DB) {
	query := `
		SELECT d.disease_code, d.description
		FROM disease d
		JOIN discover di ON d.disease_code = di.disease_code
		WHERE d.pathogen = 'bacteria' AND di.first_enc_date < '2020-01-01';
		`

	rows, err := db.Query(query)
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

	fmt.Println("\nDiseases caused by 'bacteria' discovered before 2020:")
	for _, disease := range diseases {
		fmt.Printf("Disease Code: %s, Description: %s\n", disease.DiseaseCode, disease.Description)
	}
}

type Doctor struct {
	Name    string
	Surname string
	Degree  string
}

func queryNonInfectious(db *sql.DB) {
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

	rows2, err := db.Query(query2)
	if err != nil {
		fmt.Printf("Failed to execute query: %v", err)
	}
	defer rows2.Close()

	var doctors []Doctor
	for rows2.Next() {
		var doctor Doctor
		err := rows2.Scan(&doctor.Name, &doctor.Surname, &doctor.Degree)
		if err != nil {
			fmt.Printf("Failed to scan row: %v", err)
		}
		doctors = append(doctors, doctor)
	}

	fmt.Println("\nDoctors not specialized in 'infectious diseases':")
	for _, doctor := range doctors {
		fmt.Printf("Doctor: %s %s, Degree: %s\n", doctor.Name, doctor.Surname, doctor.Degree)
	}
}

func queryDoctorsSpecializedInMoreThanTwoDiseases(db *sql.DB) {
	query := `
		SELECT u.name, u.surname, d.degree
		FROM doctor d
		JOIN users u ON d.email = u.email
		JOIN specialize s ON d.email = s.email
		GROUP BY u.name, u.surname, d.degree
		HAVING COUNT(s.id) > 2;
	`

	rows, err := db.Query(query)
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

	fmt.Println("\nDoctors specialized in more than 2 disease types:")
	for _, doctor := range doctors {
		fmt.Printf("Doctor: %s %s, Degree: %s\n", doctor.Name, doctor.Surname, doctor.Degree)
	}
}

type CountryAvgSalary struct {
	Country       string
	AverageSalary float64
}

func queryAverageSalaryByCountryForVirology(db *sql.DB) {
	query := `
		SELECT u.cname, AVG(u.salary) AS avg_salary
		FROM doctor d
		JOIN users u ON d.email = u.email
		JOIN specialize s ON d.email = s.email
		JOIN diseasetype dt ON s.id = dt.id
		WHERE dt.description = 'Viral Diseases'
		GROUP BY u.cname;
	`

	rows, err := db.Query(query)
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

	fmt.Println("\nAverage salary of doctors specialized in 'virology' by country:")
	for _, result := range results {
		fmt.Printf("Country: %s, Average Salary: %.2f\n", result.Country, result.AverageSalary)
	}
}

type DepartmentReport struct {
	Department    string
	EmployeeCount int
}

func queryDepartmentsWithCovidCases(db *sql.DB) {
	query := `
		SELECT ps.department, COUNT(DISTINCT ps.email) AS employee_count
		FROM publicservant ps
		JOIN record r ON ps.email = r.email
		JOIN disease d ON r.disease_code = d.disease_code
		WHERE d.disease_code = 'COVID-19'
		GROUP BY ps.department;
	`

	rows, err := db.Query(query)
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

	fmt.Println("\nDepartments with public servants reporting 'covid-19' cases:")
	for _, dept := range departments {
		fmt.Printf("Department: %s, Employee Count: %d\n", dept.Department, dept.EmployeeCount)
	}
}

func doubleSalaryForCovidReporters(db *sql.DB) {
	query := `
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

	result, err := db.Exec(query)
	if err != nil {
		fmt.Printf("Failed to execute query: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("\nUpdated salaries for %d public servants.\n", rowsAffected)
}

func deleteUsersWithSpecificNames(db *sql.DB) {
	query := `
		DELETE FROM users
		WHERE name ILIKE '%bek%'OR name ILIKE '%gul%';
	`

	result, err := db.Exec(query)
	if err != nil {
		fmt.Printf("Failed to execute query: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("\nDeleted %d users with names containing 'bek' or 'gul'.\n", rowsAffected)
}

func createPrimaryIndexOnUsers(db *sql.DB) {
	query := `
		ALTER TABLE users
		ADD CONSTRAINT pk_users_email PRIMARY KEY (email);
	`

	_, err := db.Exec(query)
	if err != nil {
		fmt.Printf("Failed to create primary index: %v", err)
	} else {
		fmt.Println("\nPrimary index on 'email' field in Users table created successfully.")
	}
}

func createSecondaryIndexOnDiseaseCode(db *sql.DB) {
	query := `
		CREATE INDEX idx_disease_code
		ON disease (disease_code);
	`

	_, err := db.Exec(query)
	if err != nil {
		fmt.Printf("Failed to create secondary index: %v", err)
	} else {
		fmt.Println("\nSecondary index on 'disease_code' field in Disease table created successfully.")
	}
}

type CountryPatientCount struct {
	Country       string
	TotalPatients int
}

func queryTopCountriesByPatients(db *sql.DB) {
	query := `
		SELECT r.cname, SUM(r.total_patients) AS total_patients
		FROM record r
		GROUP BY r.cname
		ORDER BY total_patients DESC
		LIMIT 2;
	`

	rows, err := db.Query(query)
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

	fmt.Println("\nTop 2 countries with the highest number of total patients recorded:")
	for _, result := range results {
		fmt.Printf("Country: %s, Total Patients: %d\n", result.Country, result.TotalPatients)
	}
}

func queryTotalCovidPatients(db *sql.DB) {
	query := `
		SELECT SUM(r.total_patients) AS total_covid_patients
		FROM record r
		JOIN disease d ON r.disease_code = d.disease_code
		WHERE d.disease_code = 'COVID-19';
	`

	var totalPatients int
	err := db.QueryRow(query).Scan(&totalPatients)
	if err != nil {
		fmt.Printf("Failed to execute query: %v", err)
	}

	fmt.Printf("\nTotal number of patients with 'covid-19': %d\n", totalPatients)
}

func createPatientDiseaseView(db *sql.DB) {
	query := `
		CREATE VIEW patient_disease_view AS
		SELECT u.name, u.surname, d.description AS disease
		FROM patients p
		JOIN users u ON p.email = u.email
		JOIN patientdisease pd ON p.email = pd.email
		JOIN disease d ON pd.disease_code = d.disease_code;
	`

	_, err := db.Exec(query)
	if err != nil {
		fmt.Printf("Failed to create view: %v", err)
	} else {
		fmt.Println("\nView 'patient_disease_view' created successfully.")
	}
}

type PatientDisease struct {
	Name    string
	Surname string
	Disease string
}

func queryPatientsFromView(db *sql.DB) {
	query := `
		SELECT name, surname, disease
		FROM patient_disease_view;
	`

	rows, err := db.Query(query)
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

	fmt.Println("\nList of patients and their diseases from the view:")
	for _, result := range results {
		fmt.Printf("Patient: %s %s, Disease: %s\n", result.Name, result.Surname, result.Disease)
	}
}
