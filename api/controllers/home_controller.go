package controllers

import (
	"github.com/INI-MED/gocrud/api/response"
	"net/http"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, "Welcome page!")

}
