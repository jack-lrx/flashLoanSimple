package pkg

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

// NewEthClient 初始化以太坊客户端
// 参数 rpcUrl: 以太坊节点地址
// 返回值: 以太坊客户端实例和错误信息
func NewEthClient(rpcUrl string) (*ethclient.Client, error) {
	// 连接以太坊节点，返回客户端实例
	return ethclient.Dial(rpcUrl)
}
