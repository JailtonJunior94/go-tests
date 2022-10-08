package controller

import (
	"fmt"
	"net/http"

	"github.com/jailtonjunior94/go-tests/internal/usecases"
)

func (b *BaseHandler) CreateClientHander(w http.ResponseWriter, r *http.Request) {
	var content usecases.CreateClientInputDTO
	fmt.Println(content)

}
