package db

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"path/filepath"

	"github.com/pressly/goose/v3"
)

var pragmas = map[string]string{
	"foreign_keys":  "ON",
	"journal_mode":  "WAL",
	"page_size":     "4096",
	"cache_size":    "-8000",
	"synchronous":   "NORMAL",
	"secure_delete": "ON",
	"busy_timeout":  "30000"
}

// connect abre uma conexão com o banco de dados sqlite e executa as migrações
func Connect(ctx context.Context, dataDir string) (*sql.DB, error) {
	if dataDir == "" {
		return nil, fmt.Errorf("data.dir não está definido")
	}

	dbPath := filepath.Join(dataDir, "crush.db")

	db, err := openDB(dbPath)
	if err != nil {
		return nil, err
	}

	if err = db.PingContext(ctx); err != nil {
		db.Close()

		return nil, fmt.Errorf("falha ao conectar-se ao banco de dados: %w", err)
	}

	goose.SetBaseFS(FS)

	if err := goose.SetDialect("sqlite3"); err != nil {
		slog.Error("falha ao definir o dialeto", "error", err)

		return nil, fmt.Errorf("falha ao definir o dialeto: %w", err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		slog.Error("falha ao aplicar migrações", "error", err)

		return nil, fmt.Errorf("falha ao aplicar migrações: %w", err)
	}

	return db, nil
}