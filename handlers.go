package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Handlers struct {
	DB *DBHandler
}

func New(db *DBHandler) *Handlers {
	return &Handlers{DB: db}
}

func (han *Handlers) HandleGet(w http.ResponseWriter, r *http.Request) {
	packages, err := han.DB.GetPackages("github.com/cli/cli")
	if err != nil {
		AppLogger.Sugar().Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(packages)
	if err != nil {
		AppLogger.Sugar().Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (han *Handlers) HandlePost(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	var packet Check
	err := json.NewDecoder(r.Body).Decode(&packet)
	if err != nil {
		AppLogger.Sugar().Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = han.DB.Insert(name, packet)
	if err != nil {
		AppLogger.Sugar().Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (han *Handlers) HandleDelete(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	err := han.DB.DeletePackage(name)
	if errors.Is(err, ErrNoFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		AppLogger.Sugar().Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
