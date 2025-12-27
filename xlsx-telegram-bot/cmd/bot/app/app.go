package app

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	DB *pgxpool.Pool
	URL string
)