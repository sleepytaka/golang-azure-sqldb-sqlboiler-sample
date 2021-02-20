package main

import (
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mssql/driver"
	"golang-azure-sqldb-sqlboiler/config"
	"golang-azure-sqldb-sqlboiler/di"
	"log"
)

func main() {
	fmt.Println("Start...")
	c, err := config.New(); if err != nil {
		panic(err)
	}

	conn  := driver.MSSQLBuildQueryString(
		c.Database.User,
		c.Database.Password,
		c.Database.Dbname,
		c.Database.Host,
		c.Database.Port,
		c.Database.SslMode)

	h, cleanup, err := di.InitializeHandler(conn)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer cleanup()

	err = h.Run(); if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println("Exit...")
}
