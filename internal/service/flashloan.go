package service

import (
	"context"
	"flashLoanSimple/internal/adapter"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type FlashLoanService struct {
	Uniswap   *adapter.UniswapAdapter
	SushiSwap *adapter.SushiSwapAdapter
}

func NewFlashLoanService(uniswap *adapter.UniswapAdapter, sushi *adapter.SushiSwapAdapter) *FlashLoanService {
	return &FlashLoanService{
		Uniswap:   uniswap,
		SushiSwap: sushi,
	}
}

func (s *FlashLoanService) DoUniswapFlashLoan(ctx context.Context, amount *big.Int, receiver common.Address, data []byte) error {
	return s.Uniswap.FlashLoan(ctx, amount, receiver, data)
}

func (s *FlashLoanService) DoSushiSwapFlashLoan(ctx context.Context, amount *big.Int, receiver common.Address, data []byte) error {
	return s.SushiSwap.FlashLoan(ctx, amount, receiver, data)
}
