package dbcontroller

import (
	"context"
	"fmt"
	"os"

	pgxpool "github.com/jackc/pgx/v4/pgxpool"
)

// PostgresController helps to manage interactions with the Postgres database
type PostgresController struct {
	dbName     string
	dbHost     string
	dbPort     int
	dbUser     string
	dbPassword string
	connURL    string
}

// NewPostgresController is the constructor for PostgresController
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

// InsertPilgrim inserts a pilgrim record into the database
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

// InsertInnkeeper inserts a innkeeper record into the database
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

// UpdatePilgrimPassword updates a pilgrim password from a record in the database
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

// SelectPilgrimByUsername retrieves a pilgrim by username from the database
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

// SelectInnkeeperByUsername retrieves a innkeeper by username from the database
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

// SelectProjectByID retrieves a project by ID from the database
func (pg *PostgresController) SelectProjectByID(id string) (string, string, error) {
	dbpool, err := pgxpool.Connect(context.Background(), pg.connURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	var title string
	var pilgrimID string
	err = dbpool.QueryRow(context.Background(),
		"SELECT pilgrimID, title FROM api.projects WHERE id=$1", id).Scan(&pilgrimID, &title)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}
	return pilgrimID, title, err
}

// SelectProjectByID retrieves all projects from the database
func (pg *PostgresController) SelectProjects() (map[string]string, error) {
	dbpool, err := pgxpool.Connect(context.Background(), pg.connURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	rows, _ := dbpool.Query(context.Background(), "SELECT pilgrimid, title FROM api.projects")

	projectsMap := make(map[string]string)

	for rows.Next() {
		var pilgrimid string
		var title string
		err := rows.Scan(&pilgrimid, &title)
		if err != nil {
			return nil, err
		}
		projectsMap[title] = pilgrimid
	}

	return projectsMap, rows.Err()
}
