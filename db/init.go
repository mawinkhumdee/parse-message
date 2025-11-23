package db

import (
	"parse-message/config"
	"parse-message/db/dbmongo/standardfield"
)

type DB struct {
	Standardfield standardfield.StandardFieldDB
}

func New(config config.DB) DB {
	return DB{
		Standardfield: standardfield.New(config),
	}
}
