package main

import (
	"github.com/surrealdb/surrealdb.go"	
	"github.com/RemoteENv-Team/broker/tools"
)

// ConnectionPool is a struct that contains (in memory) the connections to the database.
type ConnectionPool struct {
	connections []*surrealdb.DB
	connSize    int
	connString  string
}

// Stratoshpere is an interface that defines strategies/algorithms for the ConnectionPool and Type Of.
type Stratosphere interface {
	// Other Database, Connections, etc. related functions
	// Base Strategy Has been implemented in db.go
	SetUpPool(size int, connString string) ConnectionPool
	Exec(handler func(conn *surrealdb.DB) error)
}

var Stratoshpere = SetupPool(10, "ws://localhost:8000/rpc")

func SetupPool(size int, connString string) ConnectionPool {

	pool := ConnectionPool{connSize: size, connString: connString}

	for i := 0; i <= size; i++ {
		conn, _ := surrealdb.New(connString)

		pool.connections = append(pool.connections, conn)
	}

	return pool
}

// Gets the last connection from the connections array
// and execute the given handler and releases the connection back to the pool.
// If there are no connections, an temporary connection is being created.
func (c *ConnectionPool) Exec(handler func(conn *surrealdb.DB) error) {
	if len(c.connections) == 0 {
		handleEmptyConnList(c, handler)
	} else {
		conn := c.connections[len(c.connections)-1]

		handlerErr := handler(conn)
		if handlerErr != nil {
			tools.FLogFatal(handlerErr)
		}

		c.connections = append(c.connections, conn)
	}
}

func handleEmptyConnList(c *ConnectionPool, handler func(conn *surrealdb.DB) error) {
	conn, err := surrealdb.New(c.connString)

	if err != nil {
		tools.FLogFatal(err)
	}

	handlerErr := handler(conn)

	if handlerErr != nil {
		tools.FLogFatal(err)
	}

	conn.Close()
}

func main() {
	Stratoshpere.Exec(func(conn *surrealdb.DB) error {
		_, err := conn.Create("test", "test",)
		if err != nil {
			tools.FLogFatal(err)
		}

		return nil
	})
}