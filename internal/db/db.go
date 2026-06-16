package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func New(connString string) *pgxpool.Pool {
	pool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		log.Fatal("db connection error:", err)
	}
	if err := pool.Ping(context.Background()); err != nil{
		log.Fatal("db ping error:", err)
	}
	log.Println("✅ PostgreSQL connected!")
	return pool
}
