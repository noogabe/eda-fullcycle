package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	usecase "github.com.br/noogabe/eda-fullcycle/consumer/internal/usecase/get"
	"github.com/go-chi/chi/v5"
)

type WebGetBalanceHandler struct {
	GetBalanceUsecase usecase.GetBalanceUsecase
}

func NewWebGetBalanceHandler(getBalanceUsecase usecase.GetBalanceUsecase) *WebGetBalanceHandler {
	return &WebGetBalanceHandler{
		GetBalanceUsecase: getBalanceUsecase,
	}
}

func (h *WebGetBalanceHandler) GetBalance(w http.ResponseWriter, r *http.Request) {
	var input usecase.GetBalanceInputDto

	accountId := chi.URLParam(r, "id")

	if accountId == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("account id is required"))
		fmt.Println("account id is required")
		return
	}

	input.AccountId = accountId

	output, err := h.GetBalanceUsecase.Execute(input)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		fmt.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
