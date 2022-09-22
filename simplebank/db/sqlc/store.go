package db

import (
	"database/sql"
)

//store provides all functions to excecute db qureies and transactions
type store struct {
	*Queries
	db *sql.DB
}
