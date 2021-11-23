package config

import (
	"github.com/hermeznetwork/hermez-core/aggregator"
	"github.com/hermeznetwork/hermez-core/etherman"
	"github.com/hermeznetwork/hermez-core/jsonrpc"
	"github.com/hermeznetwork/hermez-core/log"
	"github.com/hermeznetwork/hermez-core/sequencer"
)

type Config struct {
	Log        log.Config
	RPC        jsonrpc.Config
	Sequencer  sequencer.Config
	Aggregator aggregator.Config
}

func Load() Config {
	return Config{
		Log: log.Config{
			Level:   "debug",
			Outputs: []string{"stdout"},
		},
		RPC: jsonrpc.Config{
			Host: "",
			Port: 8123,

			ChainID: 2576980377, // 0x99999999
		},
		Sequencer: sequencer.Config{
			Etherman: etherman.Config{},
		},
		Aggregator: aggregator.Config{
			Etherman: etherman.Config{},
		},
	}
}