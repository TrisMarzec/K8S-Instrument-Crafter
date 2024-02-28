package generator

import (
	"fmt"
	"os"
)

func GenerateYAMLFile(object interface{}, fileName string) {

	yamlData, err := yaml.Marshal(&object)

	if err != nil {
		panic(err)
	}

	err = os.WriteFile(fileName, yamlData, 0644)

	if err != nil {
		panic(err)
	}
}

func UpdateYAMLFile(object interface{}, fileName string) {

}

func DeleteYAMLFile(fileName string) {

	err := os.Remove(fileName)

	if err != nil {
		panic(err)
	}
	fmt.Println("File successfully deleted...")
}
