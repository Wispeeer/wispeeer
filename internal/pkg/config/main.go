package config

import (
	"fmt"
	"time"

	"github.com/creasty/defaults"
)

var (
	// PostsDir ...
	PostsDir string = "posts"
)

var (
	configFile = "config.yml"
	config     Config
	parserErr  error
)

func init() {
	initConfig()
}

func initConfig() {
	cf, err := getConfig(configFile)
	if err != nil {
		parserErr = fmt.Errorf("parse the %s : %v", configFile, err)
		return
	}
	if err := defaults.Set(&cf); err != nil {
		parserErr = fmt.Errorf("initial configuration : %v", err)
		return
	}
	config = cf
}

// GetWispeeerConfig ...
func GetWispeeerConfig() (Config, error) {
	return config, parserErr
}

// TimeZone ...
func TimeZone() *time.Location {
	var tz *time.Location

	if parserErr != nil {
		tz, _ = time.LoadLocation("local")
	} else {
		tz, _ = time.LoadLocation(config.Timezone)
	}
	return tz
}
