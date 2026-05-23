package utils

import (
	"os"
	"path"

	"github.com/BurntSushi/toml"
)

func GetCurrDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return dir, nil
}

func GetConfigPath() (string, error) {
	cwd, err := GetCurrDir()
	if err != nil {
		return "", err
	}

	return path.Join(cwd, "config.toml"), nil
}

type ConfigFile struct {
	Name    string `toml:"app_name"`
	Version string

	DownloadPath     string
	FilenameTemplate string
}

var Config ConfigFile

func LoadConfigFile(path string) error {
	_, err := toml.DecodeFile(path, &Config)
	if err != nil {
		return err
	}
	return nil
}

func InitConfigFile(f *os.File) error {
	c := ConfigFile{
		Name:             "Golash",
		Version:          "0.0.1",
		DownloadPath:     "download",
		FilenameTemplate: "%(title)s.%(ext)s",
	}

	err := toml.NewEncoder(f).Encode(c)
	if err != nil {
		return err
	}

	return nil
}

func GetConfigMetadata() *ConfigFile {
	return &Config
}
