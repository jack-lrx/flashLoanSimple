package adapter

import (
	"context"
	"errors"
	goethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
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

	return nil
}

// SushiSwapV2PairABI 与 Uniswap V2 相同
const SushiSwapV2PairABI = `[{"constant":true,"inputs":[],"name":"getReserves","outputs":[{"internalType":"uint112","name":"reserve0","type":"uint112"},{"internalType":"uint112","name":"reserve1","type":"uint112"},{"internalType":"uint32","name":"blockTimestampLast","type":"uint32"}],"payable":false,"stateMutability":"view","type":"function"}]`

// SushiSwapV2Factory 地址（主网）
const SushiSwapV2FactoryAddress = "0xC0AEe478e3658e2610c5F7A4A2E1777cE9e4f2Ac"

// SushiSwapV2Factory ABI（只用 getPair 方法）
const SushiSwapV2FactoryABI = `[{"constant":true,"inputs":[{"internalType":"address","name":"tokenA","type":"address"},{"internalType":"address","name":"tokenB","type":"address"}],"name":"getPair","outputs":[{"internalType":"address","name":"pair","type":"address"}],"payable":false,"stateMutability":"view","type":"function"}]`

// GetPrice 查询 SushiSwap 上 tokenA/tokenB 的价格（主网V2实现）
// 参数 ctx: 上下文
// 参数 tokenA, tokenB: 交易对
// 返回值: 价格（单位：wei），错误信息
func (s *SushiSwapAdapter) GetPrice(ctx context.Context, tokenA, tokenB string) (*big.Int, error) {
	factoryAddr := common.HexToAddress(SushiSwapV2FactoryAddress)
	factoryABI, err := abi.JSON(strings.NewReader(SushiSwapV2FactoryABI))
	if err != nil {
		return nil, err
	}
	tokenAAddr := common.HexToAddress(tokenA)
	tokenBAddr := common.HexToAddress(tokenB)
	input, err := factoryABI.Pack("getPair", tokenAAddr, tokenBAddr)
	if err != nil {
		return nil, err
	}
	msg := goethereum.CallMsg{
		To:   &factoryAddr,
		Data: input,
	}
	result, err := s.client.CallContract(ctx, msg, nil)
	if err != nil {
		return nil, err
	}
	outputs, err := factoryABI.Unpack("getPair", result)
	if err != nil || len(outputs) == 0 {
		return nil, errors.New("getPair unpack failed")
	}
	pairAddr := outputs[0].(common.Address)
	if pairAddr == (common.Address{}) {
		return nil, errors.New("pair not found")
	}
	pairABI, err := abi.JSON(strings.NewReader(SushiSwapV2PairABI))
	if err != nil {
		return nil, err
	}
	input2, err := pairABI.Pack("getReserves")
	if err != nil {
		return nil, err
	}
	msg2 := goethereum.CallMsg{
		To:   &pairAddr,
		Data: input2,
	}
	result2, err := s.client.CallContract(ctx, msg2, nil)
	if err != nil {
		return nil, err
	}
	outputs2, err := pairABI.Unpack("getReserves", result2)
	if err != nil || len(outputs2) < 2 {
		return nil, errors.New("getReserves unpack failed")
	}
	reserve0 := outputs2[0].(*big.Int)
	reserve1 := outputs2[1].(*big.Int)
	// SushiSwap V2: pair 的 token0/token1 按地址排序
	// reserve0 对应 token0，reserve1 对应 token1
	if tokenAAddr.Hex() < tokenBAddr.Hex() {
		if reserve0.Cmp(big.NewInt(0)) == 0 {
			return nil, errors.New("zero reserve0")
		}
		return new(big.Int).Div(reserve1, reserve0), nil
	} else {
		if reserve1.Cmp(big.NewInt(0)) == 0 {
			return nil, errors.New("zero reserve1")
		}
		return new(big.Int).Div(reserve0, reserve1), nil
	}
}
