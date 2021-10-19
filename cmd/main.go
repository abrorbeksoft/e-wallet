package main

import (
	"database/sql"
	"fmt"
	"github.com/abrorbeksoft/e-wallet.git/api"
	"github.com/abrorbeksoft/e-wallet.git/config"
	"github.com/abrorbeksoft/e-wallet.git/storage"
	_ "github.com/lib/pq"
)

func main()  {
	config:=config.Load();
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", config.PostgresUser,config.PostgresPassword,config.PostgresDBName)

	session, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	db:=storage.NewStorage(session)


	server:=api.New(&api.RouterOptions{
		Storage: db,
	})

	server.Run()
}


