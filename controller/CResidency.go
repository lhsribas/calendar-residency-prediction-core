package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/lhsribas/calendar-residency-prediction-core/model"
	"github.com/lhsribas/calendar-residency-prediction-core/repository"
)

var s = repository.SResidency{}

func respondWithERROR(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, msg)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func SaveCustomer(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var customer model.Customer

	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		respondWithERROR(w, http.StatusBadRequest, "Invalid request payload.")
		return
	}

	Id, err := s.SaveCustomer(customer, "root")

	if err != nil {
		respondWithERROR(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"id": Id.Hex()})
}

func SaveResidency(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var residency model.Residency

	if err := json.NewDecoder(r.Body).Decode(&residency); err != nil {
		respondWithERROR(w, http.StatusBadRequest, "Invalid request payload.")
		return
	}

	Id, err := s.SaveResidency(residency, "root")

	if err != nil {
		respondWithERROR(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"id": Id.Hex()})
}

func SaveStaffMamber(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var member model.StaffMember

	if err := json.NewDecoder(r.Body).Decode(&member); err != nil {
		respondWithERROR(w, http.StatusBadRequest, "Invalid request payload.")
		return
	}

	Id, err := s.SaveStaffMember(member, "root")

	if err != nil {
		respondWithERROR(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"id": Id.Hex()})
}

func SaveResidencyNotes(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var note model.Note

	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		respondWithERROR(w, http.StatusBadRequest, "Invalid request payload.")
		return
	}

	Id, err := s.SaveResidencyNotes(note, "root")

	if err != nil {
		respondWithERROR(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"id": Id.Hex()})
}
