package main

import (
	"fmt"
	"github.com/AgustinCardozo/labo-so/memoria/globals"
	"github.com/AgustinCardozo/labo-so/utils/config"
	"github.com/AgustinCardozo/labo-so/utils/list"
)

func main() {
	var numberList list.ArrayList[int]

	numberList.Add(10)
	numberList.Add(20)
	numberList.Add(30)

	fmt.Println(numberList.Size())

	err := config.SetupConfig("./configs/memoria.json", &globals.Config)
	if err != nil {
		panic(err)
	}

	fmt.Println(globals.Config.Resources)

	//log.Logger.Info(fmt.Sprint(globals.Config.Resources))
	//log.Logger.Info(fmt.Sprint(globals.Config))
	//
	//log.SugarLogger.Infof("Port: %s", globals.Config.Port)
	//// o bien
	//
	//err = errors.New("Soy un error")
	//log.SugarLogger.Errorf("Error %v", err)
	//log.Error.JSON(err.Error())
}
