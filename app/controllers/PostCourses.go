package controllers

import (
	"fmt"
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"../models"
	"../utils"
)

func AddCourse(db *gorm.DB, w http.ResponseWriter, r *http.Response) {
	decoder := json.NewDecoder(r.Body)
	var project models.Course
	if err := decoder.Decode(&project); err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	project.Key = utils.GenerateId()
	if err := models.AddCourse(db, &project); err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, project)
}

func RetakeCourse(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var courses models.Course

	id, _ := strconv.Atoi(vars["key"])
	err := models.GetACourse(db, id, &courses)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	decoder := json.NewDecoder(r.Body)
	if err = decoder.Decode(&courses); err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	if err = models.RetakeCourse(db, &courses); err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, courses)
}

func DropCourse(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var courses models.Course
	id, _ := strconv.Atoi(vars["key"])
	if err := models.GetACourse(db, id, &courses); err != nil {
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if err := models.DropCourse(db, &courses); err != nil {
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, courses)
}

func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
	fmt.Println("status ", status)
	var res utils.ResponseData

	res.Status = status
	res.Meta = utils.ResponseMessage(status)
	res.Data = payload

	response, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

func RespondError(w http.ResponseWriter, status int, message string) {
	var res utils.ResponseData
	rescode := utils.ResponseMessage(status)
	res.Status = status
	res.Meta = rescode
	response, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}
