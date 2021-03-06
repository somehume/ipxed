package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// CreateMachineHandler creates a new machine.
func CreateMachineHandler(w http.ResponseWriter, r *http.Request) {
	m := Machine{}
	m.Name = r.PostFormValue("name")
	m.MacAddress = r.PostFormValue("macaddress")
	m.Profile = r.PostFormValue("profile")
	m.Save()
	http.Redirect(w, r, "/machines/"+m.Name, http.StatusMovedPermanently)
}

// DeleteMachineHandler deletes a specific machine.
func DeleteMachineHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	DeleteMachineByName(vars["name"])
	http.Redirect(w, r, "/machines", http.StatusMovedPermanently)
}

// EditMachineHandler displays an HTML form for editing a machine.
func EditMachineHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	m, err := GetMachineByName(vars["name"])
	if err != nil {
		log.Println(err.Error())
	}
	p := &Page{
		Data:  m,
		Title: fmt.Sprintf("Machine - %s", m.Name),
	}
	renderTemplate(w, "templates/machines/edit.html", p)
}

// IndexMachineHandler displays a list of all machines.
func IndexMachineHandler(w http.ResponseWriter, r *http.Request) {
	machines, err := GetMachines()
	if err != nil {
		log.Println(err.Error())
	}
	p := &Page{
		Data:  machines,
		Title: "Machines",
	}
	renderTemplate(w, "templates/machines/index.html", p)
}

// NewMachineHandler displays an HTML form for creating a machine.
func NewMachineHandler(w http.ResponseWriter, r *http.Request) {
	profiles, err := GetProfiles()
	if err != nil {
		log.Println(err.Error())
	}
	data := struct {
		Profiles []Profile
	}{
		Profiles: profiles,
	}
	p := &Page{
		Data:  data,
		Title: "Create a Machine",
	}
	renderTemplate(w, "templates/machines/create.html", p)
}

// ShowMachineHandler displays a specific machine.
func ShowMachineHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	m, err := GetMachineByName(vars["name"])
	if err != nil {
		log.Println(err.Error())
	}
	p := &Page{
		Data:  m,
		Title: fmt.Sprintf("Machines - %s", m.Name),
	}
	renderTemplate(w, "templates/machines/show.html", p)
}
