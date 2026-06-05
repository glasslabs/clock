//go:build wasip1

package main

import (
	"time"
	_ "time/tzdata"

	"github.com/glasslabs/client-go"
)

// Config is the module configuration.
type Config struct {
	TimeFormat string `json:"timeFormat"`
	DateFormat string `json:"dateFormat"`
	Timezone   string `json:"timezone"`
}

// NewConfig returns a Config with default values.
func NewConfig() Config {
	return Config{
		TimeFormat: "15:04",
		DateFormat: "Monday, January 2",
		Timezone:   "Local",
	}
}

var (
	mod *client.Module
	log *client.Logger
	cfg Config
	loc *time.Location
)

func main() {
	log = client.NewLogger()

	var err error
	mod, err = client.NewModule()
	if err != nil {
		log.Error("Could not create module", "error", err.Error())
		return
	}

	cfg = NewConfig()
	if err = mod.ParseConfig(&cfg); err != nil {
		log.Error("Could not parse config", "error", err.Error())
		return
	}

	if cfg.Timezone != "" {
		loc, err = time.LoadLocation(cfg.Timezone)
		if err != nil {
			log.Error("Invalid timezone", "error", err.Error(), "timezone", cfg.Timezone)
			//nolint:gosmopolitan // Used as fallback only.
			loc = time.Local
		}
	}

	log.Info("Module ready", "module", mod.Name())

	render()

	for {
		time.Sleep(10 * time.Second)
		render()
	}
}

func render() {
	now := time.Now()
	if loc != nil {
		now = now.In(loc)
	}

	w := client.NewVStack(
		client.NewText(now.Format(cfg.TimeFormat),
			client.WithColor("#ffffff"),
			client.WithFontSize(95),
			client.WithLight(),
			client.WithAlign("right"),
		),
		client.NewText(now.Format(cfg.DateFormat),
			client.WithColor("#cccccc"),
			client.WithFontSize(24),
			client.WithCondensed(),
			client.WithLight(),
			client.WithAlign("right"),
		),
	)

	mod.Render(w)
}
