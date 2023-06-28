package gdrive

import (
	"context"
	"fmt"
	"net/http"
)

func (g *drive) Callback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	receivedState := r.URL.Query().Get("state")

	fmt.Println("CODE: ", code)
	fmt.Println("STATE: ", receivedState)

	if receivedState != g.state {
		http.Error(w, "Invalid state", http.StatusBadRequest)
		return
	}

	token, err := g.Config.Exchange(context.TODO(), code)
	if err != nil {
		http.Error(w, "Error exchanging code for token", http.StatusInternalServerError)
		return
	}

	g.saveToken(token)

	g.newHttpClient(token)
	// Utiliza el token para acceder a los recursos protegidos
	// Aquí puedes realizar solicitudes HTTP con el token de acceso a tus APIs internas

	fmt.Println("Authentication successful!")

	http.Redirect(w, r, "/drive", 200)

}
