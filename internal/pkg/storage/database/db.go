package database

import "database/sql"

type DatabaseStorageAdapter struct {
	dbObj *sql.DB
}
