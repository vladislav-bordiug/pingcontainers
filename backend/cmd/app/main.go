package main

import (
	"backend/internal/app"
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
	"time"
)

func Config() *pgxpool.Config {
	// const defaultMaxConns = int32(4)
	// const defaultMinConns = int32(0)
	const defaultMaxConnLifetime = time.Hour
	const defaultMaxConnIdleTime = time.Minute * 30
	const defaultHealthCheckPeriod = time.Minute
	const defaultConnectTimeout = time.Second * 5
	dbConfig, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Не получилось создать конфиг для бд, ошибка: ", err)
	}

	// dbConfig.MaxConns = defaultMaxConns
	// dbConfig.MinConns = defaultMinConns
	dbConfig.MaxConnLifetime = defaultMaxConnLifetime
	dbConfig.MaxConnIdleTime = defaultMaxConnIdleTime
	dbConfig.HealthCheckPeriod = defaultHealthCheckPeriod
	dbConfig.ConnConfig.ConnectTimeout = defaultConnectTimeout

	dbConfig.BeforeAcquire = func(ctx context.Context, c *pgx.Conn) bool {
		log.Println("До получения соединения к бд!")
		return true
	}

	dbConfig.AfterRelease = func(c *pgx.Conn) bool {
		log.Println("После получения соединения к бд!")
		return true
	}

	dbConfig.BeforeClose = func(c *pgx.Conn) {
		log.Println("Закрыто соединение к бд!")
	}

	return dbConfig
}

func main() {
	var err error
	connPool, err := pgxpool.NewWithConfig(context.Background(), Config())
	if err != nil {
		log.Fatal("Ошибка создания соединения к бд!", err)
	}
	defer connPool.Close()
	connection, err := connPool.Acquire(context.Background())
	if err != nil {
		log.Fatal("Ошибка получения соединения к бд!", err)
	}
	defer connection.Release()
	err = connection.Ping(context.Background())
	if err != nil {
		log.Fatal("Не получилось пропинговать бд", err)
	}
	application := app.NewApp(connPool, os.Getenv("PORT"))
	log.Fatal(application.Run())
}
