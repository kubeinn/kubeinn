package test

import (
	global "github.com/kubeinn/schutterij/internal/global"
	"log"
	"os"
)

func TestInitEnvironmentVars() {
	os.Setenv("PGDATABASE", "postgres")
	os.Setenv("PGHOST", "localhost")
	os.Setenv("PGPORT", "5432")
	os.Setenv("PGUSER", "postgres")
	os.Setenv("POSTGRES_PASSWORD", "pgpassword")
	os.Setenv("JWT_SIGNING_KEY", "bh3lfEY6f0hQ7TxHv0n8zj6s76ubN1hK")
}

func TestDatabaseConnection() {
	userID, password, err := global.PG_CONTROLLER.SelectInnkeeperByUsername("test-user-01")
	if err != nil {
		log.Println(err)
	}
	log.Println("Innkeeper: " + string(userID) + " " + password)
	userID, password, err = global.PG_CONTROLLER.SelectPilgrimByUsername("test-user-01")
	if err != nil {
		log.Println(err)
	}
	log.Println("Pilgrim password: " + string(userID) + " " + password)
}
