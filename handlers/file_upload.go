package handlers

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func FileUpload(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	sessionId := vars["sessionId"]
	instanceName := vars["instanceName"]

	s := core.SessionGet(sessionId)
	i := core.InstanceGet(s, instanceName)

	// allow up to 32 MB which is the default

	// has a url query parameter, ignore body
	if url := req.URL.Query().Get("url"); url != "" {
		err := core.InstanceUploadFromUrl(i, req.URL.Query().Get("url"))
		if err != nil {
			log.Println(err)
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		rw.WriteHeader(http.StatusOK)
		return
	} else {
		red, err := req.MultipartReader()
		if err != nil {
			log.Println(err)
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		r := req.URL.Query().Get("relative")

		for {
			p, err := red.NextPart()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Println(err)
				continue
			}

			if p.FileName() == "" {
				continue
			}

			if r != "" {
				err = core.InstanceUploadToCWDFromReader(i, p.FileName(), p)
				if err != nil {
					log.Println(err)
					rw.WriteHeader(http.StatusInternalServerError)
					return
				}
			} else {
				err = core.InstanceUploadFromReader(i, p.FileName(), "/var/run/pwd/uploads", p)
				if err != nil {
					log.Println(err)
					rw.WriteHeader(http.StatusInternalServerError)
					return
				}
			}

			log.Printf("Uploaded [%s] to [%s]\n", p.FileName(), i.Name)
		}
		rw.WriteHeader(http.StatusOK)
		return
	}

}
