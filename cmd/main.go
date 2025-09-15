package main

import (
	"github.com/gavin/flashLoanSimple/api"
	"github.com/gavin/flashLoanSimple/config"
	"github.com/gavin/flashLoanSimple/internal/adapter"
	"github.com/gavin/flashLoanSimple/internal/service"
	"github.com/gavin/flashLoanSimple/pkg"
	"log"
	"net/http"
	"os"
)

func main() {
	cfg := config.LoadConfig()
	client, err := pkg.NewEthClient(cfg.RpcUrl)
	if err != nil {
		log.Fatalf("failed to connect to eth node: %v", err)
	}
	uniswap := adapter.NewUniswapAdapter(client, cfg.UniswapAddr)
	sushiswap := adapter.NewSushiSwapAdapter(client, cfg.SushiAddr)
	svc := service.NewFlashLoanService(uniswap, sushiswap)
	handler := api.NewHandler(svc)

	http.HandleFunc("/flashloan", handler.FlashLoan)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("server listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
