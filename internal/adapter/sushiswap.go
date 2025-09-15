package adapter

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gavin/flashLoanSimple/internal/contract/uniswapv2factory"
	"github.com/gavin/flashLoanSimple/internal/contract/uniswapv2pair"
	"math/big"
	"os"
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
	amount0Out := amount
	amount1Out := big.NewInt(0)
	chainID, err := s.client.NetworkID(ctx)
	if err != nil {
		return fmt.Errorf("get chainID failed: %v", err)
	}
	privKeyHex := os.Getenv("FLASHLOAN_PRIVKEY")
	privKey, err := crypto.HexToECDSA(privKeyHex)
	if err != nil {
		return fmt.Errorf("invalid privkey: %v", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privKey, chainID)
	if err != nil {
		return fmt.Errorf("auth failed: %v", err)
	}
	auth.Context = ctx
	pair, err := uniswapv2pair.NewUniswapv2pair(s.address, s.client)
	if err != nil {
		return fmt.Errorf("new pair instance failed: %v", err)
	}
	tx, err := pair.Swap(auth, amount0Out, amount1Out, receiver, data)
	if err != nil {
		return fmt.Errorf("swap failed: %v", err)
	}
	fmt.Printf("[UniswapV2 FlashLoan] tx sent: %s\n", tx.Hash().Hex())
	return nil
}

// SushiSwapV2Factory 地址（主网）
const SushiSwapV2FactoryAddress = "0xC0AEe478e3658e2610c5F7A4A2E1777cE9e4f2Ac"

// GetPrice 查询 SushiSwap 上 tokenA/tokenB 的价格（主网V2实现）
// 参数 ctx: 上下文
// 参数 tokenA, tokenB: 交易对
// 返回值: 价格（单位：wei），错误信息
func (s *SushiSwapAdapter) GetPrice(ctx context.Context, tokenA, tokenB string) (*big.Int, *big.Int, *big.Int, error) {
	factoryAddr := common.HexToAddress(SushiSwapV2FactoryAddress)
	// 使用 abigen 生成的 UniswapV2Factory binding 查询 pair 地址
	factory, err := uniswapv2factory.NewUniswapv2factory(factoryAddr, s.client)
	if err != nil {
		return nil, nil, nil, err
	}
	tokenAAddr := common.HexToAddress(tokenA)
	tokenBAddr := common.HexToAddress(tokenB)
	pairAddr, err := factory.GetPair(&bind.CallOpts{Context: ctx}, tokenAAddr, tokenBAddr)
	if err != nil {
		return nil, nil, nil, err
	}
	if pairAddr == (common.Address{}) {
		return nil, nil, nil, errors.New("pair not found")
	}
	pair, err := uniswapv2pair.NewUniswapv2pair(pairAddr, s.client)
	if err != nil {
		return nil, nil, nil, err
	}
	reserves, err := pair.GetReserves(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, nil, nil, err
	}
	reserve0 := reserves.Reserve0
	reserve1 := reserves.Reserve1
	if reserve0 == nil || reserve1 == nil {
		return nil, nil, nil, errors.New("reserve0 or reserve1 is nil")
	}
	if tokenAAddr.Hex() < tokenBAddr.Hex() {
		if reserve0.Cmp(big.NewInt(0)) == 0 {
			return nil, nil, nil, errors.New("zero reserve0")
		}
		return new(big.Int).Div(reserve1, reserve0), reserve0, reserve1, nil
	} else {
		if reserve1.Cmp(big.NewInt(0)) == 0 {
			return nil, nil, nil, errors.New("zero reserve1")
		}
		return new(big.Int).Div(reserve0, reserve1), reserve0, reserve1, nil
	}
}
