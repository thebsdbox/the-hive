package handlers

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/crypto/acme/autocert"

	"github.com/googollee/go-socket.io"
	gh "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/play-with-docker/play-with-docker/config"
	"github.com/play-with-docker/play-with-docker/event"
	"github.com/play-with-docker/play-with-docker/pwd"
	"github.com/play-with-docker/play-with-docker/templates"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/urfave/negroni"
)

var core pwd.PWDApi
var e event.EventApi
var ws *socketio.Server

func Bootstrap(c pwd.PWDApi, ev event.EventApi) {
	core = c
	e = ev
}

func Register() {

	bypassCaptcha := len(os.Getenv("GOOGLE_RECAPTCHA_DISABLED")) > 0

	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	server.On("connection", WS)
	server.On("error", WSError)

	RegisterEvents(server)

	r := mux.NewRouter()
	corsRouter := mux.NewRouter()

	corsHandler := gh.CORS(gh.AllowCredentials(), gh.AllowedHeaders([]string{"x-requested-with", "content-type"}), gh.AllowedOrigins([]string{"*"}))

	// Specific routes
	r.HandleFunc("/ping", Ping).Methods("GET")
	corsRouter.HandleFunc("/instances/images", GetInstanceImages).Methods("GET")
	corsRouter.HandleFunc("/sessions/{sessionId}", GetSession).Methods("GET")
	corsRouter.HandleFunc("/sessions/{sessionId}/setup", SessionSetup).Methods("POST")
	corsRouter.HandleFunc("/sessions/{sessionId}/instances", NewInstance).Methods("POST")
	corsRouter.HandleFunc("/sessions/{sessionId}/instances/{instanceName}/uploads", FileUpload).Methods("POST")
	corsRouter.HandleFunc("/sessions/{sessionId}/instances/{instanceName}", DeleteInstance).Methods("DELETE")
	corsRouter.HandleFunc("/sessions/{sessionId}/instances/{instanceName}/exec", Exec).Methods("POST")

	r.HandleFunc("/p/{sessionId}", Home).Methods("GET")
	r.PathPrefix("/assets").Handler(http.FileServer(http.Dir("./www")))
	r.HandleFunc("/robots.txt", func(rw http.ResponseWriter, r *http.Request) {
		http.ServeFile(rw, r, "www/robots.txt")
	})
	r.HandleFunc("/sdk.js", func(rw http.ResponseWriter, r *http.Request) {
		http.ServeFile(rw, r, "www/sdk.js")
	})

	corsRouter.Handle("/sessions/{sessionId}/ws/", server)
	r.Handle("/metrics", promhttp.Handler())

	// Generic routes
	r.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		if bypassCaptcha {
			http.ServeFile(rw, r, "./www/bypass.html")
		} else {
			welcome, tmplErr := templates.GetWelcomeTemplate()
			if tmplErr != nil {
				log.Fatal(tmplErr)
			}
			rw.Write(welcome)
		}
	}).Methods("GET")

	corsRouter.HandleFunc("/", NewSession).Methods("POST")

	n := negroni.Classic()
	r.PathPrefix("/").Handler(negroni.New(negroni.Wrap(corsHandler(corsRouter))))
	n.UseHandler(r)

	httpServer := http.Server{
		Addr:              "0.0.0.0:" + config.PortNumber,
		Handler:           n,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
	}
	if config.UseLetsEncrypt {
		certManager := autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist(config.LetsEncryptDomains...),
			Cache:      autocert.DirCache(config.LetsEncryptCertsDir),
		}

		httpServer.TLSConfig = &tls.Config{
			GetCertificate: certManager.GetCertificate,
		}
		log.Println("Listening on port " + config.PortNumber)
		log.Fatal(httpServer.ListenAndServeTLS("", ""))
	} else {
		log.Println("Listening on port " + config.PortNumber)
		log.Fatal(httpServer.ListenAndServe())
	}

}

func RegisterEvents(s *socketio.Server) {
	ws = s
	e.OnAny(broadcastEvent)
}

func broadcastEvent(eventType event.EventType, sessionId string, args ...interface{}) {
	ws.BroadcastTo(sessionId, eventType.String(), args...)
}
