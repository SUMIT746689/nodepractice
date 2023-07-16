package pkg

import (
	"log"
	"os"
	"pos/ent"

	"entgo.io/ent/dialect"
	_ "github.com/go-sql-driver/mysql"
)

var client *ent.Client

func InitEnt()  {
	c, err := ent.Open(dialect.MySQL, os.Getenv("DB_URL"))
	if err != nil {
		log.Fatalln(err)
	}

	client = c
}

func EntClient() *ent.Client {
	return client
}
