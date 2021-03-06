/*
Sniperkit-Bot
- Status: analyzed
*/

package main

import (
	"fmt"
	"strings"

	"github.com/urfave/cli"

	"github.com/sniperkit/snk.fork.kudohamu-watchcat"
)

// Watch parses flags and starts to watch repositories.
func Watch(c *cli.Context) {
	conf := c.GlobalString("conf")
	accessToken := c.GlobalString("token")
	interval := c.GlobalString("interval")

	if interval == "" {
		interval = "30m"
	}

	watcher := watchcat.New(conf, interval, accessToken)

	for _, notifier := range strings.Split(c.GlobalString("notifiers"), ",") {
		switch notifier {
		case "std":
			watcher.AddNotifier(&watchcat.StdNotifier{})
		case "slack":
			webhookURL := c.GlobalString("slack_webhook_url")
			if webhookURL == "" {
				panic(fmt.Errorf("not specified `slack_webhook_url` flag"))
			}
			watcher.AddNotifier(&watchcat.SlackNotifier{
				WebhookURL: webhookURL,
			})
		default:
			panic(fmt.Errorf("invalid notifier: %s", notifier))
		}
	}

	if err := watcher.Watch(); err != nil {
		panic(err)
	}
}
