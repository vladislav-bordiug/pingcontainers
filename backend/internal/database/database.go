package database

import (
	"backend/internal/models"
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"log"
)

type Database interface {
	CreateTableQuery(ctx context.Context) error
	SelectStatusesQuery(ctx context.Context) ([]models.PingStatus, error)
	UpdateStatusQuery(ctx context.Context, ping models.PingStatus) error
}

type DBPool interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, arguments ...interface{}) (pgx.Rows, error)
}

type PGXDatabase struct {
	pool DBPool
}

func NewPGXDatabase(pool DBPool) *PGXDatabase {
	return &PGXDatabase{pool: pool}
}

func (db *PGXDatabase) CreateTableQuery(ctx context.Context) error {
	createTable := `
	CREATE TABLE IF NOT EXISTS pings (
		ip TEXT PRIMARY KEY,
		ping_time DOUBLE PRECISION,
		last_success TIMESTAMP
	)`
	if _, err := db.pool.Exec(context.Background(), createTable); err != nil {
		log.Fatal("Ошибка создания таблицы:", err)
		return err
	}
	return nil
}

func (db *PGXDatabase) SelectStatusesQuery(ctx context.Context) ([]models.PingStatus, error) {
	rows, err := db.pool.Query(ctx, "SELECT ip, ping_time, last_success FROM pings")
	if err != nil {
		log.Println("Ошибка получения статусов пингов")
		return nil, err
	}
	defer rows.Close()

	var statuses []models.PingStatus
	for rows.Next() {
		var ps models.PingStatus
		if err := rows.Scan(&ps.IP, &ps.PingTime, &ps.LastSuccess); err != nil {
			log.Println("Ошибка чтения информации")
			return nil, err
		}
		statuses = append(statuses, ps)
	}
	return statuses, nil
}

func (db *PGXDatabase) UpdateStatusQuery(ctx context.Context, ping models.PingStatus) error {
	upsert := `
	INSERT INTO pings (ip, ping_time, last_success)
	VALUES ($1, $2, $3)
	ON CONFLICT (ip)
	SET 
		ping_time = EXCLUDED.ping_time,
		last_success = CASE 
			WHEN EXCLUDED.last_success IS NOT NULL THEN EXCLUDED.last_success
			ELSE pings.last_success
		END
	`
	var lastSuccess any
	if ping.LastSuccess.IsZero() {
		lastSuccess = nil
	} else {
		lastSuccess = ping.LastSuccess
	}
	if _, err := db.pool.Exec(ctx, upsert, ping.IP, ping.PingTime, lastSuccess); err != nil {
		log.Printf("Ошибка базы: %v", err)
		return err
	}
	return nil
}
