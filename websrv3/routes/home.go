package routes

import (
	"encoding/json"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	log.Printf("\n Method: [%v] - Route: [%v] - Remote Address: [%v]", r.Method, r.RequestURI, r.RemoteAddr)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	response := Response{
		Msg: "server running",
	}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
