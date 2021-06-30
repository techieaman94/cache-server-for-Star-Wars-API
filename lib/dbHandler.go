package cacheServer

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

func CreateTablesIfNotExists() {
	db, _ := sql.Open("sqlite3", "./cacheDatabase.db")

	// -----people table -----------------------------------
	createPeopleTableSQL := `CREATE TABLE IF NOT EXISTS people(
	  id INT PRIMARY KEY     NOT NULL,
		name TEXT    NOT NULL,
		height TEXT    NOT NULL,
		mass TEXT    NOT NULL,
		hair_color TEXT    NOT NULL,
		skin_color TEXT    NOT NULL,
		eye_color TEXT    NOT NULL,
		birth_year TEXT    NOT NULL,
		gender TEXT    NOT NULL,
		homeworld TEXT    NOT NULL,
		films TEXT    NOT NULL,
		species TEXT    NOT NULL,
		vehicles TEXT    NOT NULL,
		starships TEXT    NOT NULL,
		created TEXT    NOT NULL,
		edited TEXT    NOT NULL,
		url TEXT    NOT NULL
		);` // SQL Statement for Create Table

	statement1, err := db.Prepare(createPeopleTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement1.Exec() // Execute SQL Statements
	log.Println("People table created")

	// -----planets table -----------------------------------
	createPlanetsTableSQL := `CREATE TABLE IF NOT EXISTS planets(
   	id INT PRIMARY KEY     NOT NULL,
		name TEXT    NOT NULL,
		rotation_period TEXT    NOT NULL,
		orbital_period TEXT    NOT NULL,
		diameter TEXT    NOT NULL,
		climate TEXT    NOT NULL,
		gravity TEXT    NOT NULL,
		terrain TEXT    NOT NULL,
		surface_water TEXT    NOT NULL,
		population TEXT    NOT NULL,
		residents TEXT    NOT NULL,
		films TEXT    NOT NULL,
		created TEXT    NOT NULL,
		edited TEXT    NOT NULL,
		url TEXT    NOT NULL
		);` // SQL Statement for Create Table

	statement2, err := db.Prepare(createPlanetsTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement2.Exec() // Execute SQL Statements
	log.Println("Planets table created")

	// -----films table -----------------------------------
	createFilmsTableSQL := `CREATE TABLE IF NOT EXISTS films(
   	id INT PRIMARY KEY     NOT NULL,
		title TEXT    NOT NULL,
		episode_id TEXT    NOT NULL,
		opening_crawl TEXT    NOT NULL,
		director TEXT    NOT NULL,
		producer TEXT    NOT NULL,
		release_date TEXT    NOT NULL,
		characters TEXT    NOT NULL,
		planets TEXT    NOT NULL,
		starships TEXT    NOT NULL,
		vehicles TEXT    NOT NULL,
		species TEXT    NOT NULL,
		created TEXT    NOT NULL,
		edited TEXT    NOT NULL,
		url TEXT    NOT NULL
		);` // SQL Statement for Create Table

	// log.Println("Create student table...")
	statement3, err := db.Prepare(createFilmsTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement3.Exec() // Execute SQL Statements
	log.Println("Films table created")

	// -----species table -----------------------------------
	createSpeciesTableSQL := `CREATE TABLE IF NOT EXISTS species(
   	id INT PRIMARY KEY     NOT NULL,
		name TEXT    NOT NULL,
		classification TEXT    NOT NULL,
		designation TEXT    NOT NULL,
		average_height TEXT    NOT NULL,
		skin_colors TEXT    NOT NULL,
		hair_colors TEXT    NOT NULL,
		eye_colors TEXT    NOT NULL,
		average_lifespan TEXT    NOT NULL,
		homeworld TEXT    NOT NULL,
		language TEXT    NOT NULL,
		people TEXT    NOT NULL,
		films TEXT    NOT NULL,
		created TEXT    NOT NULL,
		edited TEXT    NOT NULL,
		url TEXT    NOT NULL
		);` // SQL Statement for Create Table

	statement4, err := db.Prepare(createSpeciesTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement4.Exec() // Execute SQL Statements
	log.Println("Species table created")

	// -----starships table -----------------------------------
	createSpaceshipsTableSQL := `CREATE TABLE IF NOT EXISTS starships(
   	id INT PRIMARY KEY     NOT NULL,
		name TEXT    NOT NULL,
		model TEXT    NOT NULL,
		manufacturer TEXT    NOT NULL,
		cost_in_credits TEXT    NOT NULL,
		length TEXT    NOT NULL,
		max_atmosphering_speed TEXT    NOT NULL,
		crew TEXT    NOT NULL,
		passengers TEXT    NOT NULL,
		cargo_capacity TEXT    NOT NULL,
		consumables TEXT    NOT NULL,
		hyperdrive_rating TEXT    NOT NULL,
		MGLT TEXT    NOT NULL,
		starship_class TEXT NOT NULL,
		pilots TEXT NOT NULL,
		films TEXT NOT NULL,
		created TEXT    NOT NULL,
		edited TEXT    NOT NULL,
		url TEXT    NOT NULL
		);` // SQL Statement for Create Table

	// log.Println("Create student table...")
	statement5, err := db.Prepare(createSpaceshipsTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement5.Exec() // Execute SQL Statements
	log.Println("Spaceships table created")

	// -----vehicles table -----------------------------------
	createVehiclesTableSQL := `CREATE TABLE IF NOT EXISTS vehicles(
   	id INT PRIMARY KEY     NOT NULL,
		name TEXT    NOT NULL,
		model TEXT    NOT NULL,
		manufacturer TEXT    NOT NULL,
		cost_in_credits TEXT    NOT NULL,
		length TEXT    NOT NULL,
		max_atmosphering_speed TEXT    NOT NULL,
		crew TEXT    NOT NULL,
		passengers TEXT    NOT NULL,
		cargo_capacity TEXT    NOT NULL,
		consumables TEXT    NOT NULL,
		vehicle_class TEXT    NOT NULL,
		pilots TEXT NOT NULL,
		films TEXT NOT NULL,
		created TEXT    NOT NULL,
		edited TEXT    NOT NULL,
		url TEXT    NOT NULL
		);` // SQL Statement for Create Table

	// log.Println("Create student table...")
	statement6, err := db.Prepare(createVehiclesTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement6.Exec() // Execute SQL Statements
	log.Println("Vehicles table created")

}

func SearchInLocalPeople(id string) People {
	fmt.Println("Hit Point : SearchInLocalPeople ", id)
	var people People
	db, _ := sql.Open("sqlite3", "./cacheDatabase.db")

	stmt, err := db.Prepare("SELECT * FROM people WHERE id = ?")
	if err != nil {
		fmt.Println("Error 1 SearchInLocal : ", err.Error())
	}

	row := stmt.QueryRow(id)

	var films, species, vehicles, starships string

	switch err := row.Scan(&id,
		&people.Name,
		&people.Height,
		&people.Mass,
		&people.HairColor,
		&people.SkinColor,
		&people.EyeColor,
		&people.BirthYear,
		&people.Gender,
		&people.Homeworld,
		&films,
		&species,
		&vehicles,
		&starships,
		&people.Created,
		&people.Edited,
		&people.Url); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:

		if len(films) > 0 {
			people.Films = strings.Split(films, " ")
		} else {
			people.Films = []string{}
		}

		if len(species) > 0 {
			people.Species = strings.Split(species, " ")
		} else {
			people.Species = []string{}
		}

		if len(vehicles) > 0 {
			people.Vehicles = strings.Split(vehicles, " ")
		} else {
			people.Vehicles = []string{}
		}
		if len(starships) > 0 {
			people.Starships = strings.Split(starships, " ")
		} else {
			people.Starships = []string{}
		}

		fmt.Println("Result from ROW : ", people)
		// byteString := []byte(people)
		return people
	default:
		panic(err)
	}
	return people

}

func SearchInLocalPlanets(id string) Planet {
	fmt.Println("Hit Point : SearchInLocalPlanets ", id)
	var planet Planet
	db, _ := sql.Open("sqlite3", "./cacheDatabase.db")

	stmt, err := db.Prepare("SELECT * FROM planets WHERE id = ?")
	if err != nil {
		fmt.Println("Error 1 SearchInLocal : ", err.Error())
	}

	row := stmt.QueryRow(id)

	var films, residents string

	switch err := row.Scan(&id,
		&planet.Name,
		&planet.RotationPeriod,
		&planet.OrbitalPeriod,
		&planet.Diameter,
		&planet.Climate,
		&planet.Gravity,
		&planet.Terrain,
		&planet.SurfaceWater,
		&planet.Population,
		&residents,
		&films,
		&planet.Created,
		&planet.Edited,
		&planet.Url); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:

		if len(films) > 0 {
			planet.Films = strings.Split(films, " ")
		} else {
			planet.Films = []string{}
		}

		if len(residents) > 0 {
			planet.Residents = strings.Split(residents, " ")
		} else {
			planet.Residents = []string{}
		}

		fmt.Println("Result from ROW : ", planet)
		// byteString := []byte(people)
		return planet
	default:
		panic(err)
	}
	return planet

}

func SearchInLocalFilms(id string) Film {
	fmt.Println("Hit Point : SearchInLocalFilms ", id)
	var film Film
	db, _ := sql.Open("sqlite3", "./cacheDatabase.db")

	stmt, err := db.Prepare("SELECT * FROM films WHERE id = ?")
	if err != nil {
		fmt.Println("Error 1 SearchInLocal : ", err.Error())
	}

	row := stmt.QueryRow(id)

	var species, vehicles, starships, planets, characters string

	switch err := row.Scan(&id,
		&film.Title,
		&film.EpisodeId,
		&film.OpeningCrawl,
		&film.Director,
		&film.Producer,
		&film.ReleaseDate,
		&characters,
		&planets,
		&starships,
		&vehicles,
		&species,
		&film.Created,
		&film.Edited,
		&film.Url); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:

		if len(characters) > 0 {
			film.Characters = strings.Split(characters, " ")
		} else {
			film.Characters = []string{}
		}

		if len(species) > 0 {
			film.Species = strings.Split(species, " ")
		} else {
			film.Species = []string{}
		}

		if len(vehicles) > 0 {
			film.Vehicles = strings.Split(vehicles, " ")
		} else {
			film.Vehicles = []string{}
		}
		if len(starships) > 0 {
			film.Starships = strings.Split(starships, " ")
		} else {
			film.Starships = []string{}
		}

		if len(planets) > 0 {
			film.Planets = strings.Split(planets, " ")
		} else {
			film.Planets = []string{}
		}

		fmt.Println("Result from ROW : ", film)
		// byteString := []byte(people)
		return film
	default:
		panic(err)
	}
	return film

}

func SearchInLocalSpecies(id string) Species {
	fmt.Println("Hit Point : SearchInLocalSpecies ", id)
	var species Species
	db, _ := sql.Open("sqlite3", "./cacheDatabase.db")

	stmt, err := db.Prepare("SELECT * FROM species WHERE id = ?")
	if err != nil {
		fmt.Println("Error 1 SearchInLocal : ", err.Error())
	}

	row := stmt.QueryRow(id)

	var films, peoples string

	switch err := row.Scan(&id,
		&species.Name,
		&species.Classification,
		&species.Designation,
		&species.AverageHeight,
		&species.SkinColors,
		&species.HairHolors,
		&species.EyeColors,
		&species.AverageLifespan,
		&species.Homeworld,
		&species.Language,
		&peoples,
		&films,
		&species.Created,
		&species.Edited,
		&species.Url); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:

		if len(films) > 0 {
			species.Films = strings.Split(films, " ")
		} else {
			species.Films = []string{}
		}

		if len(peoples) > 0 {
			species.People = strings.Split(peoples, " ")
		} else {
			species.People = []string{}
		}

		fmt.Println("Result from ROW : ", species)
		// byteString := []byte(people)
		return species
	default:
		panic(err)
	}
	return species

}

func SearchInLocalStarships(id string) Starship {
	fmt.Println("Hit Point : SearchInLocalStarships ", id)
	var starship Starship
	db, _ := sql.Open("sqlite3", "./cacheDatabase.db")

	stmt, err := db.Prepare("SELECT * FROM starships WHERE id = ?")
	if err != nil {
		fmt.Println("Error 1 SearchInLocal : ", err.Error())
	}

	row := stmt.QueryRow(id)

	var films, pilots string

	switch err := row.Scan(&id,
		&starship.Name,
		&starship.Model,
		&starship.Manufacturer,
		&starship.CostInCredits,
		&starship.Length,
		&starship.MaxAtmospheringSpeed,
		&starship.Crew,
		&starship.Passengers,
		&starship.CargoCapacity,
		&starship.Consumables,
		&starship.HyperdriveRating,
		&starship.MGLT,
		&starship.StarshipClass,
		&pilots,
		&films,
		&starship.Created,
		&starship.Edited,
		&starship.Url); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:

		if len(films) > 0 {
			starship.Films = strings.Split(films, " ")
		} else {
			starship.Films = []string{}
		}

		if len(pilots) > 0 {
			starship.Pilots = strings.Split(pilots, " ")
		} else {
			starship.Pilots = []string{}
		}

		fmt.Println("Result from ROW : ", starship)
		// byteString := []byte(people)
		return starship
	default:
		panic(err)
	}
	return starship

}

func SearchInLocalVehicles(id string) Vehicle {
	fmt.Println("Hit Point : SearchInLocalVehicles ", id)
	var vehicle Vehicle
	db, _ := sql.Open("sqlite3", "./cacheDatabase.db")

	stmt, err := db.Prepare("SELECT * FROM vehicles WHERE id = ?")
	if err != nil {
		fmt.Println("Error 1 SearchInLocal : ", err.Error())
	}

	row := stmt.QueryRow(id)

	var films, pilots string

	switch err := row.Scan(&id,
		&vehicle.Name,
		&vehicle.Model,
		&vehicle.Manufacturer,
		&vehicle.CostInCredits,
		&vehicle.Length,
		&vehicle.MaxAtmospheringSpeed,
		&vehicle.Crew,
		&vehicle.Passengers,
		&vehicle.CargoCapacity,
		&vehicle.Consumables,
		&vehicle.VehicleClass,
		&pilots,
		&films,
		&vehicle.Created,
		&vehicle.Edited,
		&vehicle.Url); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:

		if len(films) > 0 {
			vehicle.Films = strings.Split(films, " ")
		} else {
			vehicle.Films = []string{}
		}

		if len(pilots) > 0 {
			vehicle.Pilots = strings.Split(pilots, " ")
		} else {
			vehicle.Pilots = []string{}
		}

		fmt.Println("Result from ROW : ", vehicle)
		// byteString := []byte(people)
		return vehicle
	default:
		panic(err)
	}
	return vehicle
}

func InsertIntoPeoples(id string, data []byte) bool {

	fmt.Println("Hit Point : InsertIntoPeoples ")
	db, _ := sql.Open("sqlite3", "./cacheDatabase.db")

	insertPeopleSQL := `INSERT INTO people(id, name, height ,mass ,hair_color, skin_color, eye_color, birth_year , gender ,homeworld ,films,species,
	vehicles, starships, created, edited, url) VALUES (?, ?, ?,?,?, ?, ?,?,?, ?, ?,?,?, ?, ?,?,?)`
	statement, err := db.Prepare(insertPeopleSQL) // Prepare statement.

	var people People
	json.Unmarshal(data, &people)

	// This is good to avoid SQL injections
	if err != nil {
		fmt.Println("Error 1 : ", err.Error())
		return false
	}

	res, err := statement.Exec(id,
		people.Name,
		people.Height,
		people.Mass,
		people.HairColor,
		people.SkinColor,
		people.EyeColor,
		people.BirthYear,
		people.Gender,
		people.Homeworld,
		strings.Join(people.Films, " "),
		strings.Join(people.Species, " "),
		strings.Join(people.Vehicles, " "),
		strings.Join(people.Starships, " "),
		people.Created,
		people.Edited,
		people.Url,
	)
	if err != nil {
		fmt.Println("Error 2 : ", err.Error())
		return false
	}

	fmt.Println("Inserted a single record: ", res)

	return true
}

func InsertIntoPlanets(id string, data []byte) bool {

	fmt.Println("Hit Point : InsertIntoPlanets ")
	db, _ := sql.Open("sqlite3", "./cacheDatabase.db")

	insertSQL := `INSERT INTO planets(id, name, rotation_period ,orbital_period ,diameter, climate, gravity, terrain , surface_water ,population ,residents,films,
	created, edited, url) VALUES (?, ?, ?,?,?, ?, ?,?,?, ?, ?,?,?, ?, ?)`
	statement, err := db.Prepare(insertSQL) // Prepare statement.

	var planet Planet
	json.Unmarshal(data, &planet)

	// This is good to avoid SQL injections
	if err != nil {
		fmt.Println("Error 1 : ", err.Error())
		return false
	}

	res, err := statement.Exec(id,
		planet.Name,
		planet.RotationPeriod,
		planet.OrbitalPeriod,
		planet.Diameter,
		planet.Climate,
		planet.Gravity,
		planet.Terrain,
		planet.SurfaceWater,
		planet.Population,
		strings.Join(planet.Films, " "),
		strings.Join(planet.Residents, " "),
		planet.Created,
		planet.Edited,
		planet.Url,
	)
	if err != nil {
		fmt.Println("Error 2 : ", err.Error())
		return false
	}

	fmt.Println("Inserted a single record: ", res)

	return true
}

func InsertIntoFilms(id string, data []byte) bool {

	fmt.Println("Hit Point : InsertIntoFilms ")
	db, _ := sql.Open("sqlite3", "./cacheDatabase.db")

	insertSQL := `INSERT INTO films(id, title, episode_id ,opening_crawl ,director, producer, release_date, characters , planets ,starships ,vehicles,species,
	created, edited, url) VALUES (?, ?, ?,?,?, ?, ?,?,?, ?, ?,?,?, ?, ?)`
	statement, err := db.Prepare(insertSQL) // Prepare statement.

	var film Film
	json.Unmarshal(data, &film)

	// This is good to avoid SQL injections
	if err != nil {
		fmt.Println("Error 1 : ", err.Error())
		return false
	}

	res, err := statement.Exec(id,
		film.Title,
		film.EpisodeId,
		film.OpeningCrawl,
		film.Director,
		film.Producer,
		film.ReleaseDate,
		strings.Join(film.Characters, " "),
		strings.Join(film.Planets, " "),
		strings.Join(film.Starships, " "),
		strings.Join(film.Vehicles, " "),
		strings.Join(film.Species, " "),
		film.Created,
		film.Edited,
		film.Url,
	)
	if err != nil {
		fmt.Println("Error 2 : ", err.Error())
		return false
	}

	fmt.Println("Inserted a single record: ", res)

	return true
}

func InsertIntoSpecies(id string, data []byte) bool {

	fmt.Println("Hit Point : InsertIntoSpecies ")
	db, _ := sql.Open("sqlite3", "./cacheDatabase.db")

	insertSQL := `INSERT INTO species (id, name, classification ,designation ,average_height, skin_colors, hair_colors, eye_colors , average_lifespan ,homeworld ,language,people,
	films,created, edited, url) VALUES (?, ?, ?,?,?, ?, ?,?,?, ?, ?,?,?, ?, ?,?)`
	statement, err := db.Prepare(insertSQL) // Prepare statement.

	var species Species
	json.Unmarshal(data, &species)

	// This is good to avoid SQL injections
	if err != nil {
		fmt.Println("Error 1 : ", err.Error())
		return false
	}

	res, err := statement.Exec(id,
		species.Name,
		species.Classification,
		species.Designation,
		species.AverageHeight,
		species.SkinColors,
		species.HairHolors,
		species.EyeColors,
		species.AverageLifespan,
		species.Homeworld,
		species.Language,
		strings.Join(species.People, " "),
		strings.Join(species.Films, " "),
		species.Created,
		species.Edited,
		species.Url,
	)
	if err != nil {
		fmt.Println("Error 2 : ", err.Error())
		return false
	}

	fmt.Println("Inserted a single record: ", res)

	return true
}

func InsertIntoStarships(id string, data []byte) bool {

	fmt.Println("Hit Point : InsertIntoStarships ")
	db, _ := sql.Open("sqlite3", "./cacheDatabase.db")

	insertSQL := `INSERT INTO starships (id, name, model ,manufacturer ,cost_in_credits, length, max_atmosphering_speed, crew , passengers ,cargo_capacity ,consumables,hyperdrive_rating,
	MGLT,starship_class,pilots,films,created, edited, url) VALUES (?, ?, ?,?,?, ?, ?,?,?, ?, ?,?,?, ?, ?,?,?,?,?)`
	statement, err := db.Prepare(insertSQL) // Prepare statement.

	var starship Starship
	json.Unmarshal(data, &starship)

	// This is good to avoid SQL injections
	if err != nil {
		fmt.Println("Error 1 : ", err.Error())
		return false
	}

	res, err := statement.Exec(id,
		starship.Name,
		starship.Model,
		starship.Manufacturer,
		starship.CostInCredits,
		starship.Length,
		starship.MaxAtmospheringSpeed,
		starship.Crew,
		starship.Passengers,
		starship.CargoCapacity,
		starship.Consumables,
		starship.HyperdriveRating,
		starship.MGLT,
		starship.StarshipClass,
		strings.Join(starship.Pilots, " "),
		strings.Join(starship.Films, " "),
		starship.Created,
		starship.Edited,
		starship.Url,
	)
	if err != nil {
		fmt.Println("Error 2 : ", err.Error())
		return false
	}

	fmt.Println("Inserted a single record: ", res)

	return true
}

func InsertIntoVehicles(id string, data []byte) bool {

	fmt.Println("Hit Point : InsertIntoVehicles ")
	db, _ := sql.Open("sqlite3", "./cacheDatabase.db")

	insertSQL := `INSERT INTO vehicles (id, name, model ,manufacturer ,cost_in_credits, length, max_atmosphering_speed, crew , passengers ,cargo_capacity ,consumables,vehicle_class,
	pilots,films,created, edited, url) VALUES (?, ?, ?,?,?, ?, ?,?,?, ?, ?,?,?, ?, ?,?,?)`
	statement, err := db.Prepare(insertSQL) // Prepare statement.

	var vehicle Vehicle
	json.Unmarshal(data, &vehicle)

	// This is good to avoid SQL injections
	if err != nil {
		fmt.Println("Error 1 : ", err.Error())
		return false
	}

	res, err := statement.Exec(id,
		vehicle.Name,
		vehicle.Model,
		vehicle.Manufacturer,
		vehicle.CostInCredits,
		vehicle.Length,
		vehicle.MaxAtmospheringSpeed,
		vehicle.Crew,
		vehicle.Passengers,
		vehicle.CargoCapacity,
		vehicle.Consumables,
		vehicle.VehicleClass,
		strings.Join(vehicle.Pilots, " "),
		strings.Join(vehicle.Films, " "),
		vehicle.Created,
		vehicle.Edited,
		vehicle.Url,
	)
	if err != nil {
		fmt.Println("Error 2 : ", err.Error())
		return false
	}

	fmt.Println("Inserted a single record: ", res)

	return true
}
