package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/play-with-docker/play-with-docker/provisioner"
	"github.com/play-with-docker/play-with-docker/pwd"
	"github.com/play-with-docker/play-with-docker/pwd/types"
)

func NewInstance(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	sessionId := vars["sessionId"]

	body := types.InstanceConfig{PlaygroundFQDN: req.Host}

	json.NewDecoder(req.Body).Decode(&body)

	s := core.SessionGet(sessionId)

	i, err := core.InstanceNew(s, body)
	if err != nil {
		if pwd.SessionComplete(err) {
			rw.WriteHeader(http.StatusConflict)
			return
		} else if provisioner.OutOfCapacity(err) {
			rw.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprintln(rw, `{"error": "out_of_capacity"}`)
			return
		}
		log.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
		//TODO: Set a status error
	} else {
		json.NewEncoder(rw).Encode(i)
	}
}
