package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "postgres"
	dbname   = "mygram"
	dialect  = "postgres"
)
var (
	db  *sql.DB
	err error
)

func handleDatabaseConnection() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err = sql.Open(dialect, psqlInfo)

	if err != nil {
		log.Panic("error occured while trying to validate database arguments:", err)
	}

	err = db.Ping()

	if err != nil {
		log.Panic("error occured while trying to connect to database:", err)
	}
}
func handleCreateRequiredTables() {
	userTable := `
		CREATE TABLE IF NOT EXISTS "user" (
			id SERIAL PRIMARY KEY,
			username VARCHAR(255) UNIQUE NOT NULL,
			email VARCHAR(255) UNIQUE NOT NULL,
			password TEXT NOT NULL,
			age int NOT NULL,
			created_at timestamptz DEFAULT now(),
			updated_at timestamptz DEFAULT now()
		);
	`

	photoTable := `
		CREATE TABLE IF NOT EXISTS "movies" (
			id SERIAL PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			caption TEXT NOT NULL,
			photo_url TEXT NOT NULL,
			user_id int NOT NULL,
			created_at timestamptz DEFAULT now(),
			updated_at timestamptz DEFAULT now(),
			CONSTRAINT photo_user_id_fk
				FOREIGN KEY(user_id)
					REFERENCES user(id)
						ON DELETE CASCADE
		);
	`
	commentTable := `
		CREATE TABLE IF NOT EXISTS "movies" (
			id SERIAL PRIMARY KEY,
			user_id int NOT NULL,
			photo_id int NOT NULL,
			message TEXT NOT NULL,
			created_at timestamptz DEFAULT now(),
			updated_at timestamptz DEFAULT now(),
			CONSTRAINT comment_user_id_fk
				FOREIGN KEY(user_id)
					REFERENCES user(id)
						ON DELETE CASCADE,
			CONSTRAINT comment_photo_user_id_fk
						FOREIGN KEY(photo_id)
							REFERENCES photo(id)
								ON DELETE CASCADE,
		);
	`
	socialMediaTable := `
		CREATE TABLE IF NOT EXISTS "socialmedia" (
			id SERIAL PRIMARY KEY,
			name varchar(255) NOT NULL,
			social_media_url TEXT NOT NULL,
			user_id int NOT NULL,
			created_at timestamptz DEFAULT now(),
			updated_at timestamptz DEFAULT now(),
			CONSTRAINT social_media_user_id_fk
				FOREIGN KEY(user_id)
					REFERENCES user(id)
						ON DELETE CASCADE
		);
	`

	createTableQueries := fmt.Sprintf("%s %s %s %s", userTable, photoTable, commentTable, socialMediaTable)

	_, err = db.Exec(createTableQueries)

	if err != nil {
		log.Panic("error occured while trying to create required tables:", err)
	}
}
