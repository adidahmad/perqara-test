package main

import (
	"fmt"
	"log"
	"os"

	"github.com/adidahmad/perqara-test/config"
	"github.com/adidahmad/perqara-test/modules/databases/mysql"
	"github.com/pressly/goose/v3"
	"github.com/spf13/cobra"
)

var usageCommands = `
Run database migrations

Usage:
    perqara-migrate [command]

Available Migration Commands:
    up                   Migrate the DB to the most recent version available
    up-to VERSION        Migrate the DB to a specific VERSION
    down                 Roll back the version by 1
    down-to VERSION      Roll back to a specific VERSION
    redo                 Re-run the latest migration
    status               Dump the migration status for the current DB
    version              Print the current version of the database
    create NAME [sql|go] Creates new migration file with next version	
`

func main() {
	var rootCmd = &cobra.Command{
		Use:   "perqara-migrate",
		Short: "MySql Migration Perqara Service",

		Run: func(cmd *cobra.Command, args []string) {
			config.Load()

			if len(args) == 0 {
				cmd.Help()
				os.Exit(0)
			}

			goose.SetDialect("mysql")
			dbCon, err := mysql.MySQLConnect(config.AppConf.DBConfig)
			if err != nil {
				log.Fatal(err)
				panic(err)
			}

			appPath, _ := os.Getwd()
			dir := appPath + "/cmd/migration/files"
			if len(args) == 0 {
				cmd.Help()
				os.Exit(0)
			}

			command := args[0]
			cmdArgs := args[1:]
			optAllowMissing := goose.WithAllowMissing()
			if err := goose.RunWithOptions(command, dbCon, dir, cmdArgs, optAllowMissing); err != nil {
				log.Fatalf("goose run: %v", err)
			}
		},
	}

	rootCmd.SetUsageTemplate(usageCommands)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
