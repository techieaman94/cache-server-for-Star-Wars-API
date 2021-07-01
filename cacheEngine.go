package main

import (
    "encoding/json"
    // "fmt"
    "github.com/gorilla/mux"
    "log"
    "net/http"
    "time"

    L "./lib"
)

// func homePage(w http.ResponseWriter, r *http.Request) {
//     fmt.Fprintf(w, "Welcome to the HomePage! Cache Server")
//     fmt.Println("Endpoint Hit: homePage")
// }

func checkCache(w http.ResponseWriter, r *http.Request) {
    // fmt.Println("Endpoint Hit: checkCache", r, "\n")
    w.Header().Set("Content-Type", "application/json")
    vars := mux.Vars(r)
    // fmt.Println("___ VARS ___",vars)
    var dataType, key string

    if vars["dataType"] != "" {
        dataType = vars["dataType"]
        // fmt.Println("dataType: ", dataType)
    }

    if vars["id"] != "" {
        key = vars["id"]
        // fmt.Println("Key: ", key)
    }

    if dataType != "" && key != "" {
        switch dataType {
        case "people":
            /* code */
            dataToReturn := L.SearchInLocalPeople(key)

            if dataToReturn.Name != "" {
                // fmt.Println("____ DATA FOUND ___ : ", dataToReturn)
                json.NewEncoder(w).Encode(dataToReturn)
            } else {
                // GET it from SWAPI API , store localy and return to end user
                temp := L.GetOneRecord(dataType, key)
                var people L.People
                json.Unmarshal(temp, &people)

                // fmt.Println("data in catchEngine : ", people, "\n")

                if people.Name != "" {
                    json.NewEncoder(w).Encode(people)
                } else {
                    http.NotFound(w, r)
                }
            }

        case "planets":
            /* code */
            dataToReturn := L.SearchInLocalPlanets(key)

            if dataToReturn.Name != "" {
                // fmt.Println("____ DATA FOUND ___ : ", dataToReturn)
                json.NewEncoder(w).Encode(dataToReturn)
            } else {
                // GET it from SWAPI API , store localy and return to end user
                temp := L.GetOneRecord(dataType, key)
                var planet L.Planet
                json.Unmarshal(temp, &planet)

                // fmt.Println("data in catchEngine : ", planet, "\n")

                if planet.Name != "" {
                    json.NewEncoder(w).Encode(planet)
                } else {
                    http.NotFound(w, r)
                }
            }

        case "vehicles":
            /* code */
            dataToReturn := L.SearchInLocalVehicles(key)

            if dataToReturn.Name != "" {
                // fmt.Println("____ DATA FOUND ___ : ", dataToReturn)
                json.NewEncoder(w).Encode(dataToReturn)
            } else {
                // GET it from SWAPI API , store localy and return to end user
                temp := L.GetOneRecord(dataType, key)
                var vehicle L.Vehicle
                json.Unmarshal(temp, &vehicle)

                // fmt.Println("data in catchEngine : ", vehicle, "\n")
                if vehicle.Name != "" {
                    json.NewEncoder(w).Encode(vehicle)
                } else {
                    http.NotFound(w, r)
                }

            }
        case "films":
            /* code */
            dataToReturn := L.SearchInLocalFilms(key)

            if dataToReturn.Title != "" {
                // fmt.Println("____ DATA FOUND ___ : ", dataToReturn)
                json.NewEncoder(w).Encode(dataToReturn)
            } else {
                // GET it from SWAPI API , store localy and return to end user
                temp := L.GetOneRecord(dataType, key)
                var film L.Film
                json.Unmarshal(temp, &film)

                // fmt.Println("data in catchEngine : ", film, "\n")

                if film.Title != "" {
                    json.NewEncoder(w).Encode(film)
                } else {
                    http.NotFound(w, r)
                }
            }
        case "species":
            /* code */
            dataToReturn := L.SearchInLocalSpecies(key)

            if dataToReturn.Name != "" {
                // fmt.Println("____ DATA FOUND ___ : ", dataToReturn)
                json.NewEncoder(w).Encode(dataToReturn)
            } else {
                // GET it from SWAPI API , store localy and return to end user
                temp := L.GetOneRecord(dataType, key)
                var species L.Species
                json.Unmarshal(temp, &species)

                // fmt.Println("data in catchEngine : ", species, "\n")

                if species.Name != "" {
                    json.NewEncoder(w).Encode(species)
                } else {
                    http.NotFound(w, r)
                }
            }
        case "starships":
            /* code */
            dataToReturn := L.SearchInLocalStarships(key)

            if dataToReturn.Name != "" {
                // fmt.Println("____ DATA FOUND ___ : ", dataToReturn)
                json.NewEncoder(w).Encode(dataToReturn)
            } else {
                // GET it from SWAPI API , store localy and return to end user
                temp := L.GetOneRecord(dataType, key)
                var starships L.Starship
                json.Unmarshal(temp, &starships)

                // fmt.Println("data in catchEngine : ", starships, "\n")

                if starships.Name != "" {
                    json.NewEncoder(w).Encode(starships)
                } else {
                    http.NotFound(w, r)
                }

            }

        default:
            // code
            http.NotFound(w, r)
        }

    } else if dataType != "" {

        temp := L.GetAllRecords(dataType)

        switch dataType {
        case "people":
            /* code */
            var peoplesAll L.PeopleAll

            json.Unmarshal(temp, &peoplesAll)

            // fmt.Println("data in catchEngine : ", peoplesAll, "\n")

            if peoplesAll.Count != 0 {
                json.NewEncoder(w).Encode(peoplesAll)
            } else {
                http.NotFound(w, r)
            }

        case "films":
            var filmAll L.FilmAll

            json.Unmarshal(temp, &filmAll)

            // fmt.Println("data in catchEngine : ", filmAll, "\n")

            if filmAll.Count != 0 {
                json.NewEncoder(w).Encode(filmAll)
            } else {
                http.NotFound(w, r)
            }

        case "species":
            var speciesAll L.SpeciesAll

            json.Unmarshal(temp, &speciesAll)

            // fmt.Println("data in catchEngine : ", speciesAll, "\n")

            if speciesAll.Count != 0 {
                json.NewEncoder(w).Encode(speciesAll)
            } else {
                http.NotFound(w, r)
            }
        case "starships":
            var starshipAll L.StarshipAll

            json.Unmarshal(temp, &starshipAll)

            // fmt.Println("data in catchEngine : ", starshipAll, "\n")

            if starshipAll.Count != 0 {
                json.NewEncoder(w).Encode(starshipAll)
            } else {
                http.NotFound(w, r)
            }

        case "vehicles":
            var vehicleAll L.VehicleAll

            json.Unmarshal(temp, &vehicleAll)

            // fmt.Println("data in catchEngine : ", vehicleAll, "\n")

            if vehicleAll.Count != 0 {
                json.NewEncoder(w).Encode(vehicleAll)
            } else {
                http.NotFound(w, r)
            }

        case "planets":
            var planetAll L.PlanetAll

            json.Unmarshal(temp, &planetAll)

            // fmt.Println("data in catchEngine : ", planetAll, "\n")

            if planetAll.Count != 0 {
                json.NewEncoder(w).Encode(planetAll)
            } else {
                http.NotFound(w, r)
            }

        default:
            /* code */
            http.NotFound(w, r)
        }

    }

}

func clearCacheAfterSomeInterval() {
    for {
        <-time.After(15 * time.Minute)
        // <-time.After(20 * time.Second)
        go L.ClearDB()
    }
}

func handleRequests() {

    myRouter := mux.NewRouter().StrictSlash(true)
    // myRouter.HandleFunc("/", homePage)

    myRouter.HandleFunc("/{dataType}", checkCache).Methods("GET")
    myRouter.HandleFunc("/{dataType}/{id}", checkCache).Methods("GET")

    log.Fatal(http.ListenAndServe(":10006", myRouter))
}

func main() {
    L.CreateTablesIfNotExists()
    go clearCacheAfterSomeInterval()

    handleRequests()
}
