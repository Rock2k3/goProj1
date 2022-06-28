package repository

import (
	"app2/models"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func createDbConnection() *sql.DB  {
	datasourceUrl := os.Getenv("APP2_DATASOURCE_URL")
	fmt.Println(datasourceUrl)
	db, err := sql.Open("postgres", datasourceUrl)

	if err != nil {
		panic(err)
	}

	// check the connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	// return the connection
	return db
}

func GetUsers() ([]models.User, error)  {
	dbConn := createDbConnection()
	defer dbConn.Close()

	GET_USERS_QUERY := `select * from users`

	rows, err := dbConn.Query(GET_USERS_QUERY)
	defer rows.Close()

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	var users []models.User

	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.Id, &user.Name)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		users = append(users, user)

		fmt.Println(user)
	}

	return users, err
}

