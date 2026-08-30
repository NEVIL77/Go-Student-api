package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/NEVIL77/students-api/internal/types"
	"github.com/NEVIL77/students-api/internal/utils/response"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student

		// json.NewDecoder() reads JSON data.
		// .Decode(&student) converts that JSON into a Go Student struct.
		// &student means give the address of student so Decode can fill it.
		err := json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return

		}

		slog.Info("yer le lovded")
		w.Write([]byte("Hello World"))

		response.WriteJson(w, http.StatusCreated, student)
	}
}
