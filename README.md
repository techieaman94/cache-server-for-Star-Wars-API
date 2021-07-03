# Cache server for Star Wars API


## Description

A caching engine/server built using Golang and Sqlite for [The Star Wars API](https://swapi.dev/)

 * If a user calls the API to get details of any Planet, Starship, Vehicle, People, Film or Species then, first required details are checked in local storage. 

 * If record found in local storage then it is send back to the user as json response. 

 * If required data is not found in local database then API call is done to https://swapi.dev/api/ , obtained response is stored in local database and send back to the user.

 * After every interval of 15 minutes, all records from local database tables are deleted.
 

### Project structure
```
.
├── lib
|   ├── models.go
|   ├── apiSwapi.go
|   └── dbHandler.go
├── cacheEngine.go
└── README.md
```

### Installation and running the server

* Install [Golang](https://golang.org/).
* Install sqlite using command ```go get github.com/mattn/go-sqlite3 ```
* Download or Clone this repo in local system.
* Go inside the project folder where 'cacheEngine.go' file is present.
* Open in terminal.
* Run the cache engine by command `go run cacheEngine.go`.

  Use 'GET' on endpoint http://localhost:10006/dataType or http://localhost:10006/dataType/id
  
  **_dataType_** value can be 'people', 'planets', 'vehicles', 'films', 'species' or 'starships'

  Example - 
  ```
  Method : GET
  Url : http://localhost:10006/films/3
  ```

