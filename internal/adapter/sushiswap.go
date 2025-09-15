package adapter

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

// SushiSwapAdapter 负责与 SushiSwap 闪电贷合约进行交互
// client: 以太坊客户端
// address: SushiSwap 闪电贷合约地址
type SushiSwapAdapter struct {
	client  *ethclient.Client
	address common.Address
}

// NewSushiSwapAdapter 创建 SushiSwapAdapter 实例
// 参数 client: 以太坊客户端
// 参数 address: 合约地址字符串
// 返回值: SushiSwapAdapter 实例
func NewSushiSwapAdapter(client *ethclient.Client, address string) *SushiSwapAdapter {
	return &SushiSwapAdapter{
		client:  client,
		address: common.HexToAddress(address),
	}
}

// FlashLoan 调用 SushiSwap 闪电贷合约方法
// 参数 ctx: 上下文
// 参数 amount: 闪电贷金额
// 参数 receiver: 闪电贷接收地址
// 参数 data: 附加数据
// 返回值: 错误信息
func (s *SushiSwapAdapter) FlashLoan(ctx context.Context, amount *big.Int, receiver common.Address, data []byte) error {
	// 此处应调用 SushiSwap 的闪电贷合约方法
	// 伪代码：contract.FlashLoan(opts, receiver, amount, data)
	return nil
}
