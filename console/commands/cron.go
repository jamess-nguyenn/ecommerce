package commands

import (
	"github.com/robfig/cron/v3"
)

// InitializeCron initializes and starts cron
func InitializeCron() {
	c := cron.New()

	// list cron
	c.AddFunc("* * * * *", exampleCommand)
	c.AddFunc("*/2 * * * *", sampleCommand)

	c.Start()
}
