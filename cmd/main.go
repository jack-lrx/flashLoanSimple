package main

import (
	"context"
	"github.com/gavin/flashLoanSimple/config"
	"github.com/gavin/flashLoanSimple/internal/adapter"
	"github.com/gavin/flashLoanSimple/internal/arbitrage"
	"github.com/gavin/flashLoanSimple/pkg"
	"log"
	"math/big"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg := config.LoadConfig()
	client, err := pkg.NewEthClient(cfg.RpcUrl)
	if err != nil {
		log.Fatalf("failed to connect to eth node: %v", err)
	}
	uniswap := adapter.NewUniswapAdapter(client, cfg.UniswapAddr)
	sushiswap := adapter.NewSushiSwapAdapter(client, cfg.SushiAddr)

	// 初始化套利服务
	minProfit := big.NewInt(1e15) // 0.001 ETH，实际可配置
	arb := arbitrage.NewArbitrageService(uniswap, sushiswap, minProfit)
	tokenA := cfg.TokenA
	tokenB := cfg.TokenB

	// 启动定时套利任务
	go func() {
		ticker := time.NewTicker(30 * time.Second) // 每 30 秒轮询
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				ctx := context.Background()
				err := arb.MonitorAndArbitrage(ctx, tokenA, tokenB)
				if err != nil {
					log.Printf("[Arbitrage] error: %v", err)
				}
			}
		}
	}()

	// 优雅退出
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	log.Println("shutting down...")
	os.Exit(0)
}
