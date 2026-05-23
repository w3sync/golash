package main

import (
	"context"
	"os"

	"github.com/lrstanley/go-ytdlp"
	"github.com/w3sync/golash.git/downloader"
	"github.com/w3sync/golash.git/utils"
)

const URL = "https://www.youtube.com/watch?v=dQw4w9WgXcQ&pp=ygUNcm9jayBhbmQgcm9sbA%3D%3D"

func init() {
	ytdlp.MustInstallBun(context.Background(), nil)

	configPath, err := utils.GetConfigPath()
	if err != nil {
		panic(err)
	}

	_, err = os.Stat(configPath)
	if err != nil {
		configFile, err := os.Create(configPath)
		if err != nil {
			panic(err)
		}

		err = utils.InitConfigFile(configFile)
		if err != nil {
			panic(err)
		}

		defer configFile.Close()
	} else {
		utils.LoadConfigFile(configPath)
	}
}

func main() {
	downloader.Download(URL)
}
