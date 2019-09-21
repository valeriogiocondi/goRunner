package filesListPkg

import (
	"io/ioutil"
	"log"
	"os"

	"packages/util"
)


func GetFilesList(path string) {

	filesList := GetIteratorFilesList(path)
	fetchFilesList(filesList, path)
}

func GetIteratorFilesList(path string) []os.FileInfo {

	filesList, error := ioutil.ReadDir(path)

	if error != nil {
		log.Fatal(error)
	}
	return filesList
}

func fetchFilesList(list []os.FileInfo, initPath string) {

	for _, f := range list {

		path := initPath + f.Name()

		if f.IsDir() {
			
			path += "/"

			if util.IsDirToCheck(path) {

				GetFilesList(path)
			}
		
		} else {

			if (util.IsFileToCheck(path) && util.CheckFileUpdated(path, f.ModTime().String())) {
				
				break
			}
		}
	}

	return
}