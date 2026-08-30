package student

import (
	"log/slog"
	"net/http"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		slog.Info("yer le lovded")
		w.Write([]byte("Hello World"))
	}
}
