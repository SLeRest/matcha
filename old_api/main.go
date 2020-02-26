package main

import _"github.com/caarlos0/env"

type Config struct {
	DB_USER				string	`env:"DB_USER" envDefault:"admin"`
	DB_PASSWORD			string	`env:"DB_PASSWORD" envDefault:"secret"`
	DB_HOST				string	`env:"DB_HOST" envDefault:"db-matcha"`
	DB_DBNAME			string	`env:"DB_DBNAME" envDefault:"matcha"`
	DB_PORT				int		`env:"DB_PORT" envDefault:"5046"`
	DB_SSLMODE			string	`env:"DB_SSLMODE" envDefault:"disabled"`
	DB_CONNECT_TIMEOUT	int		`env:"DB_CONNECT_TIMEOUT" envDefault:"5"`
}

func main() {
	config := config.GetConfig()
	app := &app.App{}
	app.Initialize(config)
	app.Run(":80")
}
