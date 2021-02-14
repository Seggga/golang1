package main

import (
	"fmt"

	"github.com/seggga/golang1/08_configreader/conf"
)

func main() {
	var config conf.ConfigType

	//err := conf.ReadEnv(&config) // fill in the config-structure with OS variables
	err := conf.ReadFlags(&config) // fill in the config-structure with flags-data
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	// вывод итоговой структуры
	fmt.Printf("%+v", config)
}
