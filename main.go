package main

import (
	"flag"
	"fmt"
	"os"

	stackdriver "github.com/TV4/logrus-stackdriver-formatter"
	"github.com/sirupsen/logrus"

	"github.com/IzePhanthakarn/go-phanthakarn-dev/docs"
	"github.com/IzePhanthakarn/go-phanthakarn-dev/internal/core/config"
	"github.com/IzePhanthakarn/go-phanthakarn-dev/internal/core/utils"
	"github.com/IzePhanthakarn/go-phanthakarn-dev/internal/handlers/routes"
)

func main() {
	environment := flag.String("environment", "", "set working environment")
	configs := flag.String("config", "configs", "set configs path, default as: 'configs'")
	flag.Parse()

	if environment == nil || *environment == "" {
		envVar := os.Getenv("environment")
		if envVar != "" {
			*environment = envVar
		}
	}
	logrus.Info("environment:", environment)

	// Init configuration
	err := config.InitConfig(*configs, *environment)
	if err != nil {
		panic(err)
	}
	// =======================================================

	// Init return result
	err = config.InitReturnResult("configs")
	if err != nil {
		panic(err)
	}
	// =======================================================

	// Get public key && private key (JWT)
	err = utils.ReadECDSAKey(config.CF.JWT.PrivateKeyPath, config.CF.JWT.PublicKeyPath)
	if err != nil {
		panic(err)
	}
	// =======================================================

	// programatically set swagger info
	docs.SwaggerInfo.Title = config.CF.Swagger.Title
	docs.SwaggerInfo.Description = config.CF.Swagger.Description
	docs.SwaggerInfo.Version = config.CF.Swagger.Version
	docs.SwaggerInfo.Host = fmt.Sprintf("%s%s", config.CF.Swagger.Host, config.CF.Swagger.BaseURL)
	// =======================================================

	// set logrus
	logrus.SetReportCaller(true)
	if config.CF.App.Release {
		logrus.SetFormatter(stackdriver.NewFormatter(
			stackdriver.WithService("api"),
			stackdriver.WithVersion("v1.0.0")))
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{})
	}
	logrus.Infof("Initial 'Configuration'. %+v", config.CF)

	routes.NewRouter()
	// ========================================================
}
