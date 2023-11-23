package main

import (
	"github.com/inabatron/go_crud/inits"
	"github.com/inabatron/go_crud/models"
)

func init() {
	inits.LoadEnvVars()
	inits.ConnectToDB()
}

func main() {
	inits.DB.AutoMigrate(&models.Post{})
}
