package main

import (
	"log"
	"net/http"
	"os"

	gh "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/miekg/dns"
	"github.com/play-with-docker/play-with-docker/config"
	"github.com/play-with-docker/play-with-docker/handlers"
	"github.com/play-with-docker/play-with-docker/services"
	"github.com/play-with-docker/play-with-docker/templates"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/urfave/negroni"
)

func main() {

	config.ParseFlags()

	bypassCaptcha := len(os.Getenv("GOOGLE_RECAPTCHA_DISABLED")) > 0

	// Start the DNS server
	dns.HandleFunc(".", handlers.DnsRequest)
	udpDnsServer := &dns.Server{Addr: ":53", Net: "udp"}
	go func() {
		err := udpDnsServer.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()
	tcpDnsServer := &dns.Server{Addr: ":53", Net: "tcp"}
	go func() {
		err := tcpDnsServer.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	server := services.CreateWSServer()
	server.On("connection", handlers.WS)
	server.On("error", handlers.WSError)

	err := services.LoadSessionsFromDisk()
	if err != nil && !os.IsNotExist(err) {
		log.Fatal("Error decoding sessions from disk ", err)
	}

	r := mux.NewRouter()
	corsRouter := mux.NewRouter()

	// Reverse proxy (needs to be the first route, to make sure it is the first thing we check)
	//proxyHandler := handlers.NewMultipleHostReverseProxy()
	//websocketProxyHandler := handlers.NewMultipleHostWebsocketReverseProxy()

	tcpHandler := handlers.NewTCPProxy()

	corsHandler := gh.CORS(gh.AllowCredentials(), gh.AllowedHeaders([]string{"x-requested-with", "content-type"}), gh.AllowedOrigins([]string{"*"}))

	// Specific routes
	r.Host(`{subdomain:.*}{node:pwd[0-9]{1,3}-[0-9]{1,3}-[0-9]{1,3}-[0-9]{1,3}}-{port:[0-9]*}.{tld:.*}`).Handler(tcpHandler)
	r.Host(`{subdomain:.*}{node:pwd[0-9]{1,3}-[0-9]{1,3}-[0-9]{1,3}-[0-9]{1,3}}.{tld:.*}`).Handler(tcpHandler)
	r.Host(`pwd{alias:.*}-{session:.*}-{port:[0-9]*}.{tld:.*}`).Handler(tcpHandler)
	r.Host(`pwd{alias:.*}-{session:.*}.{tld:.*}`).Handler(tcpHandler)
	r.HandleFunc("/ping", handlers.Ping).Methods("GET")
	corsRouter.HandleFunc("/instances/images", handlers.GetInstanceImages).Methods("GET")
	corsRouter.HandleFunc("/sessions/{sessionId}", handlers.GetSession).Methods("GET")
	corsRouter.HandleFunc("/sessions/{sessionId}/instances", handlers.NewInstance).Methods("POST")
	corsRouter.HandleFunc("/sessions/{sessionId}/instances/{instanceName}/uploads", handlers.FileUpload).Methods("POST")
	corsRouter.HandleFunc("/sessions/{sessionId}/instances/{instanceName}", handlers.DeleteInstance).Methods("DELETE")
	corsRouter.HandleFunc("/sessions/{sessionId}/instances/{instanceName}/exec", handlers.Exec).Methods("POST")

	h := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./www/index.html")
	}

	r.HandleFunc("/p/{sessionId}", h).Methods("GET")
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

	corsRouter.HandleFunc("/", handlers.NewSession).Methods("POST")

	n := negroni.Classic()
	r.PathPrefix("/").Handler(negroni.New(negroni.Wrap(corsHandler(corsRouter))))
	n.UseHandler(r)

	go func() {
		log.Println("Listening on port " + config.PortNumber)
		log.Fatal(http.ListenAndServe("0.0.0.0:"+config.PortNumber, n))
	}()

	// Now listen for TLS connections that need to be proxied
	handlers.StartTLSProxy(config.SSLPortNumber)
}
