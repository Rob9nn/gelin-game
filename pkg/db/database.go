package db

import (
	"database/sql"
	"errors"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

type ConnectionPool struct {
	pool               *sql.DB
	maxConnections     int
	currentConnections int
	mutex              *sync.Mutex
}

var pool *ConnectionPool

func NewConnectionPool(maxConnections int) error {
	log.Printf("Initializing connection pool")
	if maxConnections == 0 {
		return errors.New("Max connections cannot be 0")
	}

	connStr := "postgresql://pqgotest:password@localhost/pqgotest?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(maxConnections)
	log.Printf("Connection pool init with max size of %d\n", maxConnections)

	pool = &ConnectionPool{
		pool:               db,
		maxConnections:     maxConnections,
		currentConnections: 0,
		mutex:              &sync.Mutex{},
	}
	return nil
}

func GetConnection() (*sql.DB, error) {
	if pool == nil {
		log.Fatal("pool has not been initilazed")
	}
	pool.mutex.Lock()
	defer pool.mutex.Unlock()

	if pool.currentConnections == pool.maxConnections {
		return nil, errors.New("Maximum of connections reached")
	}

	pool.currentConnections++

	return pool.pool, nil
}

func CloseConnection() {
	pool.mutex.Lock()
	defer pool.mutex.Unlock()

	pool.currentConnections--
}
