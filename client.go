package gdrive

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"
)

// Recupera un token, guarda el token y luego devuelve el cliente generado.
func (g *drive) getClient() {
	// El archivo token.json almacena los tokens de acceso y actualización del usuario,
	// y se crea automáticamente cuando se completa el flujo de autorización por primera vez.
	token, err := g.tokenFromFile()
	if err != nil {
		token = g.getTokenFromWeb()
		g.saveToken(token)
	}
	g.newHttpClient(token)
}

func (g *drive) newHttpClient(token *oauth2.Token) {
	fmt.Println("nuevo http client")
	g.Client = g.Config.Client(context.Background(), token)
}
