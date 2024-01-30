package main

func main() {
	cfg := config.GetConfig()
	// err := db.InitDB(cfg)
	// if err != nil {

	// }
	_ = db.InitDb(cfg)
	api.InitServer(cfg)
}
