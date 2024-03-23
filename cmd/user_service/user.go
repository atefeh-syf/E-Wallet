package main

import (
	"github.com/atefeh-syf/yumigo/pkg/user"
	"github.com/atefeh-syf/yumigo/pkg/user/config"
	"github.com/atefeh-syf/yumigo/pkg/user/data/db"
)

// @securityDefinitions.apikey AuthBearer
// @in header
// @name Authorization
func main() {

	cfg := config.GetConfig()
	//logger := logging.NewLogger(cfg)
	

	_ = db.InitDb(cfg)
	defer db.CloseDb()
	// if err != nil {
	// 	logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)
	// }
	//migrations.Up_1()
	user.InitServer(cfg)
}