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
}

func TestDatabaseConnection() {
	password, err := global.PG_CONTROLLER.SelectInnkeeperPasswordByUsername("test-user-01")
	if err != nil {
		log.Println(err)
	}
	log.Println("Innkeeper password: " + password)
	password, err = global.PG_CONTROLLER.SelectPilgrimPasswordByUsername("test-user-01")
	if err != nil {
		log.Println(err)
	}
	log.Println("Pilgrim password: " + password)
}
