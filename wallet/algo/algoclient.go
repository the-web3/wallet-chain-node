package algo

import (
	"errors"
	"github.com/SavourDao/savour-hd/config"
	"github.com/algorand/go-algorand-sdk/client/v2/algod"
	"github.com/ethereum/go-ethereum/log"
	"net"
	"strings"
)

type algoClient struct {
	Client        *algod.Client
	confirmations uint64
}

func newAlgoClients(conf *config.Config) ([]*algoClient, error) {
	var clients []*algoClient
	for _, rpc := range conf.Fullnode.Algo.RPCs {
		client := &algoClient{
			confirmations: conf.Fullnode.Algo.Confirmations,
		}
		rpcURL := rpc.RPCURL
		domain := strings.TrimPrefix(rpc.RPCURL, "http://")
		domain = strings.TrimPrefix(domain, "https://")
		if strings.Contains(domain, ":") {
			words := strings.Split(domain, ":")
			ipAddr, err := net.ResolveIPAddr("ip", words[0])
			if err != nil {
				log.Error("resolve eth domain failed", "url", rpc.RPCURL)
				continue
			}
			log.Info("ethclient setup client", "ip", ipAddr)
			rpcURL = strings.Replace(rpc.RPCURL, words[0], ipAddr.String(), 1)
		}
		var err error
		client.Client, err = algod.MakeClient(rpcURL, conf.Fullnode.Algo.ApiToken)
		if err != nil {
			log.Error("ethclient dial failed", "err", err)
			continue
		}
		clients = append(clients, client)
	}
	if len(clients) == 0 {
		return nil, errors.New("No clients available")
	}
	return clients, nil
}

func newLocalAlgoClient(network config.NetWorkType) *algoClient {
	return &algoClient{
		Client: &algod.Client{},
	}
}

func (a algoClient) GetLatestBlockHeight() (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (a algoClient) GetAccountBalance(address string) *algod.AccountInformation {
	return a.Client.AccountInformation(address)
}
