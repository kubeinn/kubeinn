package DBController

import (
	"context"
	"fmt"
	"os"

	pgxpool "github.com/jackc/pgx/v4/pgxpool"
	// global "github.com/kubeinn/schutterij/internal/global"
)

type PostgresController struct {
	dbName     string
	dbHost     string
	dbPort     int
	dbUser     string
	dbPassword string
	connURL    string
}

func NewPostgresController(dbName string, dbHost string, dbPort int, dbUser string, dbPassword string) *PostgresController {
	pg := PostgresController{}
	pg.dbName = dbName
	pg.dbHost = dbHost
	pg.dbPort = dbPort
	pg.dbUser = dbUser
	pg.dbPassword = dbPassword
	pg.connURL = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", dbName, dbHost, dbPort, dbUser, dbPassword)
	return &pg
}

/*
PILGRIM
*/

func (pg *PostgresController) GetPilgrimPassword(email string) (string, error) {
	dbpool, err := pgxpool.Connect(context.Background(), pg.connURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	var password string
	err = dbpool.QueryRow(context.Background(),
		"SELECT passwd FROM pilgrims WHERE email=$1", email).Scan(&password)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}
	fmt.Println(password)
	return password, err
}

/*
INNKEEPER
*/

func (pg *PostgresController) GetInnkeeperPassword(email string) (string, error) {
	dbpool, err := pgxpool.Connect(context.Background(), pg.connURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	var password string
	err = dbpool.QueryRow(context.Background(),
		"SELECT passwd FROM innkeepers WHERE email=$1", email).Scan(&password)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}
	fmt.Println(password)
	return password, err
}
