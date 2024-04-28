package routes

import (
	"fmt"
	"gorm.io/gorm/clause"
	"log"
	"net/http"
	"snapshot/internal/durable"
	"snapshot/internal/model"
)

func Post(mux *http.ServeMux) {
	mux.HandleFunc("POST /{url}", func(w http.ResponseWriter, r *http.Request) {
		url, err := durable.URLDecode(r.PathValue("url"))

		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		image, err := durable.CompressedImage(url, 90)

		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		snapshot := &model.Snapshot{
			Url:   url,
			Image: image,
		}

		result := durable.Connection().Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "url"}},
			DoUpdates: clause.AssignmentColumns([]string{"image", "updated_at"}),
		}).Create(&snapshot)

		if result.Error != nil {
			log.Println(result.Error)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "image/jpeg")
		w.Header().Add("Content-Length", fmt.Sprintf("%d", len(image)))
		w.Header().Add("Content-Encoding", "gzip")
		w.WriteHeader(http.StatusCreated)

		if _, err := w.Write(image); err != nil {
			log.Fatal(err)
		}
	})
}
