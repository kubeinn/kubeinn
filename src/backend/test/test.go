package test

import (
	auth_handler "github.com/kubeinn/src/backend/internal/api/auth"
	"os"
)

func TestInitEnvironmentVars() {
	os.Setenv("PGDATABASE", "postgres")
	os.Setenv("PGHOST", "localhost")
	os.Setenv("PGPORT", "5432")
	os.Setenv("PGUSER", "postgres")
	os.Setenv("POSTGRES_PASSWORD", "pgpassword")
	os.Setenv("JWT_SIGNING_KEY", "bh3lfEY6f0hQ7TxHv0n8zj6s76ubN1hK")
	os.Setenv("POSTGREST_URL", "http://192.168.0.130:3000")
}

func TestCreateDefaultInnkeeper() {
	auth_handler.RegisterInnkeeper("admin", "admin", "admin")
}
