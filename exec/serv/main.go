package main

import (
	"encoding/json"

	"github.com/KVRes/District/serv"
	"github.com/KevinZonda/GoX/pkg/iox"
)

type Config struct {
	Protocol string `json:"protocol"`
	Addr     string `json:"addr"`
}

var cfg Config

func normalizeConfig() {
	if cfg.Protocol == "" {
		cfg.Protocol = "tcp"
	}
	if cfg.Addr == "" {
		if cfg.Protocol == "tcp" {
			cfg.Addr = "127.0.0.1:9329"
		} else if cfg.Protocol == "unix" {
			cfg.Addr = "/tmp/district.sock"
		}
	}
}

func loadConfig() {
	bs, err := iox.ReadAllByte("config.json")
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(bs, &cfg); err != nil {
		panic(err)
	}
	normalizeConfig()
}

func main() {
	loadConfig()
	svr := serv.NewServer()

	svr.Run(cfg.Protocol, cfg.Addr)
}
