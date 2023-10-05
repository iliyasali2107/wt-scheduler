package main

import (
	"github.com/iliyasali2107/wt-scheduler/internal/services/user_service"
	"github.com/iliyasali2107/wt-scheduler/internal/storage/postgres"
)

func main() {
	psql := postgres.NewPostgresStorage()

	serv := user_service.NewUserService(psql)
}
