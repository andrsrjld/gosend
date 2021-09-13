package gosend

import "database/sql"

type GoSendRepository struct {
	DB *sql.DB
}

func NewGoSendRepository(DB *sql.DB) *GoSendRepository {
	return &GoSendRepository{
		DB: DB,
	}
}
