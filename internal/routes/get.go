package routes

import (
	"fmt"
	"log"
	"net/http"
	"snapshot/internal/durable"
)

func Get(mux *http.ServeMux) {
	mux.HandleFunc("GET /{url}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		url, err := durable.URLDecode(r.PathValue("url"))

		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		fmt.Println(url)
	})
}
