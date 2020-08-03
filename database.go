package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var database *sql.DB

func dbConnect() {
	user := "username"
	password := "password"
	dbName := "name"
	host := "localhost"
	port := "5432"
	connectionString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		user, password, host, port, dbName)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	database = db
	log.Println("Connected.")
}

// Find All cities and return JSON.
func dbCityList() (Cities, error) {
	var cities Cities
	var city City

	cityResults, err := database.Query(`SELECT id, "Name", "CountryCode", "District", "Population" FROM city`)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer cityResults.Close()

	for cityResults.Next() {
		cityResults.Scan(&city.ID, &city.Name, &city.CountryCode, &city.District, &city.Population)
		cities = append(cities, city)
	}

	return cities, nil
}

// Find a single city based on ID and returns as JSON
func dbCityDisplay(id int) (City, error) {
	var city City

	err := database.QueryRow(`SELECT id, "Name", "CountryCode", "District", "Population" FROM city WHERE id=$1 LIMIT 1`, id).Scan(&city.ID, &city.Name, &city.CountryCode, &city.District, &city.Population)
	if err != nil {
		return city, err
	}
	return city, nil
}

// Create a new city based on the information supplied.
func dbCityAdd(city City) (DBUpdate, error) {
	var addResult DBUpdate

	// Create prepared statement.
	stmt, err := database.Prepare(`INSERT INTO city("Name", "CountryCode", "District", "Population") VALUES($1, $2, $3, $4)`)
	if err != nil {
		return addResult, err
	}

	// Execute the prepared statement and retrieve the results.
	res, err := stmt.Exec(city.Name, city.CountryCode, city.District, city.Population)
	if err != nil {
		log.Println(err)
		return addResult, err
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
		return addResult, err
	}

	// Populate DBUpdate struct with Last Id and num rows affected.
	addResult.ID = rowCnt
	addResult.Affected = rowCnt

	return addResult, nil
}

// Delete the city with the supplied ID.
func dbCityDelete(id int64) (DBUpdate, error) {
	var deleteResult DBUpdate

	// Create prepared statement.
	stmt, err := database.Prepare(`DELETE FROM city WHERE id=$1`)
	if err != nil {
		log.Println(err)
		return deleteResult, err
	}

	// Execute the prepared statement and retrieve the results.
	res, err := stmt.Exec(id)
	if err != nil {
		log.Println(err)
		return deleteResult, err
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
		return deleteResult, err
	}

	// Populate DBUpdate struct with Last Id and num rows affected.
	deleteResult.ID = id
	deleteResult.Affected = rowCnt

	return deleteResult, nil
}

// Update the city with supplied ID.
func dbCityUpdate(id int64, city City) (DBUpdate, error) {
	var updateResult DBUpdate
	// Create prepared statement.
	stmt, err := database.Prepare(`UPDATE city
	SET "Name"=$1, "CountryCode"=$2, "Population"=$3, "District"=$4
	WHERE id=$5`)

	// Execute the prepared statement and retrieve the results.
	res, err := stmt.Exec(city.Name, city.CountryCode, city.Population, city.District, id)
	if err != nil {
		log.Println(err)
		return updateResult, err
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
		return updateResult, err
	}

	// Populate DBUpdate struct with Last Id and num rows affected.
	updateResult.ID = id
	updateResult.Affected = rowCnt

	return updateResult, nil
}
