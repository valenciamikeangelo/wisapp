package main

import (
	"errors"
	"github.com/codegangsta/cli"
	"github.com/valenciamikeangelo/wisapp/wisappweb/server"
	"gopkg.in/yaml.v1"
	"io/ioutil"
	"log"
	"os"
)

func getConfig(c *cli.Context) (server.Config, error) {
	yamlPath := c.GlobalString("config")
	log.Print("Locating CONFIG File : " + yamlPath)
	config := server.Config{}

	if _, err := os.Stat(yamlPath); err != nil {
		return config, errors.New("config path not valid")
	}

	ymlData, err := ioutil.ReadFile(yamlPath)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal([]byte(ymlData), &config)

	return config, err
}

func main() {

	app := cli.NewApp()
	app.Name = "Wisapp Web App"
	app.Usage = "Web Client for Wisapp REST Gateway"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{"config, c", "config.yaml", "config file to use", "APP_CONFIG"},
	}

	app.Commands = []cli.Command{
		{
			Name:  "server",
			Usage: "Run the http server",
			Action: func(c *cli.Context) {
				cfg, err := getConfig(c)
				if err != nil {
					log.Fatal(err)
					return
				}

				var wisappServer = server.WisappServer{}

				if err = wisappServer.Run(cfg); err != nil {
					log.Fatal(err)
				}
			},
		},
	}
	app.Run(os.Args)
}
