package web

import (
	"encoding/json"
	"net/http"

	"github.com/jorgemarinho/temperatura-por-cep/internal/dto"
	"github.com/jorgemarinho/temperatura-por-cep/internal/usecase"
)

func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {
	cepParam := r.URL.Query().Get("cep")

	if cepParam == "" {
		http.Error(w, "invalid zipcode", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(cepParam) < 8 {
		http.Error(w, "invalid zipcode", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	buscaCepInputDTO := dto.BuscaCepInputDTO{Cep: cepParam}

	newBuscaCepUseCase := usecase.NewBuscaCepUseCase(buscaCepInputDTO)

	cep, err := newBuscaCepUseCase.Execute()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cep)
}
