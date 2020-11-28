package dbcontroller

import (
	"context"
	"fmt"
	"os"

	pgxpool "github.com/jackc/pgx/v4/pgxpool"
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
	pg.connURL = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", dbName, dbPassword, dbHost, dbPort, dbUser)
	return &pg
}

/*
PILGRIM
*/

func (pg *PostgresController) InsertPilgrim(organization string, description string, username string, email string, password string) error {
	dbpool, err := pgxpool.Connect(context.Background(), pg.connURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	tx, err := dbpool.Begin(context.Background())
	if err != nil {
		return err
	}
	// Rollback is safe to call even if the tx is already closed, so if
	// the tx commits successfully, this is a no-op
	defer tx.Rollback(context.Background())

	_, err = tx.Exec(context.Background(), "INSERT INTO api.pilgrims (organization, description, username, email, passwd) VALUES ($1, $2, $3, $4, $5)", organization, description, username, email, password)
	if err != nil {
		return err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (pg *PostgresController) UpdatePilgrimPassword(id string, password string) error {
	dbpool, err := pgxpool.Connect(context.Background(), pg.connURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	tx, err := dbpool.Begin(context.Background())
	if err != nil {
		return err
	}
	// Rollback is safe to call even if the tx is already closed, so if
	// the tx commits successfully, this is a no-op
	defer tx.Rollback(context.Background())

	_, err = tx.Exec(context.Background(), "UPDATE api.pilgrims SET passwd = $2 WHERE id = $1", id, password)
	if err != nil {
		return err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}
	return nil
}

/*
REEVE
*/

func (pg *PostgresController) SelectPilgrimByUsername(username string) (string, string, string, error) {
	dbpool, err := pgxpool.Connect(context.Background(), pg.connURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	var id string
	var password string
	var status string

	err = dbpool.QueryRow(context.Background(),
		"SELECT id,passwd,status FROM api.pilgrims WHERE username=$1", username).Scan(&id, &password, &status)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}
	return id, password, status, err
}

/*
INNKEEPER
*/

func (pg *PostgresController) SelectInnkeeperByUsername(username string) (string, string, error) {
	dbpool, err := pgxpool.Connect(context.Background(), pg.connURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	var id string
	var password string
	err = dbpool.QueryRow(context.Background(),
		"SELECT id,passwd FROM api.innkeepers WHERE username=$1", username).Scan(&id, &password)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}
	return id, password, err
}

func (pg *PostgresController) InsertInnkeeper(username string, email string, password string) error {
	dbpool, err := pgxpool.Connect(context.Background(), pg.connURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	tx, err := dbpool.Begin(context.Background())
	if err != nil {
		return err
	}
	// Rollback is safe to call even if the tx is already closed, so if
	// the tx commits successfully, this is a no-op
	defer tx.Rollback(context.Background())

	_, err = tx.Exec(context.Background(), "INSERT INTO api.innkeepers (username, email, passwd) VALUES ($1, $2, $3)", username, email, password)
	if err != nil {
		return err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}
	return nil
}

/*
VILLAGE
*/
func (pg *PostgresController) SelectVillageByOrganization(organization string) (string, error) {
	dbpool, err := pgxpool.Connect(context.Background(), pg.connURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	var id string
	err = dbpool.QueryRow(context.Background(),
		"SELECT id FROM api.villages WHERE organization=$1", organization).Scan(&id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}
	return id, err
}

func (pg *PostgresController) InsertVillage(organization string, description string) error {
	dbpool, err := pgxpool.Connect(context.Background(), pg.connURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	tx, err := dbpool.Begin(context.Background())
	if err != nil {
		return err
	}
	// Rollback is safe to call even if the tx is already closed, so if
	// the tx commits successfully, this is a no-op
	defer tx.Rollback(context.Background())

	_, err = tx.Exec(context.Background(), "INSERT INTO api.villages (organization, description) VALUES ($1, $2)", organization, description)
	if err != nil {
		return err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}
	return nil
}

/*
PROJECTS
*/
// SelectProjectById is ...
func (pg *PostgresController) SelectProjectById(id string) (string, error) {
	dbpool, err := pgxpool.Connect(context.Background(), pg.connURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	var title string
	err = dbpool.QueryRow(context.Background(),
		"SELECT title FROM api.projects WHERE id=$1", id).Scan(&title)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}
	return title, err
}
