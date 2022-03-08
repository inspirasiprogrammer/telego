package main

import (
	"fmt"
	"os"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

func main() {
	botToken := os.Getenv("TOKEN")

	bot, err := telego.NewBot(botToken, telego.WithDefaultLogger(true, true))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Get updates channel
	updates, _ := bot.UpdatesViaLongPulling(nil)
	defer bot.StopLongPulling()

	// Create bot handler and specify from where to get updates
	bh := th.NewBotHandler(bot, updates)

	// Register handler with union predicate and not predicate
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		fmt.Println("Update with message text `Hmm?` or any other, but without message.")
	}, th.Union(
		th.Not(th.AnyMassage()), // Matches to any not message update
		th.TextEqual("Hmm?"),    // Matches to message update with text `Hmm?`
	))

	// Register handler with message predicate and custom predicate
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		fmt.Println("Update with message which text is longer then 7 chars.")
	},
		th.AnyMassage(), // Matches to any message update
		func(update telego.Update) bool { // Matches to message update with text longer then 7
			return len(update.Message.Text) > 7
		},
	)

	// Register handler with commands and specific args
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		fmt.Println("Update with command `start` without args or `help` with any args")
	}, th.Union(
		th.CommandEqualArgc("start", 0),     // Matches only to `/start`
		th.CommandEqualArgv("how", "works"), // Matches only to `/how works`
		th.CommandEqual("help"),             // Matches to `/help` with any args
	))

	// Start handling updates
	bh.Start()

	// Stop handling updates
	defer bh.Stop()
}