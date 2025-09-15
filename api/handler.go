package api

import (
	"encoding/json"
	"flashLoanSimple/internal/service"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"net/http"
)

// Handler 用于处理 HTTP 请求，连接 API 层与业务逻辑层
// Service 字段为闪电贷业务服务
type Handler struct {
	Service *service.FlashLoanService
}

// FlashLoanRequest 闪电贷请求参数
// Protocol: 协议类型，"uniswap" 或 "sushiswap"
// Amount: 闪电贷金额
// Receiver: 闪电贷接收地址
// Data: 附加数据
type FlashLoanRequest struct {
	Protocol string `json:"protocol"` // "uniswap" 或 "sushiswap"
	Amount   string `json:"amount"`
	Receiver string `json:"receiver"`
	Data     string `json:"data"`
}

// NewHandler 创建一个新的 Handler 实例
// 参数 svc: 闪电贷业务服务
// 返回值: Handler 实例
func NewHandler(svc *service.FlashLoanService) *Handler {
	return &Handler{Service: svc}
}

// FlashLoan 处理闪电贷 HTTP 请求
// 解析请求参数，调用对应的闪电贷服务，返回执行结果
// 参数 w: HTTP 响应 writer
// 参数 r: HTTP 请求
func (h *Handler) FlashLoan(w http.ResponseWriter, r *http.Request) {
	var req FlashLoanRequest
	// 解析请求体
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid request"))
		return
	}
	// 金额字符串转为大整数
	amount, ok := new(big.Int).SetString(req.Amount, 10)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid amount"))
		return
	}
	// 解析接收地址和附加数据
	receiver := common.HexToAddress(req.Receiver)
	data := common.FromHex(req.Data)
	var err error
	// 根据协议类型调用不同的闪电贷服务
	if req.Protocol == "uniswap" {
		err = h.Service.DoUniswapFlashLoan(r.Context(), amount, receiver, data)
	} else if req.Protocol == "sushiswap" {
		err = h.Service.DoSushiSwapFlashLoan(r.Context(), amount, receiver, data)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("unsupported protocol"))
		return
	}
	// 返回执行结果
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("flashloan success"))
}
