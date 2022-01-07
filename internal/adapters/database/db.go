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
	pool DbPool
}

type DBConfig struct {
	Username, Hostname, Database, CACert, ClientCert, SSLKey, Password string
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
	db.pool, err = pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		l.Error("error connecting to the database: ", zap.Error(err))
		return err
	}
	return nil
}

func (d *DBImpl) Close() {
	d.pool.Close()
}

func (d *DBImpl) Query(sql string, params ...interface{}) (pgx.Rows, error) {
	return d.pool.Query(context.Background(), sql, params...)
}

// QueryRow - to fetch at most one row.
func (d *DBImpl) QueryRow(sql string, params ...interface{}) pgx.Row {
	row := d.pool.QueryRow(context.Background(), sql, params...)
	return row
}

func (d *DBImpl) Exec(sql string, params ...interface{}) (int64, error) {
	l := zap.L()
	exec, err := d.pool.Exec(context.Background(), sql, params...)

	if err != nil {
		l.Error("error when executing sql command", zap.Error(err))
		return 0, err
	}
	return exec.RowsAffected(), nil
}

func NewPostgresDatabase() *DBImpl {
	db := &DBImpl{}

	dbCfg := &DBConfig{
		Username:   "postgres",
		Hostname:   "0.0.0.0",
		Database:   "postgres",
		Port:       5432,
		PoolSize:   4,
		Insecure:   true,
		Password: "postgres",
	}

	db.CreatePool(dbCfg)
	return db
}
