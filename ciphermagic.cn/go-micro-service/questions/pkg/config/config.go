package config

import (
	"github.com/BurntSushi/toml"
	"os"
	"path/filepath"
	"runtime"
)

var (
	Config  = tomlConfig{}
	ViewDir string
	LogDir  string
)

type tomlConfig struct {
	RootDir string
	Port    string
	DBDebug bool
}

type configByOS struct {
	Windows tomlConfig
	OSX     tomlConfig
	Linux   tomlConfig
}

func init() {
	var configOS configByOS
	if _, err := toml.DecodeFile("go-micro-service/questions/config.toml", &configOS); err != nil {
		panic(err)
	}

	if runtime.GOOS == "windows" {
		Config = configOS.Windows
	} else if runtime.GOOS == "darwin" {
		Config = configOS.OSX
	} else {
		Config = configOS.Linux
	}

	ViewDir = filepath.Join(Config.RootDir, "templates")
	LogDir = filepath.Join(Config.RootDir, "logs")

	os.MkdirAll(LogDir, os.ModePerm)
}
