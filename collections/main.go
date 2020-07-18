package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/WomenWhoGoTokyo/code-reading-party/collections/handler"

	"github.com/WomenWhoGoTokyo/code-reading-party/collections/collection"

	// MySQL driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/configor"
)

// Config is DB connect setting.
var Config = struct {
	DB struct {
		Host     string `default:"localhost"`
		Name     string
		User     string `default:"root"`
		Password string `required:"true" env:"DBPassword"`
		Port     uint   `default:"3306"`
	}
}{}

func main() {
	configor.Load(&Config, "config.yml")
	db, err := sql.Open("mysql", fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?parseTime=true",
		Config.DB.User,
		Config.DB.Password,
		Config.DB.Host,
		Config.DB.Port,
		Config.DB.Name,
	))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// New Collection
	c := collection.NewCollection(db)
	// New Handlers
	handlers := handler.NewHandlers(c)

	// Set Handlers
	http.HandleFunc("/add", handlers.AddHandler)
	http.HandleFunc("/list", handlers.ListHandler)

	// Stat
	log.Fatal(http.ListenAndServe(":8080", nil))
}
