package models

import (
	"log"
	"os"
	"time"

	"xorm.io/core"

	"github.com/go-xorm/xorm"
	"github.com/spf13/pflag"
)

var (
	engine *xorm.Engine
	show   = pflag.Bool("sql.debug", false, "show sql logs")
)

func InitDB() {
	var err error
	engine, err = xorm.NewEngine("mysql", os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal(err)
	}

	if err := engine.Ping(); err != nil {
		log.Fatal(err)
	}

	engine.DatabaseTZ = time.UTC
	engine.SetColumnMapper(core.GonicMapper{})
	engine.ShowSQL(*show)
}

func CloseDB() {
	engine.Close()
}
