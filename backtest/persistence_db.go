package backtest

import (
	"gorm.io/gorm"
)

var persistenceDB *gorm.DB

// UseDatabase enables database-backed persistence for all backtest storage operations.
func UseDatabase(db *gorm.DB) {
	persistenceDB = db
}

func usingDB() bool {
	return persistenceDB != nil
}
