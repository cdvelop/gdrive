package gdrive

import (
	"net/http"
)

func (g *drive) Login(w http.ResponseWriter, r *http.Request) {
	// err := g.setOauth2Config(r, w)
	// if err != nil {
	// 	httpError(err.Error(), w)
	// 	return
	// }

	// token, err := g.tokenFromFile()
	// if err != nil {
	// 	fmt.Println("no existe token almacenado: ", err)
	// 	fmt.Println("creando nueva solicitud token")
	// 	g.state = buildUniqueKey(32) // Genera un estado aleatorio de 32 bytes para proteger contra ataques CSRF

	// 	authURL := g.Config.AuthCodeURL(g.state)
	// 	http.Redirect(w, r, authURL, http.StatusTemporaryRedirect)
	// } else {
	// 	// verificar la valides del token aquí

	// 	// crear nuevo cliente http
	// 	g.Client = g.Config.Client(context.Background(), token)
	// 	g.newHttpClient(token)

	// 	http.Redirect(w, r, "/drive", 200)
	// }

}
