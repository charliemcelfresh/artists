package translations

import (
	"io/ioutil"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v1"
)

func init() {
	err := Load()
	if err != nil {
		logrus.Fatal(err)
	}
}

type Command struct {
	Use   string `yaml:"use"`
	Short string `yaml:"short"`
}

type Flag struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
}

type CobraCommands struct {
	Root        Command `yaml:"root"`
	Server      Command `yaml:"server"`
	LoadArtists Command `yaml:"load_artists"`
}

type EnvironmentVariables struct {
	DatabaseURL string `yaml:"database_url"`
}

type SQLDrivers struct {
	Postgres string `yaml:"postgres"`
}

type Errors struct {
	ErrorLoadingEnv         string `yaml:"error_loading_env"`
	ErrorConnectingToDB     string `yaml:"error_connecting_to_db"`
	ErrorPerformingQuery    string `yaml:"error_performing_query"`
	ErrorSubcommandRequired string `yaml:"error_subcommand_required"`
}

type Paths struct {
	Artists string `yaml:"artists"`
}

type Ports struct {
	EightyEighty string `yaml:"eighty_eighty"`
}

type Flags struct {
	LoadArtistsCountFlag Flag `yaml:"load_artists_count_flag"`
}

type Translations struct {
	Commands             CobraCommands        `yaml:"commands"`
	EnvironmentVariables EnvironmentVariables `yaml:"environment_variables"`
	SQLDrivers           SQLDrivers           `yaml:"sql_drivers"`
	Errors               Errors               `yaml:"errors"`
	Paths                Paths                `yaml:"paths"`
	Ports                Ports                `yaml:"ports"`
	Flags                Flags                `yaml:"flags"`
}

var StringValues Translations

func Load() error {
	filename, _ := filepath.Abs("./translations.yml")
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlFile, &StringValues)
	if err != nil {
		return err
	}
	return nil
}
