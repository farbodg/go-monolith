package server

import "go-monolith/db"

type Config struct {
	DB   db.BaseDAL
	Port uint16
}
