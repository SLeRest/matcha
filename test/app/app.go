package app

import (
	"fmt"
	"os"
	"net/http"
	"database/sql"
	_"github.com/lib/pq"
)

type App struct {
	Router *mux.ROUTER
	DB		driver.Conn
}

func (a *App) Initialize(config *config.Config) {
	connStr := fmt.Sprintf(`
				user=%s
				password=%s
				host=%s
				dbname=%s
				port=%d
				sslmode=%s
				connect_timeout= %d`,
				config.DB_USER,
				config.DB_PASSWORD,
				config.DB_HOST,
				config.DB_DBNAME,
				config.DB_PORT,
				config.DB_SSLMODE,
				config.DB_CONNECT_TIMEOUT)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.fatal("Error initialize connection to database")
		os.exit(1)
	}
	a.Router = mux.NewRouter().StrictSlash(true)
	a.setRouters()
}

func (a *App) SetRouters() {
	a.Router.HandleFunc("/matcha/v1/user", getUsers).Methods("GET")
	//a.Router.HandleFunc("/matcha/v1/user", postUser).Methods("POST")
	//a.Router.HandleFunc("/matcha/v1/user/{id}", getUser).Methods("GET")
	//a.Router.HandleFunc("/matcha/v1/user/{id}", putUser).Methods("PUT")
	//a.Router.HandleFunc("/matcha/v1/user/{id}", deleteUser).Methods("DELETE")

	/*	DEBUG */
	a.Router.HandleFunc("/matcha/v1/config", getConfig).Methods("GET")
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
