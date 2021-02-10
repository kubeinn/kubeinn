package test

import (
	auth_handler "github.com/kubeinn/kubeinn/src/internal/api/auth"
	"os"
)

func TestInitEnvironmentVars() {
	os.Setenv("PGDATABASE", "postgres")
	os.Setenv("PGHOST", "localhost")
	os.Setenv("PGPORT", "5432")
	os.Setenv("PGUSER", "postgres")
	os.Setenv("POSTGRES_PASSWORD", "pgpassword")
	os.Setenv("JWT_SIGNING_KEY", "bh3lfEY6f0hQ7TxHv0n8zj6s76ubN1hK")
	os.Setenv("PGTURL", "localhost")
	os.Setenv("PGTPORT", "3000")
	os.Setenv("PROMETHEUS_URL", "51.222.35.240")
	os.Setenv("PROMETHEUS_PORT", "30010")
}

func TestCreateDefaultInnkeeper() {
	auth_handler.RegisterInnkeeper("admin", "admin", "admin")
}
