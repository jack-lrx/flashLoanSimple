package service

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gavin/flashLoanSimple/internal/adapter"
	"math/big"
)

// FlashLoanService 闪电贷业务服务，统一封装 Uniswap 和 SushiSwap 闪电贷逻辑
// Uniswap: Uniswap 闪电贷适配器
// SushiSwap: SushiSwap 闪电贷适配器
type FlashLoanService struct {
	Uniswap   *adapter.UniswapAdapter
	SushiSwap *adapter.SushiSwapAdapter
}

// NewFlashLoanService 创建闪电贷业务服务
// 参数 uniswap: Uniswap 闪电贷适配器
// 参数 sushi: SushiSwap 闪电贷适配器
// 返回值: 闪电贷业务服务实例
func NewFlashLoanService(uniswap *adapter.UniswapAdapter, sushi *adapter.SushiSwapAdapter) *FlashLoanService {
	return &FlashLoanService{
		Uniswap:   uniswap,
		SushiSwap: sushi,
	}
}

// DoUniswapFlashLoan 执行 Uniswap 闪电贷业务
// 参数 ctx: 上下文
// 参数 amount: 闪电贷金额
// 参数 receiver: 闪电贷接收地址
// 参数 data: 附加数据
// 返回值: 错误信息
func (s *FlashLoanService) DoUniswapFlashLoan(ctx context.Context, amount *big.Int, receiver common.Address, data []byte) error {
	// 调用 Uniswap 闪电贷适配器
	return s.Uniswap.FlashLoan(ctx, amount, receiver, data)
}

// DoSushiSwapFlashLoan 执行 SushiSwap 闪电贷业务
// 参数 ctx: 上下文
// 参数 amount: 闪电贷金额
// 参数 receiver: 闪电贷接收地址
// 参数 data: 附加数据
// 返回值: 错误信息
func (s *FlashLoanService) DoSushiSwapFlashLoan(ctx context.Context, amount *big.Int, receiver common.Address, data []byte) error {
	// 调用 SushiSwap 闪电贷适配器
	return s.SushiSwap.FlashLoan(ctx, amount, receiver, data)
}
