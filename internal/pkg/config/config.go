package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// Config ...
type Config struct {
	Site       `yaml:",inline"`
	Urls       `yaml:",inline"`
	Directory  `yaml:",inline"`
	Pagination `yaml:",inline"`
	Deployment `yaml:",inline"`
}

// Site ...
type Site struct {
	Title       string `yaml:"title" default:"Welcome to wisper"`
	Subtitle    string `yaml:"subtitle" default:"wispeeer wisper"`
	Description string `yaml:"description" default:"Generate by Wispeeer"`
	Keywords    string `yaml:"keywords" default:"blog"`
	Author      string `yaml:"author" default:"void"`
	Theme       string `yaml:"theme" default:"wisper"`
	Timezone    string `yaml:"timezone" default:"Local"`
}

// Urls ...
type Urls struct {
	Root      string `yaml:"root" default:"http://localhost:1080"`
	Permalink string `yaml:"permalink" default:"/website/"`
}

// Directory ...
type Directory struct {
	SourceDir string `yaml:"source_dir" default:"source"`
	PublicDir string `yaml:"public_dir" default:"public"`
}

// Pagination ...
type Pagination struct {
	PerPage       string `yaml:"per_page" default:"9"`
	PaginationDir string `yaml:"pagination_dir" default:"page"`
}

// Deployment ...
type Deployment struct {
	Type   string `yaml:"type" default:"git"`
	Repo   string `yaml:"repo" default:"void"`
	Branch string `yaml:"branch" default:"master"`
}

func getConfig(filename string) (Config, error) {
	config := Config{}

	configFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return config, fmt.Errorf("%s files Not Found", filename)
	}

	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		return config, fmt.Errorf("%s files is corrupted", filename)
	}
	return config, nil
}
