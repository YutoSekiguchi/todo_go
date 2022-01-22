package main

import (
	"github.com/YutoSekiguchi/todo/router"
	"github.com/YutoSekiguchi/todo/util"
)

func main() {
	db := util.InitDb()
	router.InitRouter(db)
}