package db

import (
	"database/sql"
	"golang-final-project3-team2/resources/user_resources"
	"golang-final-project3-team2/utils/helpers"
	"log"
)

func Migrations(db *sql.DB) {

	createUsersTable(db)
	createCategoriesTable(db)
	createTasksTable(db)
}

func createUsersTable(db *sql.DB) {
	createTable := `
	CREATE TABLE IF NOT EXISTS users (
    	id SERIAL PRIMARY KEY,
     	full_name VARCHAR(255) NOT NULL,
		email VARCHAR(255) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL,
		role VARCHAR(255) NOT NULL,
		created_at timestamptz DEFAULT now(),
		updated_at timestamptz DEFAULT now(),
		deleted_at timestamptz
	)
	`
	getData := `
	SELECT * FROM users
	`

	insertAdminData := `
	INSERT INTO users (full_name, email, password, role)
	values ('admin', 'admin@gmail.com', $1, $2)
	`

	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal("Error creating users table:", err.Error())
	}

	users, err := db.Query(getData)

	if err != nil {
		log.Fatal("Error querying users data:", err.Error())
	}

	if !users.Next() {
		hashPass, _ := helpers.HashPass("admin12")
		err = db.QueryRow(insertAdminData, hashPass, user_resources.RoleAdmin).Err()
		if err != nil {
			log.Fatal("Error seeding admin data:", err.Error())
		}
	}

	log.Println("success creating users table")
}

func createCategoriesTable(db *sql.DB) {
	createTable := `
	CREATE TABLE IF NOT EXISTS categories (
		id SERIAL PRIMARY KEY,
		type VARCHAR(255) NOT NULL,
		created_at timestamptz DEFAULT now(),
		updated_at timestamptz DEFAULT now()
	)
	`
	_, err = db.Exec(createTable)

	if err != nil {
		log.Fatal("Error creating categories table:", err.Error())
	}
	log.Println("success creating categories table")

}

func createTasksTable(db *sql.DB) {
	createTable := `
	CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		description TEXT NOT NULL,
		status BOOL NOT NULL,
		user_id SERIAL references users(id),
		category_id SERIAL references categories(id),
		created_at timestamptz DEFAULT now(),
		updated_at timestamptz DEFAULT now()
	)
	`
	_, err = db.Exec(createTable)

	if err != nil {
		log.Fatal("Error creating comments table:", err.Error())
	}
	log.Println("success creating comments table")

}
