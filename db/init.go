package db

import (
	"parse-message/config"
	"parse-message/db/dbmongo/message"
	"parse-message/db/dbmongo/parseresult"
	"parse-message/db/dbmongo/standardfield"
)

type DB struct {
	Standardfield standardfield.StandardFieldDB
	ParseResult   parseresult.ParseResultDB
	Message       message.MessageDB
}

func New(config config.DB) DB {
	return DB{
		Standardfield: standardfield.New(config),
		ParseResult:   parseresult.New(config),
		Message:       message.New(config),
	}
}
