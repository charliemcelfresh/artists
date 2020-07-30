package main

import "artists/cmd"

//go:generate ./artists buildYAMLStruct internal/translations/translations.yml
func main() {
	cmd.Execute()
}
