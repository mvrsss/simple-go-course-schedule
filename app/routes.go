package app

import (
	"fmt"
	"log"
	"net/http"

	"../config"
	"../migrate"
	"./controllers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Name,
		config.DB.Charset)

	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal("Could not connect database")
	}

	a.DB = migrate.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

func (a *App) setRouters() {
	a.Post("/course", a.AddCourse)
	a.Get("/course", a.ListCourse)
	a.Get("/course/{key:[1-9]+}", a.OneCourse)
	a.Put("/course/{key:[1-9]+}", a.RetakeCourse)
	a.Delete("/course/{key:[1-9]+}", a.DropCourse)
}

func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("Get")
}

func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("Put")
}

func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("Delete")
}

func (a *App) AddCourse(w http.ResponseWriter, r *http.Request) {
	controllers.AddCourse(a.DB, w, r)
}

func (a *App) ListCourse(w http.ResponseWriter, r *http.Request) {
	controllers.ListCourse(a.DB, w, r)
}

func (a *App) OneCourse(w http.ResponseWriter, r *http.Request) {
	controllers.OneCourse(a.DB, w, r)
}

func (a *App) RetakeCourse(w http.ResponseWriter, r *http.Request) {
	controllers.RetakeCourse(a.DB, w, r)
}

func (a *App) DropCourse(w http.ResponseWriter, r *http.Request) {
	controllers.DropCourse(a.DB, w, r)
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
