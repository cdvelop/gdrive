package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/cdvelop/auth"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

var auth_handler *auth.Auth

func main() {

	mux := http.NewServeMux()

	auth_handler = auth.Add("/home", nil, false, mux, []string{"google_dev"})
	// mux.HandleFunc("/login", Login)
	// mux.HandleFunc("/callback", Callback)
	mux.HandleFunc("/home", Home)

	fmt.Println("API SERVER RUN :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}

}

func Home(w http.ResponseWriter, r *http.Request) {

	client, redirect, err := auth_handler.GoogleDev(r.Host).GetHttpClient(r)
	if err != nil {
		http.Redirect(w, r, redirect, http.StatusTemporaryRedirect)
	} else {

		ctx := context.Background()

		srv, err := drive.NewService(ctx, option.WithHTTPClient(client))
		if err != nil {
			log.Fatalf("Unable to retrieve Drive client: %v", err)
		}

		f, err := srv.Files.List().PageSize(10).
			Fields("nextPageToken, files(id, name)").Do()
		if err != nil {
			log.Fatalf("Unable to retrieve files: %v", err)
		}
		fmt.Fprintln(w, "Files:")
		if len(f.Files) == 0 {
			fmt.Fprint(w, "No files found.")
		} else {
			for _, i := range f.Files {
				fmt.Fprintf(w, "%s (%s)\n", i.Name, i.Id)
			}
		}

	}
}
