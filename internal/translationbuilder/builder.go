package translationbuilder

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"strings"

	"gopkg.in/yaml.v3"
)

func BuildYAMLStructDefinition(filepath string) error {
	fileContents, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}
	yamlMap := map[string]interface{}{}
	err = yaml.Unmarshal(fileContents, yamlMap)
	if err != nil {
		return err
	}
	filepathComponents := strings.Split(filepath, "/")

	// the directory of the file we read is the package in which we want to place our generated struct
	packageName := filepathComponents[len(filepathComponents)-2]

	destinationFile, err := os.Create(strings.Trim(filepath, ".yml") + "_struct.go")
	structFileContents := bufio.NewWriter(destinationFile)
	fmt.Fprintf(structFileContents, "package %s\n", packageName)
	err = addStruct(true, strings.Trim(filepathComponents[len(filepathComponents)-1], ".yml"), yamlMap, structFileContents)
	if err != nil {
		return err
	}

	err = structFileContents.Flush()
	return err
}

func addStruct(isTopLevel bool, nameFromYAML string, contentMap map[string]interface{}, buffer io.Writer) error {
	if isTopLevel {
		_, err := fmt.Fprintf(buffer, "type %s struct {\n", buildStructNameFromYAMLName(nameFromYAML))
		if err != nil {
			return err
		}
	} else {
		_, err := fmt.Fprintf(buffer, "%s struct {\n", buildStructNameFromYAMLName(nameFromYAML))
		if err != nil {
			return err
		}
	}

	for key, value := range contentMap {
		fmt.Println(key)
		switch value.(type) {
		case string:
			_, err := fmt.Fprintf(buffer, "\t%s string `yaml:\"%s\"`\n", buildStructNameFromYAMLName(key), key)
			if err != nil {
				return err
			}

		case map[string]interface{}:
			err := addStruct(false, key, value.(map[string]interface{}), buffer)
			if err != nil {
				return err
			}
		default:
			fmt.Println(reflect.TypeOf(value).String())
		}
	}
	if isTopLevel {
		_, err := fmt.Fprintf(buffer, "}\n")
		if err != nil {
			return err
		}
	} else {
		_, err := fmt.Fprintf(buffer, "} `yaml:\"%s\"`\n", nameFromYAML)
		if err != nil {
			return err
		}
	}
	return nil
}

func buildStructNameFromYAMLName(name string) string {
	name = strings.ReplaceAll(name, "_", " ")
	name = strings.Title(name)
	name = strings.ReplaceAll(name, " ", "")
	return name
}
