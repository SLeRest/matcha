package main


import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"gopkg.in/go-playground/validator.v9"
)

type user struct {
	ID					string `json:"ID"`
	Email				string `json:"Email" validate:"required,email"`
	Password			string `json:"Password" validate:"required,min=8"`
	FirstName			string `json:"FirstName" validate:"required,min=2,max=100"`
	LastName			string `json:"LastName" validate:"required,min=2,max=100"`
	City				string `json:"City"`
	Country				string `json:"Country"`
	Age					int `json:"Age" validate:"required,min=18,max=130"`
	Gender				string `json:"Gender"`
	SexualOrientation	string `json:"SexualOrientation"`
	Description			string `json:"Description"`
	Tags				[]string `json:"Tags"` // TODO array ? json ?
	// TODO image dans BDD ? filesysteme ?
}

// TODO BDD
type allUsers []user

var users = allUsers {
	{
		ID:					"1",
		Email:				"seb@gmail.fr",
		Password:			"strong_password",
		FirstName:			"seb",
		LastName:			"lerest",
		City:				"Rennes",
		Country:			"France",
		Age:				26,
		Gender:				"Male",
		SexualOrientation:	"Female",
		Description:		"coucou c moi",
		Tags:				[]string{"nerdz", "coucou"},
	},
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser user
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Enter data for create new User.")
	}
	json.Unmarshal(reqBody, &newUser)
	validate := validator.New()
	err = validate.Struct(newUser)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// TODO requete BDD
	users = append(users, newUser)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newUser)
}

func getOneUser(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["id"]

	// TODO requete BDD
	for _, singleUser := range users {
		if singleUser.ID == userID {
			json.NewEncoder(w).Encode(singleUser)
		}
	}
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	// TODO requete BDD
	json.NewEncoder(w).Encode(users)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["id"]
	var updatedUser user

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}
	json.Unmarshal(reqBody, &updatedUser)

	for i, singleUser := range users {

		if singleUser.ID == userID {
			validate := validator.New()
			err = validate.Struct(updateUser)
			if err != nil {
				fmt.Fprintf(w, err.Error())
				w.WriteHeader(http.StatusBadRequest)
				continue
			}
			singleUser.Email = updatedUser.Email
			singleUser.Password = updatedUser.FirstName
			singleUser.LastName = updatedUser.LastName
			singleUser.City = updatedUser.City
			singleUser.Country = updatedUser.Country
			singleUser.Age = updatedUser.Age
			singleUser.Gender = updatedUser.Gender
			singleUser.SexualOrientation = updatedUser.SexualOrientation
			singleUser.Description = updatedUser.Description
			singleUser.Tags = updatedUser.Tags
			users = append(users[:i], singleUser)
			json.NewEncoder(w).Encode(singleUser)
		}
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["id"]

	for i, singleUser := range users {
		if singleUser.ID == userID {
			// TODO bizarre ca pour enelever un elem d'une liste
			users = append(users[:i], users[i+1:]...)
			fmt.Fprintf(w, "The user with ID %v has been deleted successfully", userID)
		}
	}
}
