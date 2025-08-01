package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Dashboard represents a scalable API service dashboard
type Dashboard struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Endpoints int    `json:"endpoints,omitempty"`
}

// DashboardService provides operations for the dashboard
type DashboardService interface {
	CreateDashboard(name string) (*Dashboard, error)
	GetDashboard(id string) (*Dashboard, error)
	UpdateDashboard(dashboard *Dashboard) error
	DeleteDashboard(id string) error
}

type dashboardService struct{}

func (ds *dashboardService) CreateDashboard(name string) (*Dashboard, error) {
	dashboard := &Dashboard{
		ID:   generateID(),
		Name: name,
	}
	// logic to create dashboard
	return dashboard, nil
}

func (ds *dashboardService) GetDashboard(id string) (*Dashboard, error) {
	// logic to get dashboard
	return &Dashboard{
		ID:   id,
		Name: "mocked-name",
	}, nil
}

func (ds *dashboardService) UpdateDashboard(dashboard *Dashboard) error {
	// logic to update dashboard
	return nil
}

func (ds *dashboardService) DeleteDashboard(id string) error {
	// logic to delete dashboard
	return nil
}

func generateID() string {
	// logic to generate unique id
	return "mocked-id"
}

func main() {
	ds := &dashboardService{}
	router := mux.NewRouter()

	router.HandleFunc("/api/dashboards", createDashboardHandler(ds)).Methods("POST")
	router.HandleFunc("/api/dashboards/{id}", getDashboardHandler(ds)).Methods("GET")
	router.HandleFunc("/api/dashboards/{id}", updateDashboardHandler(ds)).Methods("PUT")
	router.HandleFunc("/api/dashboards/{id}", deleteDashboardHandler(ds)).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}

func createDashboardHandler(ds DashboardService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dashboard *Dashboard
		err := json.NewDecoder(r.Body).Decode(&dashboard)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		d, err := ds.CreateDashboard(dashboard.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(d)
	}
}

func getDashboardHandler(ds DashboardService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		d, err := ds.GetDashboard(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(d)
	}
}

func updateDashboardHandler(ds DashboardService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		var dashboard *Dashboard
		err := json.NewDecoder(r.Body).Decode(&dashboard)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = ds.UpdateDashboard(dashboard)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

func deleteDashboardHandler(ds DashboardService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		err := ds.DeleteDashboard(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}