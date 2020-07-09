package main

import (
	"crypto/tls"
	"log"
	"net/http"

	"../API-Go-Blockchain/routers"
)

func main() {

	router := routers.InitRoutes()

	cfg := &tls.Config{
		// Version minima que se acepta de tls
		MinVersion: tls.VersionSSL30,
		MaxVersion: tls.VersionTLS12,
		// Causes servers to use Go's default ciphersuite preferences,
		// which are tuned to avoid attacks. Does nothing on clients.
		PreferServerCipherSuites: true,
		InsecureSkipVerify:       true,
		//CurvePreferences to avoid unoptimized curves: a client using CurveP384
		//would cause up to a second of CPU to be consumed on our machines
		//CurvePreferences: []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
		//If you can take the compatibility loss of the Modern configuration, you should then also set
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305, // Go 1.8 only
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,   // Go 1.8 only
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			// Best disabled, as they don't provide Forward Secrecy,
			// but might be necessary for some clients
			// tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			// tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
		},
	}

	/*cer, err := tls.LoadX509KeyPair("server.pem", "server.key")
	if err != nil {
		log.Println(err)
		return
	}

	cfg := &tls.Config{Certificates: []tls.Certificate{cer}}*/

	srv := &http.Server{
		Addr:         ":443",
		Handler:      router,
		TLSConfig:    cfg,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}

	//log.Fatal(http.ListenAndServe(":8081", http.HandlerFunc(redirect)))
	log.Fatal(srv.ListenAndServeTLS("server.pem", "server.key"))

}
