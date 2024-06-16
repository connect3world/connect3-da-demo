package actions

import (
	"encoding/json"
	"github.com/urfave/cli"
	"log"

	sdkutils "github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/connect3world/c3-data-avs/core/config"
	"github.com/connect3world/c3-data-avs/operator"
	"github.com/connect3world/c3-data-avs/types"
)

func RegisterOperatorWithEigenlayer(ctx *cli.Context) error {

	configPath := ctx.GlobalString(config.ConfigFileFlag.Name)
	nodeConfig := types.NodeConfig{}
	err := sdkutils.ReadYamlConfig(configPath, &nodeConfig)
	if err != nil {
		return err
	}
	// need to make sure we don't register the operator on startup
	// when using the cli commands to register the operator.
	nodeConfig.RegisterOperatorOnStartup = false
	configJson, err := json.MarshalIndent(nodeConfig, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println("Config:", string(configJson))

	operator, err := operator.NewOperatorFromConfig(nodeConfig)
	if err != nil {
		return err
	}

	err = operator.RegisterOperatorWithEigenlayer()
	if err != nil {
		return err
	}

	return nil
}
