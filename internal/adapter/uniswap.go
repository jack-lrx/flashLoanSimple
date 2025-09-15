package adapter

import (
	"context"
	"errors"
	"fmt"
	goethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gavin/flashLoanSimple/internal/contract/uniswapv2pair"
	"math/big"
	"os"
	"strings"
)

// UniswapAdapter 负责与 Uniswap 闪电贷合约进行交互
// client: 以太坊客户端
// address: Uniswap 闪电贷合约地址
type UniswapAdapter struct {
	client  *ethclient.Client
	address common.Address
}

// NewUniswapAdapter 创建 UniswapAdapter 实例
// 参数 client: 以太坊客户端
// 参数 address: 合约地址字符串
// 返回值: UniswapAdapter 实例
func NewUniswapAdapter(client *ethclient.Client, address string) *UniswapAdapter {
	return &UniswapAdapter{
		client:  client,
		address: common.HexToAddress(address),
	}
}

// UniswapV2Factory 地址（主网）
const UniswapV2FactoryAddress = "0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"

// FlashLoan 调用 Uniswap V2 Pair 的 swap 方法实现闪电贷
func (u *UniswapAdapter) FlashLoan(ctx context.Context, amount *big.Int, receiver common.Address, data []byte) error {
	amount0Out := amount
	amount1Out := big.NewInt(0)
	chainID, err := u.client.NetworkID(ctx)
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
	pair, err := uniswapv2pair.NewUniswapv2pair(u.address, u.client)
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

// GetPrice 查询 Uniswap 上 tokenA/tokenB 的价格（主网V2实现）
// 参数 ctx: 上下文
// 参数 tokenA, tokenB: 交易对
// 返回值: 价格（单位：wei），错误信息
func (u *UniswapAdapter) GetPrice(ctx context.Context, tokenA, tokenB string) (*big.Int, error) {
	factoryAddr := common.HexToAddress(UniswapV2FactoryAddress)
	// 只保留 getPair 查询部分，后续用合约绑定查询储备
	factoryABI, err := abi.JSON(strings.NewReader(`[{"constant":true,"inputs":[{"internalType":"address","name":"tokenA","type":"address"},{"internalType":"address","name":"tokenB","type":"address"}],"name":"getPair","outputs":[{"internalType":"address","name":"pair","type":"address"}],"payable":false,"stateMutability":"view","type":"function"}]`))
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
	result, err := u.client.CallContract(ctx, msg, nil)
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
	pair, err := uniswapv2pair.NewUniswapv2pair(pairAddr, u.client)
	if err != nil {
		return nil, err
	}
	reserves, err := pair.GetReserves(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}
	reserve0 := reserves.Reserve0
	reserve1 := reserves.Reserve1
	if reserve0 == nil || reserve1 == nil {
		return nil, errors.New("reserve0 or reserve1 is nil")
	}
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
