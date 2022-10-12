package controller

import (
	"encoding/json"
	"epiphanius_bot/pkg/models"
	"epiphanius_bot/pkg/requests"
	"epiphanius_bot/pkg/storage"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"gorm.io/datatypes"
)


func CreateUser(userId int64, userName string) {

	users := models.Users {
      UserName: userName,
	  UserID:   userId,
	}
	if err := storage.CreateUser(users); err != nil {
		fmt.Println(err)
	}
}


func CreateHoliday(w http.ResponseWriter, r *http.Request) {
	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var request requests.Holiday
	if err := json.Unmarshal(body, &request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Error: %v", err)))
		return
	}

	date,err := time.Parse("02.01.2006", request.Date)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Error: %v", err)))
		return
	}
    // fmt.Println("=id=id=id=")
    // fmt.Println(request.Id)
    // fmt.Println("=id=id=id=")
	holiday := models.Holiday{
		Name: request.Name,
		Description: request.Description,
		Date: datatypes.Date(date),
	}

	// Append to the Books table

	if err := storage.CreateHoliday(holiday); err != nil {
		fmt.Println(err)
	}

	// Send a 201 created response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}

func DeleteHoliday(w http.ResponseWriter, r *http.Request) {
	// Read the dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Delete that book
	storage.DeleteHoliday(id)

	log.Info("delete record")

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")
}

func GetHoliday(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	holiday, err := storage.GetHoliday(id)
	if err != nil {
		// TODO: handle error
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(holiday)
}

func GetHolidays(w http.ResponseWriter, r *http.Request) {
	holidays, err := storage.GetHolidays()
	if err != nil {
		// TODO: handle error
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(holidays)
}

func UpdateHoliday(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	// Read request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	//log.Info("this body has:", body)
	if err != nil {
		log.Fatalln(err)
	}

	var request requests.Holiday
	err = json.Unmarshal(body, &request)
	fmt.Println("=bodyUpdate=bodyUpdate")
	fmt.Println(request)
	fmt.Println("=bodyUpdate=bodyUpdate")
	//log.Info("json unmarshal update", update, id)
	if err != nil {
		// TODO: handle error
		return
	}

	
	date,err := time.Parse("02.01.2006", request.Date)
	fmt.Println("RomaRomaRoma")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Error: %v", err)))
		return
	}
	

	update := models.Holiday{
		Name: request.Name,
		Description: request.Description,
		Date: datatypes.Date(date),
	}

	err = storage.UpdateHoliday(id, update)
	if err != nil {
		log.Info("update error")
		return
	}

	log.Info("update record")
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Updated")
}
