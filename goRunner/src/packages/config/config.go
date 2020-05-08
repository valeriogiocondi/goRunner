package config

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"packages/vars"
	"packages/filesListPkg"
)


const mainPath string = "./goRunner/"

type fileConfig struct {

	TaskList []string 
	ExceptFile []string 
	ExceptDirectories []string 
}

func Load() {

	/* 
	 *	Load all files and config.json
	 *
	 */
	loadFilesList(vars.InitPath)
	loadConfig(vars.TaskList, vars.ExceptionFileExtension, vars.ExceptionDirectories)

	fmt.Println("\n\n WATCHER STARTED \n\n")
}

func loadConfig(taskList []string, exceptionFileExtension []string, exceptionDirectories []string) {

	var contentFile string
	var configOptions fileConfig;
	file, err := os.Open(mainPath + "config.json")

	if (err != nil) {
		log.Fatal(err)
	}
	defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {            
        
        contentFile += scanner.Text();
    }

	/* 
	 *	Save config.json into global data structure
	 *
	 */
    err = json.Unmarshal([]byte(contentFile), &configOptions)
	
	if (err != nil) {
		
		fmt.Println("\n\nJSON VALIDATION ERROR:")
		log.Fatal(err)
	}

 	vars.TaskList = configOptions.TaskList
	vars.ExceptionFileExtension = configOptions.ExceptFile
	vars.ExceptionDirectories = configOptions.ExceptDirectories

    return
}

func loadFilesList(initPath string) {

	filesList := filesListPkg.GetIteratorFilesList(initPath)

	for _, f := range filesList {

		path := initPath + f.Name()

		if f.IsDir() {

			loadFilesList(path + "/")
		
		} else {

			loadDataStructure(path, f.ModTime().String())
		}
	}
}

func loadDataStructure(path string, timestampLastModification string) {

	vars.MapFiles[path] = timestampLastModification
	fmt.Println("\033[1mLOADED:\033[0m " + path)
	return
}