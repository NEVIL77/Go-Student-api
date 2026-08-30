package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/NEVIL77/students-api/internal/types"
	"github.com/NEVIL77/students-api/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		slog.Info("creating a student")

		var student types.Student

		// json.NewDecoder() reads JSON data.
		// .Decode(&student) converts that JSON into a Go Student struct.
		// &student means give the address of student so Decode can fill it.
		err := json.NewDecoder(r.Body).Decode(&student)

		// if body dont pass then this error
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		if err := validator.New().Struct(student); err != nil {
			validateErrs := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrs))
			return
		}

		response.WriteJson(w, http.StatusCreated, student)
	}
}
