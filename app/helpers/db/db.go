package db

import (
	"devcodeIdentity/app/config"
	"fmt"
	"github.com/go-pg/pg"
	"time"
)

func NewDb() *pg.DB {
	fmt.Println("[DB] Connection...", config.ConnDB.ConnPath())
	opt, _ := pg.ParseURL(config.ConnDB.ConnPath())

	opt.PoolSize = 15

	db := pg.Connect(opt).WithTimeout(time.Second * 40)

	return db
}
