package pkg

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

func NewEthClient(rpcUrl string) (*ethclient.Client, error) {
	return ethclient.Dial(rpcUrl)
}
