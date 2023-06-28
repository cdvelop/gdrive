package gdrive

import (
	"log"
	"net/http"
)

func httpError(errorMessage string, w http.ResponseWriter) {
	err := "Error: " + errorMessage
	log.Println(err)
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(err))
}
