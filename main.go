package main

import (
	"log"
	"os"

	cli "github.com/jawher/mow.cli"

	"github.com/LarsFox/diploma-bot/core"
	"github.com/LarsFox/diploma-bot/tg"
)

var (
	app = cli.App("tgfinbot", "Launches a Test bot for my diploma project")

	tgToken   = app.StringOpt("tg-token", "", "Telegram Bot token")
	tgAdminID = app.IntOpt("tg-admin-id", 0, "ID of the user, who will get all the input")
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(os.Stdout)
	app.Action = appAction
}

func main() {
	if err := app.Run(os.Args); err != nil {
		log.Fatalln("cli: ", err)
	}
}

func appAction() {
	tgClient := tg.NewClient(*tgToken)
	appManager := core.NewManager(tgClient, *tgAdminID)
	appManager.Listen()
}
