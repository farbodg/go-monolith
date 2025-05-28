package data

import "go-monolith/db"

func New(baseProvider db.BaseDAL) DALProvider {
	return &dataProvider{
		BaseDAL: baseProvider,
	}
}
