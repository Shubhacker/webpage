package postgres

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

var pool *pgxpool.Pool

func InitDbPool() {
	// host := os.Getenv("DB_HOST")
	// port := os.Getenv("DB_PORT")
	// user := os.Getenv("DB_USER")
	// pass := os.Getenv("DB_PASSWORD")
	// sslmode := os.Getenv("SSLMODE")

	databaseUrl := " host=" + "localhost" + " port=" + "8888" + " user=" + "postgres" + " password=" + "RogStrix@1080" + " dbname=" + "postgres" + " sslmode=" + "disable"

	config, err := pgxpool.ParseConfig(databaseUrl)
	if err != nil {
		log.Print(err)
		log.Print("Error in config.")
		//return &pgxpool.Pool{}
	}
	pool, err = pgxpool.ConnectConfig(context.Background(), config)

	if err != nil {
		log.Print(err)
		log.Print("Could not connect to Postgres.")
	}
	log.Println("Postgres connected!")
}

func GetPool() *pgxpool.Pool {
	if pool == nil {
		InitDbPool()
	}

	connectedPoolSize := pool.AcquireAllIdle(context.Background())
	for connectedPoolSize == nil {
		log.Println("Pg Connection Lost")
		pool.Close()
		time.Sleep(2 * time.Second)
		log.Print("Reconnecting...")
		InitDbPool()
		connectedPoolSize = pool.AcquireAllIdle(context.Background())
	}
	return pool
}
