package config

import (
	"os"
)

// Config 用于存储服务端配置，如 RPC 节点和合约地址
type Config struct {
	RpcUrl      string
	UniswapAddr string
	SushiAddr   string
	TokenA      string // 套利币对A
	TokenB      string // 套利币对B
}

// LoadConfig 加载服务端配置
// 从环境变量读取 RPC 节点、Uniswap/SushiSwap 合约地址
// 返回值: 配置结构体指针
func LoadConfig() *Config {
	return &Config{
		RpcUrl:      os.Getenv("RPC_URL"),        // 以太坊节点地址
		UniswapAddr: os.Getenv("UNISWAP_ADDR"),   // Uniswap 闪电贷合约地址
		SushiAddr:   os.Getenv("SUSHISWAP_ADDR"), // SushiSwap 闪电贷合约地址
		TokenA:      os.Getenv("TOKEN_A"),        // 套利币对A
		TokenB:      os.Getenv("TOKEN_B"),        // 套利币对B
	}
}
