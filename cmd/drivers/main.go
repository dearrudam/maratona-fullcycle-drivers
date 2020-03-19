package main

import (
	"github.com/dearrudam/maratona-fullcycle-drivers/pkg/drivers"
	"github.com/gorilla/mux"
	"net/http"
)

var avaliableDrivers drivers.Drivers

func init() {
	avaliableDrivers = drivers.LoadDrivers()
}

func listDrivers(w http.ResponseWriter, r *http.Request) {
	w.Write(avaliableDrivers.MarshallToJSON())
}

func getDriverByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	driver, err := avaliableDrivers.GetDriverByID(id)
	if(err!=nil){
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Write(driver.MarshallToJSON())
}

func asJSONContentType(fn func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter,r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fn(w, r)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/drivers", asJSONContentType(listDrivers))
	r.HandleFunc("/drivers/{id}", asJSONContentType(getDriverByID))
	http.ListenAndServe(":8081", r)
}
