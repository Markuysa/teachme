package main

import (
	"context"
	"log"
	"os"

	"session/internal/app"
	"session/internal/config"

	cfgLoader "github.com/Markuysa/pkg/config"
)

const (
	cfgPathKey = "CONFIG_PATH"
)

func main() {
	ctx := context.Background()
	onBuild()

	cfgPath := os.Getenv(cfgPathKey)
	cfg := &config.Config{}

	err := cfgLoader.LoadFromYAML(cfg, cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	if err = app.Run(ctx, cfg); err != nil {
		log.Fatal(err)
	}
}
