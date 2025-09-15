package api

import (
	"encoding/json"
	"flashLoanSimple/internal/service"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"net/http"
)

type FlashLoanRequest struct {
	Protocol string `json:"protocol"` // "uniswap" 或 "sushiswap"
	Amount   string `json:"amount"`
	Receiver string `json:"receiver"`
	Data     string `json:"data"`
}

type Handler struct {
	Service *service.FlashLoanService
}

func NewHandler(svc *service.FlashLoanService) *Handler {
	return &Handler{Service: svc}
}

func (h *Handler) FlashLoan(w http.ResponseWriter, r *http.Request) {
	var req FlashLoanRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid request"))
		return
	}
	amount, ok := new(big.Int).SetString(req.Amount, 10)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid amount"))
		return
	}
	receiver := common.HexToAddress(req.Receiver)
	data := common.FromHex(req.Data)
	var err error
	if req.Protocol == "uniswap" {
		err = h.Service.DoUniswapFlashLoan(r.Context(), amount, receiver, data)
	} else if req.Protocol == "sushiswap" {
		err = h.Service.DoSushiSwapFlashLoan(r.Context(), amount, receiver, data)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("unsupported protocol"))
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("flashloan success"))
}
