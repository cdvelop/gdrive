package gdrive

import (
	"context"
	"fmt"
	"log"
	"net/http"

	drive_api "google.golang.org/api/drive/v2"
	"google.golang.org/api/option"
)

func (g *drive) Read(w http.ResponseWriter, r *http.Request) {
	fmt.Println("leyendo archivos de google drive...")

	ctx := context.Background()

	gdsrv, err := drive_api.NewService(ctx, option.WithHTTPClient(g.Client))
	if err != nil {
		log.Fatalf("No se puede obtener el cliente de Drive: %v", err)
	}

	f, err := gdsrv.Files.List().
		Fields("nextPageToken, files(id, name)").Do()
	if err != nil {
		log.Fatalf("No se pueden recuperar los archivos: %v", err)
	}

	fmt.Println("Archivos:")
	if len(f.Items) == 0 {
		fmt.Println("No se encontraron archivos.")
	} else {
		for _, i := range f.Items {
			fmt.Printf("%s (%s)\n", i.Etag, i.Id)
		}
	}
}
