package arbitrage

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gavin/flashLoanSimple/internal/adapter"
	"math/big"
)

// ArbitrageService 套利服务，负责监控价格并执行套利
// uniswap/sushiswap: DEX 适配器
// minProfit: 最小套利利润（单位：wei）
type ArbitrageService struct {
	Uniswap   *adapter.UniswapAdapter
	SushiSwap *adapter.SushiSwapAdapter
	MinProfit *big.Int
}

// NewArbitrageService 创建套利服务
func NewArbitrageService(uniswap *adapter.UniswapAdapter, sushi *adapter.SushiSwapAdapter, minProfit *big.Int) *ArbitrageService {
	return &ArbitrageService{
		Uniswap:   uniswap,
		SushiSwap: sushi,
		MinProfit: minProfit,
	}
}

// MonitorAndArbitrage 监控价格并自动执行套利
// tokenA, tokenB: 交易对
// ctx: 上下文
func (a *ArbitrageService) MonitorAndArbitrage(ctx context.Context, tokenA, tokenB string) error {
	// 查询 Uniswap 和 SushiSwap 的价格
	uniPrice, err := a.Uniswap.GetPrice(ctx, tokenA, tokenB)
	if err != nil {
		return err
	}
	sushiPrice, err := a.SushiSwap.GetPrice(ctx, tokenA, tokenB)
	if err != nil {
		return err
	}
	// 判断是否存在套利机会
	profit := new(big.Int).Sub(uniPrice, sushiPrice)
	if profit.Cmp(a.MinProfit) > 0 {
		// Uniswap 价格高，买入 SushiSwap，卖出 Uniswap
		return a.ExecuteArbitrage(ctx, tokenA, tokenB, "sushiswap", "uniswap", profit)
	}
	profit = new(big.Int).Sub(sushiPrice, uniPrice)
	if profit.Cmp(a.MinProfit) > 0 {
		// SushiSwap 价格高，买入 Uniswap，卖出 SushiSwap
		return a.ExecuteArbitrage(ctx, tokenA, tokenB, "uniswap", "sushiswap", profit)
	}
	// 没有套利机会
	return nil
}

// ExecuteArbitrage 执行套利交易
// buyDEX, sellDEX: 买入和卖出 DEX
// profit: 预期利润
func (a *ArbitrageService) ExecuteArbitrage(ctx context.Context, tokenA, tokenB, buyDEX, sellDEX string, profit *big.Int) error {
	var err error
	amountIn := profit                                        // 可根据实际业务调整
	receiver := common.HexToAddress("0x_your_wallet_address") // TODO: 替换为实际钱包地址
	// data 可用于合约回调套利逻辑，生产环境应编码具体套利操作
	data := []byte("arbitrage_callback_data")

	fmt.Printf("[Arbitrage] FlashLoan %s/%s on %s, amount: %s\n", tokenA, tokenB, buyDEX, amountIn.String())
	if buyDEX == "uniswap" {
		err = a.Uniswap.FlashLoan(ctx, amountIn, receiver, data)
	} else if buyDEX == "sushiswap" {
		err = a.SushiSwap.FlashLoan(ctx, amountIn, receiver, data)
	} else {
		return fmt.Errorf("unsupported buyDEX: %s", buyDEX)
	}
	if err != nil {
		return fmt.Errorf("flashloan failed: %v", err)
	}
	fmt.Printf("[Arbitrage] FlashLoan success, amount: %s\n", amountIn.String())
	fmt.Printf("[Arbitrage] Profit: %s wei\n", profit.String())
	return nil
}
