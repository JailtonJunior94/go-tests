package controller

import (
	"encoding/json"
	"net/http"

	"github.com/jailtonjunior94/go-tests/internal/infra/database"
	"github.com/jailtonjunior94/go-tests/internal/usecases"
)

func (b *BaseHandler) CreateClientHander(w http.ResponseWriter, r *http.Request) {
	var content usecases.CreateClientInputDTO
	err := json.NewDecoder(r.Body).Decode(&content)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	repository := database.NewClientRepository(b.db)
	usecase := usecases.NewCreateClientUseCase(repository)

	_, err = usecase.Execute(&content)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
