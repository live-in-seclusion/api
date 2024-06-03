package main

import (
	"gitlab.com/btlike/api/server"
	"gitlab.com/btlike/api/utils"
)

func main() {
	utils.Init()
	server.Run(utils.Config.Address)
}
