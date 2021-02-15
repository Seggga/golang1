package main

import (
	"fmt"
	"os"

	"github.com/seggga/golang1/08_configreader/conf"
)

func main() {
	var config conf.ConfigStruct

	//err := conf.ReadEnv(&config) // fill in the config-structure with OS variables
	//err := conf.ReadFlags(&config) // fill in the config-structure with flags-data

	confiigFile := os.Args[1]
	err := conf.ReadFile(&config, confiigFile)

	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	// вывод итоговой структуры
	fmt.Printf("%+v", config)
}
