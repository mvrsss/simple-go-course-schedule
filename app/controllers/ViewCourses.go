package controllers

import (
	"net/http"
	"strconv"

	"../models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func ListCourse(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	courses := []models.Course{}
	if err := models.GetAllCourses(db, &courses); err != nil {
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, courses)
}

func OneCourse(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var courses models.Course
	id, _ := strconv.Atoi(vars["key"])
	if err := models.GetACourse(db, id, &courses); err != nil {
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, courses)
  return
}
