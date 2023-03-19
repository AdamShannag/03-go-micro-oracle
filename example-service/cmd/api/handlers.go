package main

import (
	"net/http"

	"github.com/AdamShannag/toolkit/v2"
	"github.com/go-chi/chi"
)

func (app *Config) createUser(w http.ResponseWriter, req *http.Request) {
	var t toolkit.Tools
	user := app.Models.User

	err := t.ReadJSON(w, req, &user)
	if err != nil {
		t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	_, err = user.Insert(user)
	if err != nil {
		t.ErrorJSON(w, err)
		return
	}

	err = t.WriteJSON(w, http.StatusAccepted, toolkit.JSONResponse{
		Message: "User added!",
		Error:   false,
	})

	if err != nil {
		t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
}

func (app *Config) getUser(w http.ResponseWriter, req *http.Request) {
	var t toolkit.Tools
	user := app.Models.User

	id := chi.URLParam(req, "id")

	u, err := user.GetOne(id)
	if err != nil {
		t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = t.WriteJSON(w, http.StatusAccepted, u)
	if err != nil {
		t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
}

func (app *Config) getUsers(w http.ResponseWriter, req *http.Request) {
	var t toolkit.Tools
	user := app.Models.User

	users, err := user.GetAll()
	if err != nil {
		t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = t.WriteJSON(w, http.StatusAccepted, users)
	if err != nil {
		t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
}

func (app *Config) updateUser(w http.ResponseWriter, req *http.Request) {
	var t toolkit.Tools
	user := app.Models.User

	err := t.ReadJSON(w, req, &user)
	if err != nil {
		t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = user.Update()
	if err != nil {
		t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = t.WriteJSON(w, http.StatusAccepted, toolkit.JSONResponse{
		Message: "User updated!",
		Error:   false,
	})

	if err != nil {
		t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

}

func (app *Config) deleteUser(w http.ResponseWriter, req *http.Request) {
	var t toolkit.Tools
	user := app.Models.User

	id := chi.URLParam(req, "id")

	err := user.DeleteByID(id)
	if err != nil {
		t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = t.WriteJSON(w, http.StatusAccepted, toolkit.JSONResponse{
		Message: "User updated!",
		Error:   false,
	})

	if err != nil {
		t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
}
