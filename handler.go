package handler

import (
	"activity-tracker/database"
	"activity-tracker/model"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"time"
)

type ActivityHandler struct {
}

func (a ActivityHandler) AddActivity(w http.ResponseWriter, r *http.Request) {
	var ac model.Activity
	err := json.NewDecoder(r.Body).Decode(&ac)
	db, err := database.NewDataBaseConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print(err.Error())
	}
	defer db.Db.Close()
	var actime time.Time
	if ac.Time == "" {
		actime = time.Now()
	}
	db.AddActivities(ac.Name, ac.Type, actime)
	w.Write(append([]byte("SuccessFully Added Activity!!")))
}

func (a ActivityHandler) GetActivity(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["ID"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}
	key := keys[0]
	fmt.Printf(key)
	db, err := database.NewDataBaseConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print(err.Error())
	}
	defer db.Db.Close()
	actvity := db.GetActivity(key)
	details := fmt.Sprintf("Activity details are : ID : %d, Name : %s , Type : %s, Time : %s",actvity.ID, actvity.Name,actvity.Type,actvity.Time)
	w.Write(append([]byte(details)))

}
