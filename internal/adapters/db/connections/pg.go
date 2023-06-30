package connections

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresManager struct {
	db *sqlx.DB
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func NewPostgresManager() *PostgresManager {
	return &PostgresManager{}
}

func (p *PostgresManager) Connect(ctx context.Context, config PostgresConfig) error {
	db, err := sqlx.Connect("postgres", p.buildConnectionString(config))
	if err != nil {
		return err
	}

	err = db.PingContext(ctx)
	if err != nil {
		return err
	}

	p.db = db
	return nil
}

func (p *PostgresManager) buildConnectionString(config PostgresConfig) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.Database)
}

func (p *PostgresManager) GetConnection() *sqlx.DB {
	return p.db
}

func (p *PostgresManager) CloseConnection() error {
	return p.db.Close()
}
