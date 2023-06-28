package gdrive

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"golang.org/x/oauth2"
)

// Guarda un token en una ruta de archivo.
func (g drive) saveToken(token *oauth2.Token) {
	fmt.Printf("Guardando el archivo de credenciales en: %s\n", g.token_file)
	f, err := os.OpenFile(g.token_file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("No se puede almacenar en caché el token de OAuth: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

// Recupera un token desde un archivo local.
func (g drive) tokenFromFile() (*oauth2.Token, error) {
	f, err := os.Open(g.token_file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Solicita un token desde la web y luego devuelve el token recuperado.
func (g drive) getTokenFromWeb() *oauth2.Token {
	authURL := g.Config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Ve al siguiente enlace en tu navegador y luego escribe el "+
		"código de autorización: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("No se puede leer el código de autorización %v", err)
	}

	tok, err := g.Config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("No se puede obtener el token desde la web %v", err)
	}
	return tok
}
