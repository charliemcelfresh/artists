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
