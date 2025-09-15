package arbitrage

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gavin/flashLoanSimple/internal/adapter"
	"math/big"
	"os"
	"strings"
	"time"
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
	uniPrice, uniR0, uniR1, err := a.Uniswap.GetPrice(ctx, tokenA, tokenB)
	if err != nil {
		return err
	}
	sushiPrice, sushiR0, sushiR1, err := a.SushiSwap.GetPrice(ctx, tokenA, tokenB)
	if err != nil {
		return err
	}
	// 判断是否存在套利机会
	profit := new(big.Int).Sub(uniPrice, sushiPrice)
	if profit.Cmp(a.MinProfit) > 0 {
		// Uniswap 价格高，买入 SushiSwap，卖出 Uniswap
		return a.ExecuteArbitrage(ctx, tokenA, tokenB, "sushiswap", "uniswap", profit, sushiR0, sushiR1)
	}
	profit = new(big.Int).Sub(sushiPrice, uniPrice)
	if profit.Cmp(a.MinProfit) > 0 {
		// SushiSwap 价格高，买入 Uniswap，卖出 SushiSwap
		return a.ExecuteArbitrage(ctx, tokenA, tokenB, "uniswap", "sushiswap", profit, uniR0, uniR1)
	}
	// 没有套利机会
	return nil
}

// ExecuteArbitrage 执行套利交易
// buyDEX, sellDEX: 买入和卖出 DEX
// profit: 预期利润
func (a *ArbitrageService) ExecuteArbitrage(ctx context.Context, tokenA, tokenB, buyDEX, sellDEX string,
	profit, reserve0, reserve1 *big.Int) error {
	// 1. 动态计算最佳借款金额（池子储备的 5%）
	var amountIn *big.Int
	var err error
	// 取较小储备的 5% 作为借款金额
	reserveMin := reserve0
	if reserve1.Cmp(reserveMin) < 0 {
		reserveMin = reserve1
	}
	amountIn = new(big.Int).Div(reserveMin, big.NewInt(20)) // 5%
	if amountIn.Cmp(big.NewInt(1e15)) < 0 {
		amountIn = big.NewInt(1e15) // 最小 0.001 ETH
	}
	if amountIn.Cmp(profit) > 0 {
		amountIn = profit // 不超过预期利润
	}

	// 2. receiver 地址通过环境变量读取
	receiverHex := os.Getenv("ARBITRAGE_RECEIVER")
	if receiverHex == "" {
		return fmt.Errorf("ARBITRAGE_RECEIVER env not set")
	}
	receiver := common.HexToAddress(receiverHex)

	// 3. data 字段 ABI 编码套利参数
	arbAbi, err := abi.JSON(strings.NewReader(`[{"name":"arbitrageCallback","type":"function","inputs":[{"name":"tokenA","type":"address"},{"name":"tokenB","type":"address"},{"name":"amountIn","type":"uint256"},{"name":"buyDEX","type":"string"},{"name":"sellDEX","type":"string"}] }]`))
	if err != nil {
		return fmt.Errorf("abi parse failed: %v", err)
	}
	data, err := arbAbi.Pack("arbitrageCallback", common.HexToAddress(tokenA), common.HexToAddress(tokenB), amountIn, buyDEX, sellDEX)
	if err != nil {
		return fmt.Errorf("abi pack failed: %v", err)
	}

	fmt.Printf("[Arbitrage] FlashLoan %s/%s on %s, amount: %s\n", tokenA, tokenB, buyDEX, amountIn.String())
	var tx *types.Transaction
	if buyDEX == "uniswap" {
		err = a.Uniswap.FlashLoan(ctx, amountIn, receiver, data)
	} else if buyDEX == "sushiswap" {
		err = a.SushiSwap.FlashLoan(ctx, amountIn, receiver, data)
	}
	if err != nil {
		return fmt.Errorf("flashloan failed: %v", err)
	}
	fmt.Printf("[Arbitrage] FlashLoan tx sent: %s\n", tx.Hash().Hex())

	// 4. 等待交易回执
	client := a.Uniswap.Client() // 任选一个 client
	for i := 0; i < 30; i++ {
		receipt, err := client.TransactionReceipt(ctx, tx.Hash())
		if err == nil && receipt != nil {
			if receipt.Status == 1 {
				fmt.Printf("[Arbitrage] FlashLoan success, amount: %s\n", amountIn.String())
				fmt.Printf("[Arbitrage] Profit: %s wei\n", profit.String())
				return nil
			} else {
				return fmt.Errorf("tx failed, status: %d", receipt.Status)
			}
		}
		time.Sleep(2 * time.Second)
	}
	return fmt.Errorf("tx receipt timeout: %s", tx.Hash().Hex())
}
