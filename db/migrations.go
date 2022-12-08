package db

import (
	"database/sql"
	"golang-final-project4-team2/resources/user_resources"
	"golang-final-project4-team2/utils/helpers"
	"log"
)

func Migrations(db *sql.DB) {

	createUsersTable(db)
	createCategoriesTable(db)
	createProductsTable(db)
	createTransactionHistoriesTable(db)
}

func createUsersTable(db *sql.DB) {
	createTable := `
	CREATE TABLE IF NOT EXISTS users (
    	id SERIAL PRIMARY KEY,
     	full_name VARCHAR(255) NOT NULL,
		email VARCHAR(255) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL,
		role VARCHAR(255) NOT NULL,
		balance INTEGER CONSTRAINT balance_constraint CHECK (balance >= 0 and balance <= 100000000) NOT NULL ,
		created_at timestamptz DEFAULT now(),
		updated_at timestamptz DEFAULT now(),
		deleted_at timestamptz
	) 
	`
	getData := `
	SELECT * FROM users
	`

	insertAdminData := `
	INSERT INTO users (full_name, email, password, role, balance)
	values ('admin', 'admin@gmail.com', $1, $2, 0)
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
		sold_product_amount INTEGER DEFAULT 0,
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

func createProductsTable(db *sql.DB) {
	createTable := `
	CREATE TABLE IF NOT EXISTS products (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		price INTEGER CONSTRAINT price_constraint CHECK (price >= 0 and price <= 50000000) NOT NULL,
		stock INTEGER CONSTRAINT stock_constraint CHECK (stock >= 5) NOT NULL,
		category_id SERIAL references categories(id),
		created_at timestamptz DEFAULT now(),
		updated_at timestamptz DEFAULT now()
	)
	`
	_, err = db.Exec(createTable)

	if err != nil {
		log.Fatal("Error creating products table:", err.Error())
	}
	log.Println("success creating products table")
}

func createTransactionHistoriesTable(db *sql.DB) {
	createTable := `
	CREATE TABLE IF NOT EXISTS transaction_histories
	(
		id          SERIAL PRIMARY KEY,
		product_id  SERIAL references products (id),
		user_id     SERIAL references users (id),
		quantity    INTEGER NOT NULL,
		total_price INTEGER NOT NULL,
		created_at  timestamptz DEFAULT now(),
		updated_at  timestamptz DEFAULT now()
	)
	`
	_, err = db.Exec(createTable)

	if err != nil {
		log.Fatal("Error creating transaction histories table:", err.Error())
	}
	log.Println("success creating transaction histories table")
}
