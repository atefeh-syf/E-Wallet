package main

import (
	"github.com/atefeh-syf/E-Wallet/api"
	"github.com/atefeh-syf/E-Wallet/config"
	"github.com/atefeh-syf/E-Wallet/data/db"
)

func main() {
	cfg := config.GetConfig()
	// err := db.InitDB(cfg)
	// if err != nil {

	// }
	_ = db.InitDb(cfg)
	api.InitServer(cfg)
}
