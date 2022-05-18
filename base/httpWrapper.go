package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

type pluginRegister struct {
	Number int `json:"number"`
}

func register(w http.ResponseWriter, r *http.Request) {
	// time start
	start := time.Now()

	// get json from request

	// get body from request as string

	var pr pluginRegister
	err := json.NewDecoder(r.Body).Decode(&pr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Debug().Msgf("%+v", pr)

	success(w)

	// time end
	end := time.Now()
	elapsed := end.Sub(start)
	log.Debug().Timestamp().Int("elapsed-micro", int(elapsed.Microseconds())).Str("API-Method", "register").Msg("Time elapsed")
}

func success(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}