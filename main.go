package main

import (
	"log"

	"git.tulz.ir/monitoring/server/models"
	monitor "git.tulz.ir/monitoring/server/services/monitor"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/pflag"
)

func main() {
	pflag.Parse()
	models.InitDB()

	log.Fatal(monitor.Start())
	models.CloseDB()
}
