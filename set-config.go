package gdrive

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2/google"
)

func (g *drive) setOauth2Config(r *http.Request, w http.ResponseWriter) error {

	n := data{
		Protocol: "http",
		Host:     r.Host,
	}

	if r.TLS != nil {
		n.Protocol = "https"
	}

	fmt.Printf("Protocol:[%v] Host:[%v]\n", n.Protocol, n.Host)

	if n.Protocol == "" || n.Host == "" {
		return fmt.Errorf("datos vacíos Protocol:%v Host:%v", n.Protocol, n.Host)
	}

	// pars := n.Parse(os.Getenv("GOOGLE_DEV_AUTH"))

	// var err error
	pars, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	g.Config, err = google.ConfigFromJSON(pars)
	if err != nil {
		return fmt.Errorf("google json config %v", err.Error())
	}

	return nil
}
