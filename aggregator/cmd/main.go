package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"

	"github.com/connect3world/c3-data-avs/aggregator"
	"github.com/connect3world/c3-data-avs/core/config"
)

var (
	// Version is the version of the binary.
	Version   string
	GitCommit string
	GitDate   string
)

func main() {

	app := cli.NewApp()
	app.Flags = config.Flags
	app.Version = fmt.Sprintf("%s-%s-%s", Version, GitCommit, GitDate)
	app.Name = "c3-da-aggregator"
	app.Usage = "C3 DA Aggregator"
	app.Description = "Service that sends address to fetch C3 data by operator nodes."

	app.Action = aggregatorMain
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln("Application failed.", "Message:", err)
	}
}

func aggregatorMain(ctx *cli.Context) error {

	log.Println("Initializing Aggregator")
	config, err := config.NewConfig(ctx)
	if err != nil {
		return err
	}
	configJson, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		config.Logger.Fatalf(err.Error())
	}
	fmt.Println("Config:", string(configJson))

	agg, err := aggregator.NewAggregator(config)
	if err != nil {
		return err
	}

	err = agg.Start(context.Background())
	if err != nil {
		return err
	}

	return nil

}
