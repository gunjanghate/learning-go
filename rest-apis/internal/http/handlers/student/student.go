package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gunjanghate/learning-go/internal/types"
	"github.com/gunjanghate/learning-go/internal/utils/response"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Handling request for /api/students")
		// 		{
		//     "name": "GG",
		//     "email": "gg@gmail.com",
		//     "age" : 20
		// }  // this is coming in request body

		var stu types.Student

		err := json.NewDecoder(r.Body).Decode(&stu)
		if errors.Is(err, io.EOF){
			response.WriteJson(w, http.StatusBadRequest,response.GenError(fmt.Errorf("empty body")))
			return
		}

		if err!=nil{
			response.WriteJson(w, http.StatusBadRequest, response.GenError((err)))
			slog.Error("error decoding request body", slog.String("error", err.Error()))
			return
		}

		// validate request body
		if err := validator.New().Struct(stu); err != nil{
			validateErrs := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidationErr(validateErrs))
			return
		}
		
		response.WriteJson(w, http.StatusCreated, map[string]string{"success": "OK", "name": stu.Name, "email": stu.Email})


	}
}
