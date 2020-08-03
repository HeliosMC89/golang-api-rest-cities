package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the City Database.")
}

func CityList(w http.ResponseWriter, r *http.Request) {
	// Query The database
	cities, err := dbCityList()
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, JSONResponse{
			Error: true,
			Msg:   "Internal Error",
			Code:  http.StatusInternalServerError,
			Data:  nil,
		})
	}

	// Format the response.
	WriteJSON(w, http.StatusOK, JSONResponse{
		Error: false,
		Msg:   "List of City",
		Code:  http.StatusOK,
		Data:  cities,
	})
}

func CityDisplay(w http.ResponseWriter, r *http.Request) {
	// Get URL parameter with the city ID search for.
	vars := mux.Vars(r)
	cityID, err := strconv.Atoi(vars["id"])
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, JSONResponse{
			Error: true,
			Msg:   "Bad request",
			Code:  http.StatusBadRequest,
			Data:  nil,
		})
		return
	}

	// Query the database.
	city, err := dbCityDisplay(cityID)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, JSONResponse{
			Error: true,
			Msg:   "Internal server error",
			Code:  http.StatusInternalServerError,
			Data:  nil,
		})
		return
	}

	// Send the response.
	WriteJSON(w, http.StatusOK, JSONResponse{
		Error: false,
		Msg:   "Display City",
		Code:  http.StatusOK,
		Data:  city,
	})
}

func CityAdd(w http.ResponseWriter, r *http.Request) {
	var city City

	// Read the body of the request.
	err := ReadJSON(r, &city)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, JSONResponse{
			Error: true,
			Msg:   "Body malformed",
			Code:  http.StatusUnprocessableEntity,
			Data:  nil,
		})
		return
	}
	defer r.Body.Close()

	// Write to the database
	addResult, err := dbCityAdd(city)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, JSONResponse{
			Error: true,
			Msg:   "internal error",
			Code:  http.StatusInternalServerError,
			Data:  nil,
		})
		return
	}

	// Format response
	WriteJSON(w, http.StatusCreated, JSONResponse{
		Error: false,
		Msg:   "City added",
		Code:  http.StatusCreated,
		Data:  addResult,
	})
}

func CityUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cityID, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, JSONResponse{
			Error: true,
			Msg:   "Bad request",
			Code:  http.StatusBadRequest,
			Data:  nil,
		})
		return
	}

	var city City
	// Read the body of the request.
	err = ReadJSON(r, &city)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, JSONResponse{
			Error: true,
			Msg:   "Body malformed",
			Code:  http.StatusUnprocessableEntity,
			Data:  nil,
		})
		return
	}
	defer r.Body.Close()

	// Query the database.
	updateResult, err := dbCityUpdate(cityID, city)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, JSONResponse{
			Error: true,
			Msg:   "internal error",
			Code:  http.StatusInternalServerError,
			Data:  nil,
		})
		return
	}

	// Send the response
	WriteJSON(w, http.StatusOK, JSONResponse{
		Error: false,
		Msg:   "Ok city update",
		Code:  http.StatusOK,
		Data:  updateResult,
	})

}

func CityDelete(w http.ResponseWriter, r *http.Request) {
	// Get URL parameters with the city ID to delete.
	vars := mux.Vars(r)
	cityID, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, JSONResponse{
			Error: true,
			Msg:   "Bad request",
			Code:  http.StatusBadRequest,
			Data:  nil,
		})
		return
	}

	// Query the database.
	deleteResult, err := dbCityDelete(cityID)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, JSONResponse{
			Error: true,
			Msg:   "internal error",
			Code:  http.StatusInternalServerError,
			Data:  nil,
		})
		return
	}

	// Send the response
	WriteJSON(w, http.StatusOK, JSONResponse{
		Error: false,
		Msg:   "Ok city deleted",
		Code:  http.StatusOK,
		Data:  deleteResult,
	})

}
