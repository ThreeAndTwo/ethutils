package ethutils

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/metachris/flashbotsrpc"
)

type MEVOptions struct {
	Client     *flashbotsrpc.FlashbotsRPC
	PrivateKey *ecdsa.PrivateKey
}

func NewMEV(rpc string, privateKey *ecdsa.PrivateKey) *MEVOptions {
	return &MEVOptions{Client: flashbotsrpc.New(rpc), PrivateKey: privateKey}
}

func (mev *MEVOptions) GetUserStats(blockNumber uint64) (flashbotsrpc.FlashbotsUserStats, error) {
	result, err := mev.Client.FlashbotsGetUserStats(mev.PrivateKey, blockNumber)
	return result, err
}

func (mev *MEVOptions)SendBundle(txs []string, blockNumber uint64) (flashbotsrpc.FlashbotsSendBundleResponse, error) {
	sendBundleArgs := flashbotsrpc.FlashbotsSendBundleRequest{
		Txs:         txs,
		BlockNumber: fmt.Sprintf("0x%x", blockNumber),
	}
	return mev.Client.FlashbotsSendBundle(mev.PrivateKey, sendBundleArgs)
}


