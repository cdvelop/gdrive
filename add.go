package gdrive

import (
	"net/http"

	"golang.org/x/oauth2"
)

type drive struct {
	state string
	mux   *http.ServeMux
	*oauth2.Config
	*http.Client
	data
	token_file string
}

func Add(mux *http.ServeMux) {

	gd := drive{
		mux:        mux,
		token_file: "token.json",
	}

	gd.addHttpHandlers()

}
