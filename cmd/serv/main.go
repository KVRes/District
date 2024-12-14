package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/KVRes/District/serv"
	"github.com/KevinZonda/GoX/pkg/iox"
)

type Config struct {
	Protocol string  `json:"protocol"`
	Addr     string  `json:"addr"`
	HA       *string `json:"ha"`
}

func (c Config) GetHA() string {
	if c.HA == nil {
		return "data"
	}
	return *c.HA
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
		log.Println("no config file, use default one")
		goto end
	}
	if err := json.Unmarshal(bs, &cfg); err != nil {
		panic(err)
	}
end:
	normalizeConfig()
}

func main() {
	fmt.Println(`===================================================
 ____  _     _        _      _     ___ ____   ____ 
|  _ \(_)___| |_ _ __(_) ___| |_  |_ _|  _ \ / ___|
| | | | / __| __| '__| |/ __| __|  | || |_) | |    
| |_| | \__ \ |_| |  | | (__| |_   | ||  __/| |___ 
|____/|_|___/\__|_|  |_|\___|\__| |___|_|    \____|
===================================================
                   District IPC
              by KevinZonda Research
          https://research.kevinzonda.com                 
===================================================`)
	loadConfig()
	ha := cfg.GetHA()
	err := os.MkdirAll(ha, 0755)
	log.Println("create ha dir", ha)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}
	svr := serv.NewServer()

	log.Println("start server at:", cfg.Protocol, cfg.Addr)

	svr.Run(cfg.Protocol, cfg.Addr)
}
