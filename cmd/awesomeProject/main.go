// Copyright (C) 2021-2022 Amuzed GmbH finn@amuzed.io.
// This file is part of the project AMUZED.
// AMUZED can not be copied and/or distributed without the express.
// permission of Amuzed GmbH.

package main

import (
	"awesomeProject"
	"awesomeProject/database"
	"context"
	"github.com/dmitrymomot/go-env"
	"github.com/spf13/cobra"
	"github.com/zeebo/errs"
	"os"
)

var (
	databaseConnStr      = env.MustString("DATABASE_URL")
	consoleServerAddress = env.MustString("CONSOLE_SERVER_ADDR")
)

// Config contains configurable values for amuzed project.
type Config struct {
	Database              string `json:"database"`
	awesomeProject.Config `json:"config"`
}

// commands.
var (
	// root cmd.
	rootCmd = &cobra.Command{
		Use:   "awesome",
		Short: "cli for interacting with awesome project",
	}

	runCmd = &cobra.Command{
		Use:         "run",
		Short:       "runs the program",
		RunE:        cmdRun,
		Annotations: map[string]string{"type": "run"},
	}
)

func init() {
	rootCmd.AddCommand(runCmd)
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func cmdRun(cmd *cobra.Command, args []string) (err error) {
	ctx := context.Background()

	runCfg := createConfig()

	db, err := database.New(ctx, runCfg.Database)
	if err != nil {
		return err
	}
	defer db.Close()

	db.CreateSchema(ctx)

	peer, err := awesomeProject.New(runCfg.Config, db)
	if err != nil {
		return err
	}

	runError := peer.Run(ctx)
	closeError := peer.Close()

	return errs.Combine(runError, closeError)
}

func createConfig() Config {
	var cfg Config
	cfg.Console.Server.Address = consoleServerAddress
	cfg.Database = databaseConnStr
	return cfg
}
