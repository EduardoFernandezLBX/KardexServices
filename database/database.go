package database

import (
	"context"
	"fmt"

	"github.com/EduardoFernandezLBX/ProjectKardex.git/settings"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func New(ctx context.Context, s *settings.Settings) (*sqlx.DB, error) {
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?tls=true",
		s.DB.Username,
		s.DB.Password,
		s.DB.Host,
		s.DB.Name,
	)

	return sqlx.ConnectContext(ctx, "mysql", connectionString)
}
