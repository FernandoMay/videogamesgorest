package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func setupVideogamesRoutes(router *mux.Router) {
	enableCORS(router)

	router.HandleFunc("/videogames", func(w http.ResponseWriter, r *http.Request) {
		videogames, err := getVideogames()
		if err != nil {
			respondWithSuccess(videogames, w)
		} else {
			respondWithError(err, w)
		}
	}).Methods(http.MethodGet)

	router.HandleFunc("/videogame/{id}", func(w http.ResponseWriter, r *http.Request) {
		idString := mux.Vars(r)["id"]
		id, err := stringToInt64(idString)
		if err != nil {
			respondWithError(err, w)
			return
		}
		videogame, err := getVideogameById(id)
		if err != nil {
			respondWithSuccess(videogame, w)
		} else {
			respondWithError(err, w)
		}
	}).Methods(http.MethodGet)

	router.HandleFunc("/videogame/{id}", func(w http.ResponseWriter, r *http.Request) {
		idString := mux.Vars(r)["id"]
		id, err := stringToInt64(idString)
		if err != nil {
			respondWithError(err, w)
			return
		}
		err = deleteVideogame(id)
		if err != nil {
			respondWithSuccess(true, w)
		} else {
			respondWithError(err, w)
		}
	}).Methods(http.MethodDelete)

	router.HandleFunc("/videogame", func(w http.ResponseWriter, r *http.Request) {
		var videogame VideoGame
		err := json.NewDecoder(r.Body).Decode(&videogame)
		if err != nil {
			respondWithError(err, w)

		} else {
			err := createVideogame(videogame)
			if err != nil {
				respondWithSuccess(videogame, w)
			} else {
				respondWithError(err, w)
			}
		}
	}).Methods(http.MethodPost)

	router.HandleFunc("/videogame", func(w http.ResponseWriter, r *http.Request) {
		var videogame VideoGame
		err := json.NewDecoder(r.Body).Decode(&videogame)
		if err != nil {
			respondWithError(err, w)

		} else {
			err := createVideogame(videogame)
			if err != nil {
				respondWithSuccess(videogame, w)
			} else {
				respondWithError(err, w)
			}
		}
	}).Methods(http.MethodPut)

}

func enableCORS(router *mux.Router) {
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", AllowedCorsDomain)
	}).Methods(http.MethodOptions)
	router.Use(middlewareCors)
}

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", AllowedCorsDomain)
		w.Header().Set("Access-Control-Allow-Credntials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		next.ServeHTTP(w, r)
	})
}

func respondWithError(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(err.Error())
}

func respondWithSuccess(data interface{}, w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
