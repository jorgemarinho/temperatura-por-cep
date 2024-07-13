package web

import (
	"encoding/json"
	"net/http"

	"github.com/jorgemarinho/temperatura-por-cep/internal/dto"
	"github.com/jorgemarinho/temperatura-por-cep/internal/errors"
	"github.com/jorgemarinho/temperatura-por-cep/internal/usecase"
)

func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {
	cepParam := r.URL.Query().Get("cep")

	if cepParam == "" {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	if len(cepParam) < 8 {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	buscaCepInputDTO := dto.BuscaCepInputDTO{Cep: cepParam}

	newBuscaCepUseCase := usecase.NewBuscaCepUseCase(buscaCepInputDTO)

	cep, err := newBuscaCepUseCase.Execute()

	if err != nil {
		code := http.StatusInternalServerError
		message := err.Error()

		if httpErr, ok := err.(*errors.HTTPError); ok {
			code = httpErr.Code
			message = httpErr.Message
		}

		http.Error(w, message, code)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cep)
}
