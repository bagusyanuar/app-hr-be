package bootstrap

import (
	"fmt"

	"github.com/bagusyanuar/app-hr-be/internal/config"
)

func initialize() *config.AppConfig {
	viper := config.NewViper()
	app := config.NewFiber(viper)
	logger := config.NewLogger(viper)
	defer logger.Sync()
	cfgDB := config.NewDatabaseConfig(viper)
	db := config.NewDatabaseConnection(cfgDB)

	return &config.AppConfig{
		App:    app,
		Viper:  viper,
		DB:     db,
		Logger: logger,
	}
}

func Start() {
	cfg := initialize()

	envPort := cfg.Viper.GetString("APP_PORT")
	port := fmt.Sprintf(":%s", envPort)
	server := cfg.App
	fmt.Println("Fiber server running on", port)
	if err := server.Listen(port); err != nil {
		panic(err)
	}
}
