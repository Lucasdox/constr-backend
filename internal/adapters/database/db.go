package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type DBImpl struct {
	dbPool DbPool
}

type DBConfig struct {
	Username, Hostname, Database, CACert, ClientCert, SSLKey string
	Port, PoolSize               int
	Insecure                     bool
}

type DbPool interface {
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
	Close()
}

func (db *DBImpl) CreatePool(cfg *DBConfig) error {
	l := zap.L()
	var connString string
	if cfg.Insecure {
		connString = fmt.Sprintf("postgresql://%s:@%s:%d/%s?sslmode=disable", cfg.Username, cfg.Hostname, cfg.Port, cfg.Database)
	} else {
		connString = fmt.Sprintf(
			"postgresql://%s:@%s:%d/%s?sslmode=verify-full&sslrootcert=%s&sslcert=%s&sslkey=%s&pool_max_conns=%d",
			cfg.Username, cfg.Hostname, cfg.Port, cfg.Database, cfg.CACert, cfg.ClientCert, cfg.SSLKey, cfg.PoolSize,
		)
	}

	// Set connection pool configuration, with maximum connection pool size.
	config, err := pgxpool.ParseConfig(connString)

	if err != nil {
		l.Error("error configuring the database: ", zap.Error(err))
		return err
	}

	// Create a connection pool to the "bnko" database.
	db.dbPool, err = pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		l.Error("error connecting to the database: ", zap.Error(err))
		return err
	}
	return nil
}
