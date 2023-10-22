//go:build js && wasm

package main

import (
	_ "embed"
	"fmt"
	"time"
	_ "time/tzdata"

	"github.com/glasslabs/client-go"
)

var (
	//go:embed assets/style.css
	css []byte

	//go:embed assets/index.html
	html []byte
)

// Config is the module configuration.
type Config struct {
	TimeFormat string
	DateFormat string
	Timezone   string
}

// NewConfig returns a Config with default values set.
func NewConfig() Config {
	return Config{
		TimeFormat: "15:04",
		DateFormat: "Monday, January 2",
		Timezone:   "Local",
	}
}

func main() {
	log := client.NewLogger()
	mod, err := client.NewModule()
	if err != nil {
		log.Error("Could not create module", "error", err.Error())
		return
	}

	cfg := NewConfig()
	if err = mod.ParseConfig(&cfg); err != nil {
		log.Error("Could not parse config", "error", err.Error())
		return
	}

	log.Info("Loading Module", "module", mod.Name())

	m := &Module{
		mod: mod,
		cfg: cfg,
		log: log,
	}

	if err = m.setup(); err != nil {
		log.Error("Could not setup module", "error", err.Error())
		return
	}

	tick := time.NewTicker(10 * time.Second)
	defer tick.Stop()

	for {
		m.update()

		<-tick.C
	}
}

// Module runs the module.
type Module struct {
	mod *client.Module
	cfg Config

	loc *time.Location

	log *client.Logger
}

func (m *Module) setup() error {
	if m.cfg.Timezone != "" {
		l, err := time.LoadLocation(m.cfg.Timezone)
		if err != nil {
			return fmt.Errorf("invalid timezone: %w", err)
		}
		m.loc = l
	}

	if err := m.mod.LoadCSS(string(css)); err != nil {
		return fmt.Errorf("loading css: %w", err)
	}
	m.mod.Element().SetInnerHTML(string(html))

	return nil
}

func (m *Module) update() {
	now := time.Now()
	if m.loc != nil {
		now = now.In(m.loc)
	}

	m.mod.Element().QuerySelector(".time").SetInnerHTML(now.Format(m.cfg.TimeFormat))
	m.mod.Element().QuerySelector(".date").SetInnerHTML(now.Format(m.cfg.DateFormat))
}
