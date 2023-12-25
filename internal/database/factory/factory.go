package factory

import (
	"database/internal/database"
	"database/internal/database/compute"
	"database/internal/database/storage"
	in_mem "database/internal/database/storage/in-mem"
	"database/internal/shared/logger"
)

func CreateDatabase(logger logger.Logger) (*database.Database, error) {
	return database.NewDatabase(
		logger,
		storage.NewEngine(in_mem.NewInMemoryStorage()),
		compute.NewAnalyzer(),
		compute.NewParser(),
	), nil
}
