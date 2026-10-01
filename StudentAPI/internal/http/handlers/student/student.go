package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/Vivek09Chahal/studentsAPI/internal/types"
	"github.com/Vivek09Chahal/studentsAPI/internal/utils/response"
	"github.com/go-playground/validator.git"
	"github.com/go-playground/validator/v10"
)

func New() http.HandlerFunc {
    return  func(w http.ResponseWriter, r *http.Request) {

        var student types.Student

        err := json.NewDecoder(r.Body).Decode(&student)

        if errors.Is(err, io.EOF)  {
            response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
            return
        }

        if err != nil {
            response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
            return
        }

        // request validate
        if err := validator.New().Struct(student); err !=  nil {

            validateErrs := err.(validator.ValidationErrors)
            response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrs))
            return
        }
        
        // slog.Info("creating a student")
        response.WriteJson(w, http.StatusCreated, map[string]string{"success": "ok"})  
    }
}