package config

import (
	"os"
)

type Config struct {
	RpcUrl      string
	UniswapAddr string
	SushiAddr   string
}

func LoadConfig() *Config {
	return &Config{
		RpcUrl:      os.Getenv("RPC_URL"),
		UniswapAddr: os.Getenv("UNISWAP_ADDR"),
		SushiAddr:   os.Getenv("SUSHISWAP_ADDR"),
	}
}
