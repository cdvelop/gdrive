package gdrive

import (
	"bytes"
	"log"
	"text/template"
)

type data struct {
	Protocol string //ej: http, https
	Host     string //ej: localhost:8080,172.0.0.2:200,web.com
}

//layout: {"web":{"client_id":"XXX","project_id":"XXX","auth_uri":"https://accounts.google.com/o/oauth2/auth","token_uri":"https://oauth2.googleapis.com/token","auth_provider_x509_cert_url":"https://www.googleapis.com/oauth2/v1/certs","client_secret":"XXXX",
// "redirect_uris":["{{.Protocol}}://{{.Host}}/callback"]}}

func (d data) Parse(layout string) []byte {

	t, err := template.New("").Parse(string(layout))
	if err != nil {
		log.Fatal(err)
		return nil
	}
	var buf bytes.Buffer
	err = t.Execute(&buf, d)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return buf.Bytes()
}
