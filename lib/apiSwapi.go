package cacheServer

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var baseUrl = "https://swapi.dev/api/"

func GetOneRecord(dataType string, id string) []byte {
	fmt.Println("Endpoint Hit: GetOneRecords", "\n")
	Url := baseUrl + dataType + "/" + id

	req, _ := http.NewRequest("GET", Url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	// fmt.Println("body GetOneRecords :", string(body), "\n")

	switch dataType {
	case "people":

		res := InsertIntoPeoples(id, body)

		if res == true {
			fmt.Println("_____ INSERTED RECORD IN people _____")
		} else {
			fmt.Println("_____ NOT INSERTED RECORD IN people _____")
		}

	case "planets":

		res := InsertIntoPlanets(id, body)

		if res == true {
			fmt.Println("_____ INSERTED RECORD IN planets _____")
		} else {
			fmt.Println("_____ NOT INSERTED RECORD IN planets _____")
		}
	case "films":

		res := InsertIntoFilms(id, body)

		if res == true {
			fmt.Println("_____ INSERTED RECORD IN films _____")
		} else {
			fmt.Println("_____ NOT INSERTED RECORD IN films _____")
		}
	case "species":

		res := InsertIntoSpecies(id, body)

		if res == true {
			fmt.Println("_____ INSERTED RECORD IN species _____")
		} else {
			fmt.Println("_____ NOT INSERTED RECORD IN species _____")
		}
	case "starships":

		res := InsertIntoStarships(id, body)

		if res == true {
			fmt.Println("_____ INSERTED RECORD IN starships _____")
		} else {
			fmt.Println("_____ NOT INSERTED RECORD IN starships _____")
		}
	case "vehicles":

		res := InsertIntoVehicles(id, body)

		if res == true {
			fmt.Println("_____ INSERTED RECORD IN vehicles _____")
		} else {
			fmt.Println("_____ NOT INSERTED RECORD IN vehicles _____")
		}

	default:
	}

	return body
}

func GetAllRecords(dataType string) []byte {

	// fmt.Println("Endpoint Hit: GetAllRecords", "\n")
	Url := baseUrl + dataType

	req, _ := http.NewRequest("GET", Url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// fmt.Println("body GetAllRecords :", string(body), "\n")

	return body

}
