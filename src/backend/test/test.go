package test

import (
	auth_handler "github.com/kubeinn/schutterij/internal/api/auth"
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

func TestCreateDefaultInnkeeper() {
	auth_handler.RegisterInnkeeper("admin", "admin", "admin")
}

func TestCreateDefaultReeve() {
	auth_handler.RegisterReeve("village-1", "reeve", "reeve", "reeve")
}
