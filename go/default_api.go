/*
 * star world API
 *
 * the api to query the information about *Star War* you can check all the Star Wars data you've ever wanted Planets Spaceships Vehicles People Films and Species From all SEVEN Star Wars films
 *
 * API version: 0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	// db "github.com/S-Vanguard/StarWar_Server/db"
	db "github.com/S-Vanguard/StarWar_Server/mysqlDB"
)

func FilmsGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	values := r.URL.Query()
	var page int

	// get data of this page
	var result [10]string
	if _, ok := values["page"]; ok {
		// has parameter
		if temp, err := strconv.Atoi(values["page"][0]); err == nil {
			page = temp
			result = db.GetFilmsByPage(page)
		} else {
			log.Fatal(err)
			return
		}
	} else {
		// no parameter, default to 1
		page = 1
		result = db.GetFilmsByPage(1)
	}

	// make json of FilmsList
	filmsList := &FilmsList{}
	filmsList.Count = 7

	// validation
	if !(page >= 1 && page <= int(filmsList.Count)/10+1) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not found"))
		return
	}

	// set next
	if page == int(filmsList.Count)/10+1 {
		filmsList.Next = "null"
	} else {
		filmsList.Next = "http://localhost:8080/?page=" + strconv.Itoa(page+1)
	}

	// set previous
	if page == 1 {
		filmsList.Previous = "null"
	} else {
		filmsList.Previous = "http://localhost:8080/?page=" + strconv.Itoa(page-1)
	}

	// set results
	for i := 0; i < 10; i++ {
		if result[i] == "" {
			break
		}
		film := &Films{}
		err := json.Unmarshal([]byte(result[i]), &film)
		if err == nil {
			filmsList.Results = append(filmsList.Results, *film)
		} else {
			log.Fatal(err)
			return
		}
	}

	// make json to []byte, and response
	b, err := json.Marshal(&filmsList)
	if err == nil {
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	} else {
		log.Fatal(err)
	}
}

func FilmsIdGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	id := strings.Trim(r.URL.Path, "/films/")
	json := db.GetFilmByID(id)
	if json == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not found"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(json))
	}
}

func PeopleGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	values := r.URL.Query()
	var page int

	// get data of this page
	var result [10]string
	if _, ok := values["page"]; ok {
		// has parameter
		if temp, err := strconv.Atoi(values["page"][0]); err == nil {
			page = temp
			result = db.GetPeopleByPage(page)
		} else {
			log.Fatal(err)
			return
		}
	} else {
		// no parameter, default to 1
		page = 1
		result = db.GetPeopleByPage(1)
	}

	// make json of PeopleList
	peopleList := &PeopleList{}
	peopleList.Count = 87

	// validation
	if !(page >= 1 && page <= int(peopleList.Count)/10+1) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not found"))
		return
	}

	// set next
	if page == int(peopleList.Count)/10+1 {
		peopleList.Next = "null"
	} else {
		peopleList.Next = "http://localhost:8080/?page=" + strconv.Itoa(page+1)
	}

	// set previous
	if page == 1 {
		peopleList.Previous = "null"
	} else {
		peopleList.Previous = "http://localhost:8080/?page=" + strconv.Itoa(page-1)
	}

	// set results
	for i := 0; i < 10; i++ {
		if result[i] == "" {
			break
		}
		person := &People{}
		err := json.Unmarshal([]byte(result[i]), &person)
		if err == nil {
			peopleList.Results = append(peopleList.Results, *person)
		} else {
			log.Fatal(err)
			return
		}
	}

	// make json to []byte, and response
	b, err := json.Marshal(&peopleList)
	if err == nil {
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	} else {
		log.Fatal(err)
	}
}

func PeopleIdGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	id := strings.Trim(r.URL.Path, "/people/")
	json := db.GetPeopleByID(id)
	if json == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not found"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(json))
	}
}

func PlanetsGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	values := r.URL.Query()
	var page int

	// get data of this page
	var result [10]string
	if _, ok := values["page"]; ok {
		// has parameter
		if temp, err := strconv.Atoi(values["page"][0]); err == nil {
			page = temp
			result = db.GetPlanetsByPage(page)
		} else {
			log.Fatal(err)
			return
		}
	} else {
		// no parameter, default to 1
		page = 1
		result = db.GetPlanetsByPage(1)
	}

	// make json of PlanetsList
	planetsList := &PlanetsList{}
	planetsList.Count = 61

	// validation
	if !(page >= 1 && page <= int(planetsList.Count)/10+1) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not found"))
		return
	}

	// set next
	if page == int(planetsList.Count)/10+1 {
		planetsList.Next = "null"
	} else {
		planetsList.Next = "http://localhost:8080/?page=" + strconv.Itoa(page+1)
	}

	// set previous
	if page == 1 {
		planetsList.Previous = "null"
	} else {
		planetsList.Previous = "http://localhost:8080/?page=" + strconv.Itoa(page-1)
	}

	// set results
	for i := 0; i < 10; i++ {
		if result[i] == "" {
			break
		}
		planet := &Planets{}
		err := json.Unmarshal([]byte(result[i]), &planet)
		if err == nil {
			planetsList.Results = append(planetsList.Results, *planet)
		} else {
			log.Fatal(err)
			return
		}
	}

	// make json to []byte, and response
	b, err := json.Marshal(&planetsList)
	if err == nil {
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	} else {
		log.Fatal(err)
	}
}

func PlanetsIdGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	id := strings.Trim(r.URL.Path, "/planets/")
	json := db.GetPlanetByID(id)
	if json == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not found"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(json))
	}
}

func SpeciesGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	values := r.URL.Query()
	var page int

	// get data of this page
	var result [10]string
	if _, ok := values["page"]; ok {
		// has parameter
		if temp, err := strconv.Atoi(values["page"][0]); err == nil {
			page = temp
			result = db.GetSpeciesByPage(page)
		} else {
			log.Fatal(err)
			return
		}
	} else {
		// no parameter, default to 1
		page = 1
		result = db.GetSpeciesByPage(1)
	}

	// make json of SpeciesList
	speciesList := &SpeciesList{}
	speciesList.Count = 37

	// validation
	if !(page >= 1 && page <= int(speciesList.Count)/10+1) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not found"))
		return
	}

	// set next
	if page == int(speciesList.Count)/10+1 {
		speciesList.Next = "null"
	} else {
		speciesList.Next = "http://localhost:8080/?page=" + strconv.Itoa(page+1)
	}

	// set previous
	if page == 1 {
		speciesList.Previous = "null"
	} else {
		speciesList.Previous = "http://localhost:8080/?page=" + strconv.Itoa(page-1)
	}

	// set results
	for i := 0; i < 10; i++ {
		if result[i] == "" {
			break
		}
		specie := &Species{}
		err := json.Unmarshal([]byte(result[i]), &specie)
		if err == nil {
			speciesList.Results = append(speciesList.Results, *specie)
		} else {
			log.Fatal(err)
			return
		}
	}

	// make json to []byte, and response
	b, err := json.Marshal(&speciesList)
	if err == nil {
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	} else {
		log.Fatal(err)
	}
}

func SpeciesIdGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	id := strings.Trim(r.URL.Path, "/species/")
	json := db.GetSpeciesByID(id)
	if json == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not found"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(json))
	}
}

func StarshipsGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	values := r.URL.Query()
	var page int

	// get data of this page
	var result [10]string
	if _, ok := values["page"]; ok {
		// has parameter
		if temp, err := strconv.Atoi(values["page"][0]); err == nil {
			page = temp
			result = db.GetStarshipsByPage(page)
		} else {
			log.Fatal(err)
			return
		}
	} else {
		// no parameter, default to 1
		page = 1
		result = db.GetStarshipsByPage(1)
	}

	// make json of StarshipsList
	starShipsList := &StarshipsList{}
	starShipsList.Count = 37

	// validation
	if !(page >= 1 && page <= int(starShipsList.Count)/10+1) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not found"))
		return
	}

	// set next
	if page == int(starShipsList.Count)/10+1 {
		starShipsList.Next = "null"
	} else {
		starShipsList.Next = "http://localhost:8080/?page=" + strconv.Itoa(page+1)
	}

	// set previous
	if page == 1 {
		starShipsList.Previous = "null"
	} else {
		starShipsList.Previous = "http://localhost:8080/?page=" + strconv.Itoa(page-1)
	}

	// set results
	for i := 0; i < 10; i++ {
		if result[i] == "" {
			break
		}
		starship := &Starships{}
		err := json.Unmarshal([]byte(result[i]), &starship)
		if err == nil {
			starShipsList.Results = append(starShipsList.Results, *starship)
		} else {
			log.Fatal(err)
			return
		}
	}

	// make json to []byte, and response
	b, err := json.Marshal(&starShipsList)
	if err == nil {
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	} else {
		log.Fatal(err)
	}
}

func StarshipsIdGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	id := strings.Trim(r.URL.Path, "/starships/")
	json := db.GetStarshipByID(id)
	if json == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not found"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(json))
	}
}

func VehiclesGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	values := r.URL.Query()
	var page int

	// get data of this page
	var result [10]string
	if _, ok := values["page"]; ok {
		// has parameter
		if temp, err := strconv.Atoi(values["page"][0]); err == nil {
			page = temp
			result = db.GetVehiclesByPage(page)
		} else {
			log.Fatal(err)
			return
		}
	} else {
		// no parameter, default to 1
		page = 1
		result = db.GetVehiclesByPage(1)
	}

	// make json of VehiclesList
	vehiclesList := &VehiclesList{}
	vehiclesList.Count = 39

	// validation
	if !(page >= 1 && page <= int(vehiclesList.Count)/10+1) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not found"))
		return
	}

	// set next
	if page == int(vehiclesList.Count)/10+1 {
		vehiclesList.Next = "null"
	} else {
		vehiclesList.Next = "http://localhost:8080/?page=" + strconv.Itoa(page+1)
	}

	// set previous
	if page == 1 {
		vehiclesList.Previous = "null"
	} else {
		vehiclesList.Previous = "http://localhost:8080/?page=" + strconv.Itoa(page-1)
	}

	// set results
	for i := 0; i < 10; i++ {
		if result[i] == "" {
			break
		}
		vehicle := &Vehicles{}
		err := json.Unmarshal([]byte(result[i]), &vehicle)
		if err == nil {
			vehiclesList.Results = append(vehiclesList.Results, *vehicle)
		} else {
			log.Fatal(err)
			return
		}
	}

	// make json to []byte, and response
	b, err := json.Marshal(&vehiclesList)
	if err == nil {
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	} else {
		log.Fatal(err)
	}
}

func VehiclesIdGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	id := strings.Trim(r.URL.Path, "/vehicles/")
	json := db.GetVehicleByID(id)
	if json == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not found"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(json))
	}
}
