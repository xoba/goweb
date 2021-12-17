package main

import (
	"flag"
	"log"
	"net/http"
	"strings"

	"github.com/caddyserver/certmagic"
)

func main() {
	var hosts string
	var prod bool
	flag.BoolVar(&prod, "p", true, "production mode")
	flag.StringVar(&hosts, "h", "", "csv of valid hosts")
	flag.Parse()
	h := http.FileServer(http.Dir("content"))
	if prod {
		if err := Run(h, strings.Split(hosts, ",")...); err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal(http.ListenAndServe(":8080", h))
	}
}

func Run(h http.Handler, hosts ...string) error {
	certmagic.DefaultACME.Agreed = true
	return certmagic.HTTPS(hosts, h)
}
